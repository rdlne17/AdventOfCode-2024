package main

import (
	"regexp"
	"strconv"
)

func sol1(lines []string) int {

	ret := 0

	for _, line := range lines {

		matchs := regexp.MustCompile(`mul\((?P<num1>\d+),(?P<num2>\d+)\)`)

		for _, match := range matchs.FindAllStringSubmatch(line, -1) {

			num1, _ := strconv.Atoi(match[matchs.SubexpIndex("num1")])
			num2, _ := strconv.Atoi(match[matchs.SubexpIndex("num2")])

			ret += num1 * num2

		}
	}

	return ret

}
