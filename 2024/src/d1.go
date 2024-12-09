package main

import (
	"fmt" 
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

type Lists struct {
    Right []int
    Left  []int
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func countOccurrences(slice []int, condition func(int) bool) int {
	count := 0
	for _, v := range slice {
		if condition(v) {
			count++
		}
	}
	return count
}

func main() {
	lists := Lists{
        Right: []int{},
        Left:  []int{},
    }
	
	file, err := os.Open("/data/d1.txt")
	
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()

		ids := strings.Split(line, "   ")

		if len(ids) > 1 {
			left, err := strconv.Atoi(ids[0])
			right, err := strconv.Atoi(ids[1])
			if err != nil {
				panic(err)
			}

			lists.Left = append(lists.Left, left)
			lists.Right = append(lists.Right, right)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(lists.Left[:])
	sort.Ints(lists.Right[:])

	totalDistance := 0
	similarityScore:= 0

	for index, value := range lists.Left {
		distance := absInt(value - lists.Right[index])

		timesInRightList := countOccurrences(lists.Right, func(id int) bool {
			return id == value
		})

		totalDistance += distance
		similarityScore += value * timesInRightList
	}

	fmt.Println("Total distance:", totalDistance)
	fmt.Println("Similarity score:", similarityScore)
}
