package main

import (
	"fmt"
	CoreFunctions "golang_basics/mini_project_1/core_functions"
)

func displayGrid(grid CoreFunctions.Grid) {

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
				}
			} else {
				fmt.Printf("   *  ")
			}
		}
		fmt.Println()
	}
}

func fieldData(grid CoreFunctions.Grid) {

	fmt.Printf("\n")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("=====================| Field Data |=====================")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("\n")

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j]["value"] != "" {
				fmt.Printf("(X%s, Y%s): ", grid[i][j]["coord_x"], grid[i][j]["coord_y"])
				fmt.Println(grid[i][j]["value"])
			}
		}
	}
}

func main() {
	var grid_personel_role = CoreFunctions.GenerateGrid()
	var grid_personel_full = CoreFunctions.GenerateGrid()
	var personels = CoreFunctions.GetPersonels()

	var coord_x = 0
	var coord_y = 0

	for i := 0; i < len(personels); i++ {
		if coord_x < 5 && coord_y < 5 {
			grid_personel_role[coord_x][coord_y]["value"] = personels[i]["role"]
			coord_x++
			coord_y++
		}
	}

	coord_x = 0
	coord_y = 0

	for i := 0; i < len(personels); i++ {
		if coord_x < 4 && coord_y < 4 {
			grid_personel_full[coord_x][coord_y]["value"] = personels[i]["role"] + ", " + personels[i]["rank"] + ", " + personels[i]["alias"] + ", " + personels[i]["name"]
			coord_x++
			coord_y++
		}
	}

	displayGrid(grid_personel_role)
	fieldData(grid_personel_full)

}
