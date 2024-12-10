package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func compareMaps(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, val := range map1 {
		if val2, ok := map2[key]; !ok || val != val2 {
			return false
		}
	}

	return true
}

func InArrayString(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func InArrayPosition(value map[string]int, array []map[string]int) bool {
	for _, v := range array {
		if compareMaps(v, value) {
			return true
		}
	}
	return false
}

func SearchByValue(m map[string]string, value string) (string, bool) {
	for key, v := range m {
		if v == value {
			return key, true
		}
	}
	return "", false
}

type Guard struct {
	Y         int
	X         int
	OnMap     bool
	Direction string
}

func (guard *Guard) Turn() {
	switch guard.Direction {
	case "up":
		guard.Direction = "right"
	case "right":
		guard.Direction = "down"
	case "down":
		guard.Direction = "left"
	case "left":
		guard.Direction = "up"
	}
}

func NewGuard(grid [][]string) *Guard {
	var Y int
	var X int
	var Direction string

	var directions = map[string]string{
		"up":    "^",
		"right": ">",
		"down":  "v",
		"left":  "<",
	}

	for y, row := range grid {
		for x, content := range row {
			if InArrayString(content, []string{"^", ">", "v", "<"}) {
				direction, _ := SearchByValue(directions, content)

				Y = y
				X = x
				Direction = direction
			}
		}
	}

	return &Guard{Y: Y, X: X, Direction: Direction, OnMap: true}
}

func LoadLinesFromFile(filepath string) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]string

	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return lines
}

func main() {
	grid := LoadLinesFromFile("/data/d6.txt")

	guard := NewGuard(grid)

	var positions = []map[string]int{
		{
			"y": guard.Y,
			"x": guard.X,
		},
	}

	for guard.OnMap {
		var nextPosition = map[string]int{
			"y": 0,
			"x": 0,
		}

		switch guard.Direction {
		case "up":
			nextPosition["y"] = guard.Y - 1
			nextPosition["x"] = guard.X
		case "right":
			nextPosition["y"] = guard.Y
			nextPosition["x"] = guard.X + 1
		case "down":
			nextPosition["y"] = guard.Y + 1
			nextPosition["x"] = guard.X
		case "left":
			nextPosition["y"] = guard.Y
			nextPosition["x"] = guard.X - 1
		}

		inBounds :=
			(guard.Direction == "up" && nextPosition["y"] >= 0) ||
				(guard.Direction == "down" && nextPosition["y"] <= len(grid)-1) ||
				(guard.Direction == "left" && nextPosition["x"] >= 0) ||
				(guard.Direction == "right" && nextPosition["x"] <= len(grid[0])-1)

		if !inBounds {
			guard.OnMap = false
			break
		}

		if grid[nextPosition["y"]][nextPosition["x"]] == "#" {
			guard.Turn()
			continue
		}

		if !InArrayPosition(nextPosition, positions) {
			positions = append(positions, nextPosition)
		}

		guard.Y = nextPosition["y"]
		guard.X = nextPosition["x"]
	}

	result := len(positions)

	fmt.Println("Guard will visit", result, "positions")
}
