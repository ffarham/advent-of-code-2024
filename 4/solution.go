package main

import (
	"bufio"
	"fmt"
	"os"
)

func getContents() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	data := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		lineArr := []string{}
		for _, char := range line {
			lineArr = append(lineArr, string(char))
		}
		data = append(data, lineArr)
	}

	return data
}

func findXmasAtIndex(xIndex, yIndex int, matrix *[][]string) int {
	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1

	output := 0

	// look right
	if xIndex < maxX-2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex][xIndex+1] + (*matrix)[yIndex][xIndex+2] + (*matrix)[yIndex][xIndex+3]
		if candidate == "XMAS" {
			output++
		}
	}

	// look left
	if xIndex > 2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex][xIndex-1] + (*matrix)[yIndex][xIndex-2] + (*matrix)[yIndex][xIndex-3]
		if candidate == "XMAS" {
			output++
		}
	}

	// look up
	if yIndex > 2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex-1][xIndex] + (*matrix)[yIndex-2][xIndex] + (*matrix)[yIndex-3][xIndex]
		if candidate == "XMAS" {
			output++
		}
	}

	// look down
	if yIndex < maxY-2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex+1][xIndex] + (*matrix)[yIndex+2][xIndex] + (*matrix)[yIndex+3][xIndex]
		if candidate == "XMAS" {
			output++
		}
	}

	// look top-left
	if yIndex > 2 && xIndex > 2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex-1][xIndex-1] + (*matrix)[yIndex-2][xIndex-2] + (*matrix)[yIndex-3][xIndex-3]
		if candidate == "XMAS" {
			output++
		}
	}

	// look top-right
	if yIndex > 2 && xIndex < maxX-2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex-1][xIndex+1] + (*matrix)[yIndex-2][xIndex+2] + (*matrix)[yIndex-3][xIndex+3]
		if candidate == "XMAS" {
			output++
		}
	}

	// look bottom-right
	if yIndex < maxY-2 && xIndex < maxX-2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex+1][xIndex+1] + (*matrix)[yIndex+2][xIndex+2] + (*matrix)[yIndex+3][xIndex+3]
		if candidate == "XMAS" {
			output++
		}
	}

	// look bottom-left
	if yIndex < maxY-2 && xIndex > 2 {
		candidate := (*matrix)[yIndex][xIndex] + (*matrix)[yIndex+1][xIndex-1] + (*matrix)[yIndex+2][xIndex-2] + (*matrix)[yIndex+3][xIndex-3]
		if candidate == "XMAS" {
			output++
		}
	}

	return output

}

func findMasAtIndex(xIndex, yIndex int, matrix *[][]string) bool {
	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1

	// if center isnt A, can never be X-MAS
	if (*matrix)[yIndex][xIndex] != "A" {
		return false
	}

	// if index on boundary, cant form X
	if xIndex == 0 || xIndex == maxX || yIndex == 0 || yIndex == maxY {
		return false
	}

	// check y=x diagonal
	if (((*matrix)[yIndex+1][xIndex-1]+(*matrix)[yIndex][xIndex]+(*matrix)[yIndex-1][xIndex+1] == "MAS") ||
		((*matrix)[yIndex-1][xIndex+1]+(*matrix)[yIndex][xIndex]+(*matrix)[yIndex+1][xIndex-1] == "MAS")) &&
		(((*matrix)[yIndex-1][xIndex-1]+(*matrix)[yIndex][xIndex]+(*matrix)[yIndex+1][xIndex+1] == "MAS") ||
			((*matrix)[yIndex+1][xIndex+1]+(*matrix)[yIndex][xIndex]+(*matrix)[yIndex-1][xIndex-1] == "MAS")) {
		return true
	}

	return false

}

func part1(matrix *[][]string) {

	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1

	ans := 0

	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			ans += findXmasAtIndex(j, i, matrix)
		}
	}

	fmt.Printf("Part 1 Ans: %v\n", ans)
}

func part2(matrix *[][]string) {
	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1

	ans := 0

	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if findMasAtIndex(j, i, matrix) {
				ans += 1
			}
		}
	}

	fmt.Printf("Part 2 Ans: %v\n", ans)
}

func main() {
	fmt.Println("Advent of code day 4 solution")

	matrix := getContents()

	part1(&matrix)
	part2(&matrix)

}
