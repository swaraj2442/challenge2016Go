package permissions

import (
	"bufio"
	"fmt"
	"os"
	"real_image_challenge/pkg/regions"
	"strings"
)

type Permissions struct {
	Includes map[string]bool
	Excludes map[string]bool
	Parent   *Permissions
}

func NewPermissions(parent *Permissions) *Permissions {
	return &Permissions{
		Includes: make(map[string]bool),
		Excludes: make(map[string]bool),
		Parent:   parent,
	}
}

func (p *Permissions) AddInclude(region string) { p.Includes[region] = true }
func (p *Permissions) AddExclude(region string) { p.Excludes[region] = true }

func (p *Permissions) HasPermission(region string) bool {
	if p.Excludes[region] {
		return false
	}
	if p.Includes[region] {
		return true
	}
	if p.Parent != nil {
		return p.Parent.HasPermission(region)
	}
	return false
}

func LoadPermissionsTxt(filename string, regionMap map[string][]string) map[string]*Permissions {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	distributors := make(map[string]*Permissions)
	var current, parent string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "Permissions for") {
			parts := strings.Split(line, "<")
			current = strings.TrimSpace(strings.Split(parts[0], "for")[1])
			parent = ""
			if len(parts) > 1 {
				parent = strings.TrimSpace(parts[1])
			}

			if parent == "" {
				distributors[current] = NewPermissions(nil)
			} else if parentPerms, exists := distributors[parent]; exists {
				distributors[current] = NewPermissions(parentPerms)
			} else {
				fmt.Printf("Warning: Parent distributor %s not found for %s\n", parent, current)
				distributors[current] = NewPermissions(nil)
			}
		} else if strings.HasPrefix(line, "INCLUDE:") || strings.HasPrefix(line, "EXCLUDE:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue
			}
			action, region := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

			regionCode, exists := regions.MatchRegion(region, regionMap)
			if !exists {
				fmt.Printf("Warning: Unknown region %s\n", region)
				continue
			}

			if current == "" {
				fmt.Println("Error: No distributor defined before permissions")
				continue
			}

			if action == "INCLUDE" {
				distributors[current].AddInclude(regionCode)
			} else if action == "EXCLUDE" {
				distributors[current].AddExclude(regionCode)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return distributors
}
