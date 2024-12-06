package main

import (
	"fmt"
	"strings"
)

func sol1(lines []string) int {

	ret := 0

	rules := make(map[int][]int)
	var page, subpage, pages int

	for i, line := range lines {

		if line == "" {
			pages = i
			break
		}

		_, err := fmt.Sscanf(line, "%d|%d", &page, &subpage)
		if err != nil {
			continue
		}

		rules[page] = append(rules[page], subpage)

	}

	for _, current_line := range lines[pages+1:] {

		var int_current_line []int

		for _, num := range strings.Split(current_line, ",") {
			_, err := fmt.Sscanf(num, "%d", &page)
			if err != nil {
				continue
			}
			int_current_line = append(int_current_line, page)
		}

		bad := false

		for i, p := range int_current_line {

			for _, previous_page := range int_current_line[:i] {
				for _, bad_pages := range rules[p] {
					if bad_pages == previous_page {
						bad = true
						break
					}
				}
				if bad {
					break
				}

			}

			if bad {
				break
			}

		}

		if !bad {

			ret += int_current_line[len(int_current_line)/2]
		}

	}

	return ret

}
