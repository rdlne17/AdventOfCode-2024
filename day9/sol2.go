package main

import (
	"slices"
)

func sol2(lines []string) int {

	ret := 0

	var puzzle []rune

	puzzle = append(puzzle, []rune(lines[0])...)

	drive := build_drive(puzzle)
	reformat_drive2(drive)

	for i, val := range drive {
		if val != -1 {
			ret += i * val
		}
	}

	return ret

}

func reformat_drive2(drive []int) {

	var bad int
	var bad_len int

	last_num_first_pos := find_last_num(drive, len(drive)-1)
	len_last := len_last_num(drive, drive[last_num_first_pos], last_num_first_pos)
	last_num_first_pos -= len_last - 1

	// Hate that go does not have a while construct
	for last_num_first_pos > 0 {

		bad = slices.Index(drive, -1)

		for bad != -1 && bad < last_num_first_pos {

			bad_len = len_bad(drive, bad)

			if bad_len >= len_last {

				for i := range len_last {
					drive[i+bad] = drive[last_num_first_pos+i]
					drive[last_num_first_pos+i] = -1
				}

				break

			} else {

				bad += 1

			}

		}

		last_num_first_pos = find_last_num(drive, last_num_first_pos-1)
		len_last = len_last_num(drive, drive[last_num_first_pos], last_num_first_pos)
		last_num_first_pos -= (len_last - 1)
	}

}

func len_last_num(drive []int, target, pos int) int {
	if pos < 0 || drive[pos] != target {
		return 0
	}

	return 1 + len_last_num(drive, target, pos-1)
}

func len_bad(drive []int, pos int) int {
	if drive[pos] != -1 {
		return 0
	}

	return 1 + len_bad(drive, pos+1)
}
