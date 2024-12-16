package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
)

type location struct {
	x, y int
}

func (v *location) getAllNeighbours() []location {
	return []location{
		{x: v.x - 1, y: v.y},
		{x: v.x + 1, y: v.y},
		{x: v.x, y: v.y - 1},
		{x: v.x, y: v.y + 1},
	}
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

func (v *location) samePlantType(target rune, matrix *[][]rune) bool {
	return (*matrix)[v.y][v.x] == target
}

type line struct {
	loc, formLoc location
	isVertical   bool
}

func getContents() [][]rune {

	file, _ := os.Open("input.txt")
	defer file.Close()

	output := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		output = append(output, row)
	}

	return output
}

func dfs(pos location, plantType rune, visited []location, matrix *[][]rune) []location {
	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1

	visited = append(visited, pos)
	for _, neighbour := range pos.getValidNeighbours(maxX, maxY) {
		if slices.Contains(visited, neighbour) {
			continue
		}
		if !neighbour.samePlantType(plantType, matrix) {
			continue
		}
		newVisited := dfs(neighbour, plantType, visited, matrix)
		visited = newVisited
	}
	return visited

}

func removeElement(e location, arr []location) []location {
	for i := 0; i < len(arr); i++ {
		if arr[i] == e {
			return slices.Delete(arr, i, i+1)
		}
	}
	panic(errors.New("could not find element to remove"))
}

func getRegions(matrix *[][]rune) [][]location {

	maxX, maxY := len((*matrix)[0])-1, len(*matrix)-1
	queue := []location{}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			queue = append(queue, location{x, y})
		}
	}

	regions := [][]location{}

	for len(queue) > 0 {
		loc := queue[0]
		visited := dfs(loc, (*matrix)[loc.y][loc.x], []location{}, matrix)
		regions = append(regions, visited)
		for _, visit := range visited {
			queue = removeElement(visit, queue)
		}
	}

	return regions

}

func determineParameter(region []location) int {

	output := 0
	for _, plant := range region {
		for _, neighbour := range plant.getAllNeighbours() {
			if slices.Contains(region, neighbour) {
				continue
			}
			output++
		}
	}
	return output
}

func part1() {
	matrix := getContents()
	regions := getRegions(&matrix)

	ans := 0
	for _, region := range regions {
		area := len(region)
		parameter := determineParameter(region)
		ans += area * parameter
	}

	fmt.Printf("Part 1 Ans: %v\n", ans)
}

func getFromLocation(original line, i int) location {
	if original.isVertical {
		fromLeft := original.formLoc.x < original.loc.x
		if fromLeft {
			return location{original.loc.x - 1, original.loc.y + i}
		} else {
			return location{original.loc.x + 1, original.loc.y + i}
		}
	} else {
		fromUp := original.formLoc.y < original.loc.y
		if fromUp {
			return location{original.loc.x + i, original.loc.y - 1}
		} else {
			return location{original.loc.x + i, original.loc.y + 1}
		}
	}
}

func determineSides(region []location) int {

	shell := []line{}
	for _, plant := range region {
		// left right down up
		for i, neighbour := range plant.getAllNeighbours() {
			if slices.Contains(region, neighbour) {
				continue
			}
			shell = append(shell, line{neighbour, plant, i <= 1})
		}
	}

	toRemove := []line{}
	for _, s := range shell {
		if slices.Contains(toRemove, s) {
			continue
		}

		if s.isVertical {

			i := 1
			candidate1 := line{location{s.loc.x, s.loc.y - i}, getFromLocation(s, -i), true}
			for slices.Contains(shell, candidate1) {
				toRemove = append(toRemove, candidate1)
				i++
				candidate1 = line{location{s.loc.x, s.loc.y - i}, getFromLocation(s, -i), true}
			}

			i = 1
			candidate2 := line{location{s.loc.x, s.loc.y + i}, getFromLocation(s, i), true}
			for slices.Contains(shell, candidate2) {
				toRemove = append(toRemove, candidate2)
				i++
				candidate2 = line{location{s.loc.x, s.loc.y + i}, getFromLocation(s, i), true}
			}
		} else {

			i := 1
			candidate1 := line{location{s.loc.x - i, s.loc.y}, getFromLocation(s, -i), false}
			for slices.Contains(shell, candidate1) {
				toRemove = append(toRemove, candidate1)
				i++
				candidate1 = line{location{s.loc.x - i, s.loc.y}, getFromLocation(s, -i), false}
			}
			i = 1
			candidate2 := line{location{s.loc.x + i, s.loc.y}, getFromLocation(s, i), false}
			for slices.Contains(shell, candidate2) {
				toRemove = append(toRemove, candidate2)
				i++
				candidate2 = line{location{s.loc.x + i, s.loc.y}, getFromLocation(s, i), false}
			}
		}
	}

	return len(shell) - len(toRemove)
}

func part2() {
	matrix := getContents()
	regions := getRegions(&matrix)

	ans := 0
	for _, region := range regions {
		area := len(region)
		sides := determineSides(region)
		ans += area * sides
	}

	fmt.Printf("Part 2 Ans: %v\n", ans)
}

func main() {
	part1()
	part2()
}
