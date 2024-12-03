package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fi, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Day 1 solution: %d\n", sol1(lines))
	fmt.Printf("Day 2 solution: %d\n", sol2(lines))

}
