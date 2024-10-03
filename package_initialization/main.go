package main

import (
	"fmt"
	"golang_basics/package_initialization/database"
)

func main() {
	db := database.GetDatabase()

	fmt.Println(db)
}