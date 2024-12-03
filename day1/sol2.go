package main

import (
	"fmt"
	"sort"
	"strings"
)

func sol2(lines []string) int {

	var first, second []int
	ret := 0

	for _, line := range lines {
		var num1, num2 int
		result := strings.Fields(line)

		_, err := fmt.Sscanf(result[0], "%d", &num1)
		if err != nil {
			continue
		}

		_, err = fmt.Sscanf(result[1], "%d", &num2)
		if err != nil {
			continue
		}

		first = append(first, num1)
		second = append(second, num2)

	}

	sort.Ints(first)
	sort.Ints(second)

	for _, num1 := range first {
		occurences := 0
		for _, num2 := range second {
			if num2 > num1 {
				break
			} else if num1 == num2 {
				occurences++
			} else {
				continue
			}
		}

		ret += num1 * occurences

	}

	return ret

}
