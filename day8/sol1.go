package main

func sol1(lines []string) int {

	ret := 0

	var puzzle []string

	puzzle = append(puzzle, lines...)

	ret = solve_antennas(puzzle)

	return ret

}

func solve_antennas(puzzle []string) int {

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

		for _, vals := range create_anti_by_antenna(coords_by_antenna[key]) {
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

func create_anti_by_antenna(coords [][]int) [][]int {
	var tmp_antis [][]int

	for i, coord := range coords {

		first_x, first_y := coord[0], coord[1]

		for _, sub_coord := range coords[i+1:] {

			second_x, second_y := sub_coord[0], sub_coord[1]

			x_diff := first_x - second_x
			y_diff := first_y - second_y

			tmp_antis = append(tmp_antis, []int{first_x + x_diff, first_y + y_diff})
			tmp_antis = append(tmp_antis, []int{second_x - x_diff, second_y - y_diff})

		}

	}

	return tmp_antis

}

func matrix_index(sublist [][]int, target []int) bool {

	for _, vals := range sublist {

		if vals[0] == target[0] && vals[1] == target[1] {
			return true
		}

	}

	return false
}
