package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getFileContents(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Print(err)
	}

	return string(file)
}

func main() {
	file := getFileContents("/data/d3.txt")

	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	matches := re.FindAllStringSubmatch(file, -1)

	var firsts []int
	var seconds []int

	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		firsts = append(firsts, first)
		seconds = append(seconds, second)
	}

	total := 0

	for index, first := range firsts {
		second := seconds[index]

		total += first * second
	}

	fmt.Println("Total:", total)
}
