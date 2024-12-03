package main

import (
	"fmt"
	"strings"
)

func sol2(lines []string) int {

	ret := 0
	var num float64

	for _, line := range lines {
		var nums []float64

		result := strings.Fields(line)
		for _, current := range result {
			_, err := fmt.Sscanf(current, "%f", &num)
			if err != nil {
				continue
			}

			nums = append(nums, num)
		}

		safe := check_safe(nums)

		if safe {
			ret += 1
		}

	}

	return ret

}

func check_safe(nums []float64) bool {

	last := float64(-1)
	var dec, inc bool

	l := len(nums)

	for i := range l {

		dec = true
		inc = true
		last = -1

		numsCopy := remove_elem(nums, i)

		for _, num := range numsCopy {
			// skip over first
			if last == -1 {
				last = num
				continue
			}

			if dec {
				dec = is_dec(last, num)
			}

			if inc {
				inc = is_inc(last, num)
			}

			if !dec && !inc {
				break
			}

			last = num

		}

		if dec || inc {
			return dec || inc
		}
	}

	return dec || inc

}

func remove_elem(l []float64, i int) []float64 {

	var lCopy []float64

	lCopy = append(lCopy, l[:i]...)
	lCopy = append(lCopy, l[i+1:]...)

	return lCopy
}
