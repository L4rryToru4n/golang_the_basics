package main

import (
	"fmt"
	CoreFunctions "golang_basics/mini_project_1/core_functions"
)

func main() {
	var grid = CoreFunctions.GenerateGrid()
	var personels = CoreFunctions.GetPersonels()

	fmt.Println(grid)
	fmt.Println(personels)
}
