package main

import (
	"strings"
)

func sol2(lines []string) int {

	ret := 0

	matrix := make([][]string, len(lines))
	for i := range matrix {
		matrix[i] = make([]string, len(lines[i]))
	}

	for i, line := range lines {

		for j, letter := range strings.Split(line, "") {
			matrix[i][j] = letter
		}
	}

	for i, line := range matrix {
		for j := range line {
			ret += solve_crossword2(matrix, i, j)
		}
	}

	return ret

}

func solve_crossword2(list [][]string, x, y int) int {

	ret := 0

	if check_spot(list, x, y, 1) { // Checks if current spot is X if so check all directions

		if check_spot(list, x+2, y, 1) && check_spot(list, x+1, y+1, 2) && check_spot(list, x, y+2, 3) && check_spot(list, x+2, y+2, 3) { // M on left
			ret += 1
		}

		if check_spot(list, x, y+2, 1) && check_spot(list, x+1, y+1, 2) && check_spot(list, x+2, y, 3) && check_spot(list, x+2, y+2, 3) { // M on top
			ret += 1
		}

		if check_spot(list, x+2, y, 1) && check_spot(list, x+1, y-1, 2) && check_spot(list, x+2, y-2, 3) && check_spot(list, x, y-2, 3) { // M on right
			ret += 1
		}

		if check_spot(list, x, y+2, 1) && check_spot(list, x-1, y+1, 2) && check_spot(list, x-2, y, 3) && check_spot(list, x-2, y+2, 3) { // M on bottom
			ret += 1
		}

	}
	return ret

}
