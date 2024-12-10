package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LoadLinesFromFile(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func FormatData(file []string, pattern string, separator string) ([][]int, error) {
	re := regexp.MustCompile(pattern)

	var filtered [][]int

	for _, line := range file {
		if re.MatchString(line) {
			stringNumbers := strings.Split(line, separator)

			var intNumbers []int

			for _, number := range stringNumbers {
				num, err := strconv.Atoi(number)
				if err != nil {
					continue
				}
				intNumbers = append(intNumbers, num)
			}

			filtered = append(filtered, intNumbers)
		}
	}

	return filtered, nil
}

func InArray(value int, array []int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func ArraySearch(value int, array []int) (int, error) {
	for i, v := range array {
		if v == value {
			return i, nil
		}
	}
	return -1, errors.New("Value not found in array")
}

func main() {
	file := LoadLinesFromFile("/data/d5.txt")

	pageOrderingRules, err := FormatData(file, `([0-9]{2})\|([0-9]{2})`, "|")

	if err != nil {
		panic(err)
	}

	pagesToProduce, err := FormatData(file, `(^[0-9,]*$)`, ",")

	if err != nil {
		panic(err)
	}

	var isInRightOrder [][]int

	for _, pages := range pagesToProduce {
		var correctOrder bool

		for range pages {
			for _, rule := range pageOrderingRules {
				if InArray(rule[0], pages) && InArray(rule[1], pages) {
					ruleZeroIndex, ruleZeroErr := ArraySearch(rule[0], pages)
					ruleOneIndex, ruleOneErr := ArraySearch(rule[1], pages)

					if ruleZeroErr != nil || ruleOneErr != nil || ruleOneIndex <= ruleZeroIndex {
						correctOrder = false
						break
					}
				}

				correctOrder = true
			}
		}

		if correctOrder {
			isInRightOrder = append(isInRightOrder, pages)
		}
	}

	result := 0

	for _, pages := range isInRightOrder {
		result += pages[(len(pages)-1)/2]
	}

	fmt.Println(result)
}
