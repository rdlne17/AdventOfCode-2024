package main

import (
	"strings"
)

func sol1(lines []string) int {

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
			ret += solve_crossword1(matrix, i, j)
		}
	}

	return ret

}

func solve_crossword1(list [][]string, x, y int) int {

	ret := 0

	if check_spot(list, x, y, 0) { // Checks if current spot is X if so check all directions

		if check_spot(list, x+1, y, 1) && check_spot(list, x+2, y, 2) && check_spot(list, x+3, y, 3) { // bottom
			ret += 1
		}

		if check_spot(list, x-1, y, 1) && check_spot(list, x-2, y, 2) && check_spot(list, x-3, y, 3) { // top
			ret += 1
		}

		if check_spot(list, x+1, y+1, 1) && check_spot(list, x+2, y+2, 2) && check_spot(list, x+3, y+3, 3) { // bottom right
			ret += 1
		}

		if check_spot(list, x-1, y+1, 1) && check_spot(list, x-2, y+2, 2) && check_spot(list, x-3, y+3, 3) { // top right
			ret += 1
		}

		if check_spot(list, x+1, y-1, 1) && check_spot(list, x+2, y-2, 2) && check_spot(list, x+3, y-3, 3) { // bottom left
			ret += 1
		}

		if check_spot(list, x-1, y-1, 1) && check_spot(list, x-2, y-2, 2) && check_spot(list, x-3, y-3, 3) { // top left
			ret += 1
		}

		if check_spot(list, x, y-1, 1) && check_spot(list, x, y-2, 2) && check_spot(list, x, y-3, 3) { // left
			ret += 1
		}

		if check_spot(list, x, y+1, 1) && check_spot(list, x, y+2, 2) && check_spot(list, x, y+3, 3) { // right
			ret += 1
		}

	}
	return ret

}

func check_spot(list [][]string, x, y, current_num int) bool {
	if x < 0 || x >= len(list) || y < 0 || y >= len(list[x]) {
		return false
	}

	switch current_num {
	case 0:
		if list[x][y] == "X" {
			return true
		}
	case 1:
		if list[x][y] == "M" {
			return true
		}
	case 2:
		if list[x][y] == "A" {
			return true
		}
	case 3:
		if list[x][y] == "S" {
			return true
		}
	}
	return false
}
