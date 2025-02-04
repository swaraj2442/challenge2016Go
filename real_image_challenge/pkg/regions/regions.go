package regions

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func LoadRegionsCSV(filename string) map[string][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	regionMap := make(map[string][]string)
	for _, row := range rows {
		if len(row) < 6 {
			continue
		}
		city, state, country := strings.TrimSpace(row[3]), strings.TrimSpace(row[4]), strings.TrimSpace(row[5])
		key1 := fmt.Sprintf("%s-%s-%s", city, state, country)
		key2 := fmt.Sprintf("%s-%s", state, country)
		key3 := country

		regionMap[strings.ToUpper(key1)] = []string{key1, key2, key3}
		regionMap[strings.ToUpper(key2)] = []string{key2, key3}
		regionMap[strings.ToUpper(key3)] = []string{key3}
	}

	return regionMap
}

func MatchRegion(region string, regionMap map[string][]string) (string, bool) {
	region = strings.ToUpper(region)
	if val, exists := regionMap[region]; exists {
		return val[0], true
	}
	return "", false
}
