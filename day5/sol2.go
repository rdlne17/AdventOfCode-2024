package main

import (
	"fmt"
	"slices"
	"strings"
)

func sol2(lines []string) int {

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

		bad := is_bad(int_current_line, rules)

		if bad {
			bad_sorted := sort_line(int_current_line, rules)
			ret += bad_sorted[len(bad_sorted)/2]
		}

	}

	return ret

}

func is_bad(int_current_line []int, rules map[int][]int) bool {
	bad := false

	for i, p := range int_current_line {
		// Check if previous numbers are in the current numbers rules
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

	return bad
}

func sort_line(pages []int, rules map[int][]int) []int {

	var tmp_pages []int

	for _, page := range pages {

		if len(tmp_pages) == 0 {
			tmp_pages = append(tmp_pages, page)
			continue
		}

		insert_spot := -1

		for j, tmp_page := range tmp_pages {

			if slices.Index(rules[page], tmp_page) != -1 {
				if insert_spot == -1 {
					insert_spot = j
				}

			} else if slices.Index(rules[tmp_page], page) != -1 {
				insert_spot = j + 1
			}

		}

		tmp_pages = slices.Insert(tmp_pages, insert_spot, page)

	}

	return tmp_pages
}
