package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	GUARD    = '^'
	OBSTACLE = '#'
)

type location struct {
	x, y int
}

func (v *location) isValid(maxX, maxY int) bool {
	return v.x >= 0 && v.x <= maxX && v.y >= 0 && v.y <= maxY
}

type playerState struct {
	loc  location
	face rune
}

func (v *playerState) nextLocation() location {
	switch v.face {
	case 'N':
		return location{v.loc.x, v.loc.y - 1}
	case 'E':
		return location{v.loc.x + 1, v.loc.y}
	case 'S':
		return location{v.loc.x, v.loc.y + 1}
	case 'W':
		return location{v.loc.x - 1, v.loc.y}
	default:
		panic(errors.New("Invalid face"))
	}
}

func (v *playerState) moveAhead() {
	v.loc = v.nextLocation()
}

func (v *playerState) turnRight() {
	rotate := map[rune]rune{'N': 'E', 'E': 'S', 'S': 'W', 'W': 'N'}
	v.face = rotate[v.face]
}

func getContents() [][]rune {
	file, _ := os.Open("input.txt")
	defer file.Close()

	matrix := [][]rune{}
	scanner := bufio.NewScanner(file)
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

func countVisited(matrix *[][]rune) int {
	maxX, maxY := len((*matrix)[0])-1, len((*matrix))-1

	output := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if (*matrix)[y][x] == 'X' {
				output++
			}
		}
	}
	return output
}

func getplayerState(matrix *[][]rune) playerState {

	maxX, maxY := len((*matrix)[0])-1, len((*matrix))-1
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if GUARD == (*matrix)[y][x] {
				return playerState{loc: location{x: x, y: y}, face: 'N'}
			}
		}
	}
	panic(errors.New("playerState not found"))
}

func simulate(markVisitedPositions bool, matrix *[][]rune) bool {

	maxX, maxY := len((*matrix)[0])-1, len((*matrix))-1

	player := getplayerState(matrix)
	visited := make(map[playerState]bool)

	if markVisitedPositions {
		(*matrix)[player.loc.y][player.loc.x] = 'X'
	}

	for {

		// check if current location is valid
		currLoc := player.loc
		if !currLoc.isValid(maxX, maxY) {
			break
		}

		// mark current position
		if markVisitedPositions {
			(*matrix)[currLoc.y][currLoc.x] = 'X'
		}

		// updated and check for visited state
		if visited[player] {
			return true
		}
		visited[player] = true

		// check for obstacle to rotate: edge case double rotate
		newLoc := player.nextLocation()
		if newLoc.isValid(maxX, maxY) {
			if OBSTACLE == (*matrix)[newLoc.y][newLoc.x] {
				player.turnRight()
				newNewLoc := player.nextLocation()
				if OBSTACLE == (*matrix)[newNewLoc.y][newNewLoc.x] {
					player.turnRight()
				}
			}
		}
		player.moveAhead()
	}
	return false
}

func rotateMatrix(matrix *[][]rune) [][]rune {
	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1
	player := getplayerState(matrix)

	rotatedMatrix := [][]rune{}
	for y := 0; y <= maxY; y++ {
		tempRow := []rune{}
		for x := 0; x <= maxX; x++ {
			tempRow = append(tempRow, 0)
		}
		rotatedMatrix = append(rotatedMatrix, tempRow)
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			rotatedMatrix[x][maxY-y] = (*matrix)[y][x]
		}
	}

	rotatedMatrix[player.loc.y][player.loc.x] = GUARD

	return rotatedMatrix
}

func part1() {

	matrix := getContents()

	simulate(true, &matrix)

	fmt.Printf("Part 1 Ans: %v\n", countVisited(&matrix))
}

func part2() {

	matrix := getContents()
	maxX, maxY := len(matrix[0])-1, len(matrix)-1

	ans := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if matrix[y][x] != '.' {
				continue
			}

			matrix[y][x] = '#'
			stuck := simulate(false, &matrix)
			if stuck {
				ans++
			}
			matrix[y][x] = '.'
		}
	}
	fmt.Printf("Part 2 Ans: %v\n", ans)

}

func main() {
	part1()
	part2()
}
