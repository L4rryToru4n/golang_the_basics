package CoreFunctions

import "strconv"

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
