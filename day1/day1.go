package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func day1(lines []string) int {

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

	for i, num := range first {
		ret += int(math.Abs(float64(second[i] - num)))
	}

	return ret

}
