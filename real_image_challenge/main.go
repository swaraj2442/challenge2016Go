package main

import (
	"real_image_challenge/pkg/permissions"
	"real_image_challenge/pkg/queries"
	"real_image_challenge/pkg/regions"
)

func main() {
	regionMap := regions.LoadRegionsCSV("data/cities.csv")
	distributors := permissions.LoadPermissionsTxt("data/permissions.txt", regionMap)
	queries.LoadQueriesTxt("data/queries.txt", distributors, regionMap)
}
