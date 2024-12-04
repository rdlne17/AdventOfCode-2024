package main

import (
	"regexp"
	"strconv"
	"strings"
)

func sol2(lines []string) int {

	ret := 0

	line := strings.Join(lines[:], "")

	bad := regexp.MustCompile(`don't\(\).*?do\(\)`)

	replaced := bad.ReplaceAllString(line, "")

	matchs := regexp.MustCompile(`mul\((?P<num1>\d+),(?P<num2>\d+)\)`)

	for _, match := range matchs.FindAllStringSubmatch(replaced, -1) {

		num1, _ := strconv.Atoi(match[matchs.SubexpIndex("num1")])
		num2, _ := strconv.Atoi(match[matchs.SubexpIndex("num2")])

		ret += num1 * num2

	}

	return ret

}
