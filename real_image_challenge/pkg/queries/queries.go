package queries

import (
	"bufio"
	"fmt"
	"os"
	"real_image_challenge/pkg/permissions"
	"real_image_challenge/pkg/regions"
	"strings"
)

// LoadQueriesTxt loads queries and checks permissions for distributors
func LoadQueriesTxt(filename string, distributors map[string]*permissions.Permissions, regionMap map[string][]string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if !strings.HasPrefix(line, "Can ") || !strings.HasSuffix(line, "?") {
			fmt.Println("Error: Invalid query format:", line)
			continue
		}

		line = strings.TrimSuffix(strings.TrimPrefix(line, "Can "), "?")
		parts := strings.Split(line, " distribute in ")
		if len(parts) != 2 {
			fmt.Println("Error: Malformed query:", line)
			continue
		}

		distributor := strings.TrimSpace(parts[0])
		region := strings.TrimSpace(parts[1])

		regionCode, exists := regions.MatchRegion(region, regionMap)
		if !exists {
			fmt.Printf("Warning: Unknown region %s\n", region)
			continue
		}

		perm, exists := distributors[distributor]
		if !exists {
			fmt.Printf("Warning: Distributor %s not found\n", distributor)
			continue
		}

		result := "NO"
		if perm.HasPermission(regionCode) {
			result = "YES"
		}

		fmt.Printf("Query: Can %s distribute in %s? %s\n", distributor, region, result)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
