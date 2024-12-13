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

	ret_slice := solve_puzzle2(puzzle_map, current_x, current_y, 0)

	slices.SortFunc(ret_slice, matrix_sort)

	ret_slice = slices.CompactFunc(ret_slice, matrix_eq)

	ret = detect_potential_loops(ret_slice)

	return ret

}

func solve_puzzle2(puzzle_map map[int][]rune, current_x, current_y, num_rotations int) [][]int {

	tmp_coords := [][]int{}

	// Base case of out of puzzle
	if current_x > len(puzzle_map)-1 || current_x < 0 || current_y > len(puzzle_map[current_y])-1 || current_y < 0 {
		return tmp_coords
	}

	// Hit obsticle so rotate and try again
	if puzzle_map[current_x][current_y] == '#' {

		switch num_rotations % 4 { // Undo last move
		case 0: // Going up
			current_x += 1

		case 1: // Going right
			current_y -= 1

		case 2: // Going down
			current_x -= 1

		default: // Going left
			current_y += 1
		}

		return slices.Concat(tmp_coords, solve_puzzle2(puzzle_map, current_x, current_y, num_rotations+1))
	}

	tmp_coords = [][]int{{current_x, current_y}}

	switch num_rotations % 4 {
	case 0: // Going up
		tmp_coords = slices.Concat(tmp_coords, solve_puzzle2(puzzle_map, current_x-1, current_y, num_rotations))

	case 1: // Going right
		tmp_coords = slices.Concat(tmp_coords, solve_puzzle2(puzzle_map, current_x, current_y+1, num_rotations))

	case 2: // Going down
		tmp_coords = slices.Concat(tmp_coords, solve_puzzle2(puzzle_map, current_x+1, current_y, num_rotations))

	default: // Going left
		tmp_coords = slices.Concat(tmp_coords, solve_puzzle2(puzzle_map, current_x, current_y-1, num_rotations))
	}

	return tmp_coords

}

func detect_potential_loops(visited [][]int) int {
	ret := 0

	return ret
}
