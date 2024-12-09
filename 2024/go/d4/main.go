package main

import (
	"bufio"
	"fmt"
	"os"
)

func strSplit(s string) []string {
	runes := []rune(s)
	result := make([]string, len(runes))

	for i, r := range runes {
		result[i] = string(r)
	}

	return result
}

func loadGridFromFile(filepath string) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string

	for scanner.Scan() {
		grid = append(grid, strSplit(scanner.Text()))
	}

	return grid
}

func countWord(grid [][]string, word string) int {
	count := 0

	rows := len(grid)
	cols := len(grid[0])

	directions := [][]int{
		{0, 1},   // Horizontal right
		{0, -1},  // Horizontal left
		{1, 0},   // Vertical down
		{-1, 0},  // Vertical up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}

	for y, row := range grid {
		for x := range row {
			for _, direction := range directions {
				dy, dx := direction[0], direction[1]

				result := ""

				for index := range strSplit(word) {
					ny := y + dy*index
					nx := x + dx*index

					if nx < 0 || ny < 0 || nx >= cols || ny >= rows {
						continue
					}

					result = result + grid[ny][nx]
				}

				if result == word {
					count += 1
				}
			}
		}
	}

	return count
}

func main() {
	grid := loadGridFromFile("/data/d4.txt")

	count := countWord(grid, "XMAS")

	fmt.Println("Total occurances of 'XMAS':", count)
}
