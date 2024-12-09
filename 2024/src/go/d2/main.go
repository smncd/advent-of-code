package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number interface {
	int | float64
}

func absInt(x int) int {
	if x < 0 {
		return 0 - x
	}

	return x - 0
}

func isSafeRange[T Number](number T) bool {
	return number <= 3 && number >= 1
}

func loadLevelsListFromFile(filepath string) [][]int {

	var levelsList [][]int

	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		var levels []int

		for _, level := range strings.Split(line, " ") {
			level, err := strconv.Atoi(level)
			if err != nil {
				continue
			}

			levels = append(levels, level)
		}

		levelsList = append(levelsList, levels)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return levelsList
}

func processLevels(levels []int) bool {
	isSafe := false

	isIncreasing := true
	isDecreasing := true

	for index, level := range levels {

		var next int

		if index+1 < len(levels) {
			next = levels[index+1]
		} else {
			continue
		}

		difference := next - level

		if difference > 0 {
			isIncreasing = isIncreasing && isSafeRange(difference)
			isDecreasing = false
		} else if difference < 0 {
			isDecreasing = isDecreasing && isSafeRange(absInt(difference))
			isIncreasing = false
		} else {
			isIncreasing = false
			isDecreasing = false
		}

		isSafe = isIncreasing || isDecreasing
	}

	return isSafe
}

func main() {
	levelsList := loadLevelsListFromFile("/data/d2.txt")

	safeReports := 0
	safeReportsWithDampener := 0

	for _, levels := range levelsList {
		isSafe := processLevels(levels)

		if isSafe {
			safeReports += 1
		} else {
			for index := range levels {
				levelsToCheck := append([]int(nil), levels...)

				if index >= 0 && index < len(levelsToCheck) {
					levelsToCheck = append(levelsToCheck[:index], levelsToCheck[index+1:]...)
				}

				if processLevels(levelsToCheck) {
					safeReportsWithDampener += 1
					break
				}
			}
		}
	}

	total := safeReports + safeReportsWithDampener

	fmt.Println("Number of safe reports:", safeReports)
	fmt.Println("Number of safe reports (with Problem Dampener):", safeReportsWithDampener)
	fmt.Println("Total:", total)
}
