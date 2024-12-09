package main

func sol2(lines []string) int {

	ret := 0

	var puzzle []string

	puzzle = append(puzzle, lines...)

	ret = solve_antennas2(puzzle)

	return ret

}

func solve_antennas2(puzzle []string) int {

	ret := 0

	coords_by_antenna := make(map[rune][][]int)

	for x, line := range puzzle {
		for y, char := range line {
			if char != '.' {
				coords_by_antenna[char] = append(coords_by_antenna[char], []int{x, y})
			}
		}
	}

	max_x := len(puzzle)
	max_y := len(puzzle[0])

	var retvals [][]int

	for key := range coords_by_antenna {

		for _, vals := range create_anti_by_antenna2(coords_by_antenna[key], max_x, max_y) {
			// fmt.Println(vals)
			if vals[0] < 0 || vals[1] < 0 || vals[0] >= max_x || vals[1] >= max_y {
				continue
			}
			contains_coord := matrix_index(retvals, []int{vals[0], vals[1]})
			if contains_coord {
				continue

			}
			retvals = append(retvals, vals)
		}
	}

	ret = len(retvals)

	return ret

}

func create_anti_by_antenna2(coords [][]int, max_x, max_y int) [][]int {
	var tmp_antis [][]int

	for i, coord := range coords {

		first_x, first_y := coord[0], coord[1]

		for _, sub_coord := range coords[i+1:] {

			second_x, second_y := sub_coord[0], sub_coord[1]

			x_diff := first_x - second_x
			y_diff := first_y - second_y

			tmp_antis = append(tmp_antis, create_anti_recurse(first_x, first_y, x_diff, y_diff, max_x, max_y)...)         // Go in one directino based on first antenna
			tmp_antis = append(tmp_antis, create_anti_recurse(second_x, second_y, -1*x_diff, -1*y_diff, max_x, max_y)...) // Go in other direction based on second antenna

		}
		tmp_antis = append(tmp_antis, coord) // Add the antenna cords since they will always be part of the anti's according to the rules

	}

	return tmp_antis

}

func create_anti_recurse(val1, val2, val1_diff, val2_diff, val1_max, val2_max int) [][]int {

	var tmp_antis [][]int

	if val1 < 0 || val2 < 0 || val1 >= val1_max || val2 >= val2_max {
		return tmp_antis
	}

	tmp_antis = append(tmp_antis, []int{val1 + val1_diff, val2 + val2_diff})

	tmp_antis = append(tmp_antis, create_anti_recurse(val1+val1_diff, val2+val2_diff, val1_diff, val2_diff, val1_max, val2_max)...)

	return tmp_antis
}
