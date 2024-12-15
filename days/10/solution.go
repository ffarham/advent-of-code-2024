package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type location struct {
	x, y int
}

func (v *location) getValue(matrix *[][]int) int {
	return (*matrix)[v.y][v.x]
}

func (v *location) getValidNeighbours(maxX, maxY int) []location {
	output := []location{}
	if v.x > 0 {
		output = append(output, location{x: v.x - 1, y: v.y})
	}
	if v.x < maxX {
		output = append(output, location{x: v.x + 1, y: v.y})
	}
	if v.y > 0 {
		output = append(output, location{x: v.x, y: v.y - 1})
	}
	if v.y < maxY {
		output = append(output, location{x: v.x, y: v.y + 1})
	}
	return output
}

func (v *location) isValidNextStep(prev location, matrix *[][]int) bool {
	return (*matrix)[v.y][v.x] == (*matrix)[prev.y][prev.x]+1
}

func (v *location) nextSteps(maxX, maxY int, matrix *[][]int) []location {
	output := []location{}
	neighbours := v.getValidNeighbours(maxX, maxY)
	for _, neighbour := range neighbours {
		if !neighbour.isValidNextStep(*v, matrix) {
			continue
		}
		output = append(output, neighbour)
	}
	return output
}

func getContents() [][]int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	matrix := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []int{}
		for _, char := range line {
			if char == '.' {
				row = append(row, -1)
			} else {
				row = append(row, int(char-'0'))
			}
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func getStartingPositions(matrix *[][]int) []location {
	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1
	output := []location{}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if (*matrix)[y][x] == 0 {
				output = append(output, location{x: x, y: y})
			}
		}
	}
	return output
}

func getTrails(currPos location, matrix *[][]int, visited []location) (int, []location) {

	if (*matrix)[currPos.y][currPos.x] == 9 {
		if slices.Contains(visited, currPos) {
			return 0, visited
		}
		visited = append(visited, currPos)
		return 1, visited
	}

	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1
	output := 0
	for _, newPos := range currPos.nextSteps(maxX, maxY, matrix) {
		score, newVisited := getTrails(newPos, matrix, visited)
		output += score
		visited = newVisited
	}
	return output, visited
}

func getDistinctTrails(currPos location, matrix *[][]int) int {
	if (*matrix)[currPos.y][currPos.x] == 9 {
		return 1
	}
	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1
	output := 0
	for _, newPos := range currPos.nextSteps(maxX, maxY, matrix) {
		output += getDistinctTrails(newPos, matrix)
	}
	return output
}

func part1() {
	matrix := getContents()

	ans := 0
	for _, loc := range getStartingPositions(&matrix) {
		visited := []location{}
		score, _ := getTrails(loc, &matrix, visited)
		ans += score
	}

	fmt.Printf("Part 1 Ans: %v\n", ans)
}

func part2() {
	matrix := getContents()

	ans := 0
	for _, loc := range getStartingPositions(&matrix) {
		ans += getDistinctTrails(loc, &matrix)
	}

	fmt.Printf("Part 2 Ans: %v\n", ans)
}

func main() {

	part1()
	part2()

}
