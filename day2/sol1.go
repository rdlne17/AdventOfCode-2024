package main

import (
	"fmt"
	"strings"
)

func sol1(lines []string) int {

	ret := 0

	for _, line := range lines {
		var num float64
		last := float64(-1)
		result := strings.Fields(line)
		inc, dec := true, true

		for _, current := range result {
			_, err := fmt.Sscanf(current, "%f", &num)
			if err != nil {
				continue
			}

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

			if !inc && !dec {
				break
			}

			last = num

		}

		if inc || dec {
			ret += 1
		}

	}

	return ret

}

func is_dec(num1, num2 float64) bool {
	retval := num1 - num2
	return (retval <= 3 && retval > 0)
}

func is_inc(num1, num2 float64) bool {
	retval := num2 - num1
	return (retval <= 3 && retval > 0)
}
