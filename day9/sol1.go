package main

import (
	"slices"
	"strconv"
)

func sol1(lines []string) int {

	ret := 0

	var puzzle []rune

	puzzle = append(puzzle, []rune(lines[0])...)

	drive := build_drive(puzzle)
	reformat_drive(drive)

	for i, val := range drive {
		if val != -1 {
			ret += i * val
		}
	}

	return ret

}

func build_drive(puzzle []rune) []int {

	var drive []int

	for x, char := range puzzle {

		i, _ := strconv.Atoi(string(char))

		if x%2 == 0 {

			for range i {
				drive = append(drive, x/2)
			}
		} else {
			for range i {
				drive = append(drive, -1)
			}
		}
	}

	return drive

}

func reformat_drive(drive []int) {

	bad := slices.Index(drive, -1)
	last_num := find_last_num(drive, len(drive)-1)

	// Hate that go does not have a while construct
	for bad < last_num {

		drive[bad] = drive[last_num]
		drive[last_num] = -1

		bad = slices.Index(drive, -1)
		last_num = find_last_num(drive, last_num-1)

	}

}

func find_last_num(drive []int, pos int) int {
	if drive[pos] != -1 {
		return pos
	} else {
		return find_last_num(drive, pos-1)
	}
}
