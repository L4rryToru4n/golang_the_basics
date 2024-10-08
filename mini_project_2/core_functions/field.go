package CoreFunctions

import (
	"fmt"
	"strconv"
)

type Grid [5][5]map[string]string

func GenerateGrid() Grid {

	var newGrid Grid

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {

			// Initizalizing map
			newGrid[i][j] = make(map[string]string)

			var coord_x string = strconv.Itoa(i)
			var coord_y string = strconv.Itoa(j)

			newGrid[i][j]["coord_x"] = coord_x
			newGrid[i][j]["coord_y"] = coord_y
			newGrid[i][j]["value"] = ""
		}
	}

	return newGrid
}

func DisplayGrid(grid Grid) {

	fmt.Printf("\n")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("========================| Grid |========================")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("\n")

	fmt.Printf("    ")
	for i := 0; i < len(grid); i++ {
		fmt.Printf("  X%d  ", i)
	}
	fmt.Println()

	fmt.Printf("     ")
	for i := 0; i < len(grid)*5; i++ {
		fmt.Printf("-")
	}

	fmt.Println()
	for x := 0; x < len(grid); x++ {
		fmt.Printf("Y%d", x)
		fmt.Printf("|")
		for y := 0; y < len(grid); y++ {
			if grid[x][y]["value"] != "" {
				if grid[x][y]["value"] == "Viper Pilot" {
					fmt.Printf("   V  ")
				} else if grid[x][y]["value"] == "Raptor Pilot" {
					fmt.Printf("   R  ")
				} else if grid[x][y]["value"] == "Raider Pilot" {
					fmt.Printf("   C  ")
				}
			} else {
				fmt.Printf("   .  ")
			}
		}
		fmt.Println()
	}
}

func FieldDataTitleBanner() {
	fmt.Printf("\n")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("=====================| Field Data |=====================")
	fmt.Println("--------------------------------------------------------")
}

func FieldData(faction_name string, grid Grid) {

	if faction_name == "Human" {
		fmt.Printf("\n")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>> { Humans } <<<<<<<<<<<<<<<<<<<<<<<")
	} else if faction_name == "Cylon" {
		fmt.Printf("\n")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>> { Cylons } <<<<<<<<<<<<<<<<<<<<<<<")
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j]["value"] != "" {
				fmt.Printf("(X%s, Y%s): ", grid[i][j]["coord_x"], grid[i][j]["coord_y"])
				fmt.Println(grid[i][j]["value"])
			}
		}
	}
}
