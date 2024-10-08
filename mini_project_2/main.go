package main

import (
	CoreFunctions "golang_basics/mini_project_2/core_functions"
)

func main() {
	var grid_factions_role = CoreFunctions.GenerateGrid()
	var grid_human_full = CoreFunctions.GenerateGrid()
	var grid_cylon_full = CoreFunctions.GenerateGrid()

	var personels = CoreFunctions.GetPilots(1)
	var raiders = CoreFunctions.GetRaiders(1)

	// CoreFunctions.GetPilots(0) // <- Will spawn an error
	// CoreFunctions.GetPilot(8) // <- Will spawn an error

	var coord_x = 0
	var coord_y = 0

	// Place human ships unto the grid
	for i := 0; i < len(personels); i++ {
		if coord_x < 5 && coord_y < 5 {
			grid_factions_role[coord_x][coord_y]["value"] = personels[i].Role
			coord_x++
			coord_y++
		}
	}

	coord_x = 0
	coord_y = 2

	// Place cylon ships unto the grid
	for i := 0; i < len(raiders); i++ {
		if coord_y < 5 {
			grid_factions_role[coord_x][coord_y]["value"] = raiders[i].Role
			coord_y++
		}
	}

	coord_x = 0
	coord_y = 0

	// Add humans into the Field Data
	for i := 0; i < len(personels); i++ {
		if coord_x < 4 && coord_y < 4 {
			grid_human_full[coord_x][coord_y]["value"] = personels[i].Role + ", " + personels[i].Rank + ", " + personels[i].Alias + ", " + personels[i].Name
			coord_x++
			coord_y++
		}
	}

	coord_x = 2
	coord_y = 0

	// Add cylons into the Field Data
	for i := 0; i < len(raiders); i++ {
		if coord_x < 5 {
			grid_cylon_full[coord_x][coord_y]["value"] = raiders[i].Role + ", " + raiders[i].Rank + ", " + raiders[i].Alias + ", " + raiders[i].Name
			coord_x++
		}
	}

	CoreFunctions.DisplayGrid(grid_factions_role)
	CoreFunctions.FieldDataTitleBanner()
	CoreFunctions.FieldData("Human", grid_human_full)
	CoreFunctions.FieldData("Cylon", grid_cylon_full)
}
