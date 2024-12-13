package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const DOT = '.'

type location struct {
	x, y int
}

func (v location) addX(value int) location {
	return location{x: v.x + value, y: v.y}
}

func (v location) addY(value int) location {
	return location{x: v.x, y: v.y + value}
}

func (v location) isValid(maxX, maxY int) bool {
	return v.x >= 0 && v.x <= maxX && v.y >= 0 && v.y <= maxY
}

func (v location) String() string {
	return fmt.Sprintf("location{x: %v, y: %v}", v.x, v.y)
}

func getContents() [][]rune {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func getNextPoint(point location, maxX, maxY int) location {
	if point.x == maxX {
		if point.y == maxY {
			return location{x: point.x + 1, y: point.y + 1}
		}
		return location{x: 0, y: point.y + 1}
	}
	return location{x: point.x + 1, y: point.y}
}

func getAntiNodes(i, j location, maxX, maxY int) []location {

	xDiff := i.x - j.x
	yDiff := i.y - j.y

	output := []location{}

	farFromI := i.addX(-2 * xDiff).addY(-2 * yDiff)
	if farFromI.isValid(maxX, maxY) {
		output = append(output, farFromI)
	}
	behindI := i.addX(xDiff).addY(yDiff)
	if behindI.isValid(maxX, maxY) {
		output = append(output, behindI)
	}

	return output

}

func getLinearAntiNodes(i, j location, maxX, maxY int) []location {
	xDiff := i.x - j.x
	yDiff := i.y - j.y

	output := []location{i, j}

	direction1 := i.addX(xDiff).addY(yDiff)
	for direction1.isValid(maxX, maxY) {
		output = append(output, direction1)
		direction1 = direction1.addX(xDiff).addY(yDiff)
	}

	direction2 := j.addX(-xDiff).addY(-yDiff)
	for direction2.isValid(maxX, maxY) {
		output = append(output, direction2)
		direction2 = direction2.addX(-xDiff).addY(-yDiff)
	}

	return output
}

func main() {
	matrix := getContents()

	maxX, maxY := len(matrix[0])-1, len(matrix)-1

	part1Visited := []location{}
	part2Visited := []location{}

	for firstY := 0; firstY <= maxY; firstY++ {
		for firstX := 0; firstX <= maxY; firstX++ {

			firstFreq := matrix[firstY][firstX]
			firstPoint := location{x: firstX, y: firstY}
			nextPoint := getNextPoint(firstPoint, maxX, maxY)

			// fmt.Println("first: ", firstPoint)
			for secondY := nextPoint.y; secondY <= maxY; secondY++ {
				if secondY != nextPoint.y {
					nextPoint.x = 0
				}
				for secondX := nextPoint.x; secondX <= maxX; secondX++ {
					secondPoint := location{x: secondX, y: secondY}
					// fmt.Println("second: ", secondPoint)
					secondFreq := matrix[secondY][secondX]

					if firstFreq != secondFreq || DOT == firstFreq || DOT == secondFreq {
						continue
					}

					part1AntiNodes := getAntiNodes(firstPoint, secondPoint, maxX, maxY)
					for _, antiNode := range part1AntiNodes {
						if slices.Contains(part1Visited, antiNode) {
							continue
						}
						part1Visited = append(part1Visited, antiNode)

					}

					part2AntiNodes := getLinearAntiNodes(firstPoint, secondPoint, maxX, maxY)
					for _, antiNode := range part2AntiNodes {
						if slices.Contains(part2Visited, antiNode) {
							continue
						}
						part2Visited = append(part2Visited, antiNode)

					}

				}
			}
		}
	}

	fmt.Printf("Part 1 Ans: %v\n", len(part1Visited))
	fmt.Printf("Part 1 Ans: %v\n", len(part2Visited))

}
