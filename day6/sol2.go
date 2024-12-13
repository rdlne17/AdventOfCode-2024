package main

import (
	"slices"
)

func sol2(lines []string) int {

	ret := 0

	puzzle_map := make(map[int][]rune, len(lines))
	var current_x, current_y int

	for i, line := range lines {

		puzzle_map[i] = []rune(line)

		for j, c := range puzzle_map[i] {

			if c == '^' {
				current_x = i
				current_y = j
			}

		}

	}

	ret_slice := solve_puzzle(puzzle_map, current_x, current_y, 0)

	slices.SortFunc(ret_slice, matrix_sort)

	ret_slice = slices.CompactFunc(ret_slice, matrix_eq)

	ret = detect_potential_loops(ret_slice)

	return ret

}

func detect_potential_loops(visited [][]int) int {
	ret := 0

	return ret
}
