package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	maxX = 101
	maxY = 103
)

type historian struct {
	p location
	v velocity
}

func (v *historian) move() {
	newX := (v.p.x + v.v.x)
	newY := (v.p.y + v.v.y)
	if newX < 0 {
		newX += maxX
	} else if newX >= maxX {
		newX -= maxX
	}
	if newY < 0 {
		newY += maxY
	} else if newY >= maxY {
		newY -= maxY
	}
	v.p = location{newX, newY}
}

func (v *historian) getQuadrant() int {
	if v.p.x < (maxX-1)/2 {
		if v.p.y < (maxY-1)/2 {
			return 0
		} else if v.p.y > (maxY-1)/2 {
			return 2
		}
	} else if v.p.x > (maxX-1)/2 {
		if v.p.y < (maxY-1)/2 {
			return 1
		} else if v.p.y > (maxY-1)/2 {
			return 3
		}
	}
	return -1
}

type location struct {
	x, y int
}

type velocity struct {
	x, y int
}

func getContents() []historian {
	file, _ := os.Open("input.txt")
	defer file.Close()

	historians := []historian{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, " ")

		positionStr := lineArr[0][2:]
		velicityStr := lineArr[1][2:]

		positionArr := strings.Split(positionStr, ",")
		velicityArr := strings.Split(velicityStr, ",")

		posX, _ := strconv.Atoi(positionArr[0])
		posY, _ := strconv.Atoi(positionArr[1])
		velX, _ := strconv.Atoi(velicityArr[0])
		velY, _ := strconv.Atoi(velicityArr[1])

		historians = append(historians, historian{location{posX, posY}, velocity{velX, velY}})
	}
	return historians
}

func plot(historians []historian) {

	grid := [][]int{}
	for _ = range maxY {
		row := []int{}
		for _ = range maxX {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}

	for _, h := range historians {
		x, y := h.p.x, h.p.y
		grid[y][x]++
	}

	gridStr := [][]string{}
	for _, row := range grid {
		newRow := []string{}
		for _, value := range row {
			if value == 0 {
				newRow = append(newRow, ".")
			} else {
				strValue := strconv.Itoa(value)
				newRow = append(newRow, strValue)
			}
		}
		gridStr = append(gridStr, newRow)
	}

	fmt.Println("###################")
	for _, row := range gridStr {
		fmt.Println(row)
	}
	fmt.Println("###################")
}

func part1() {

	n := 100
	historians := getContents()

	for _ = range n {
		for i := range historians {
			historians[i].move()
		}
	}

	quadCounts := []int{0, 0, 0, 0}
	for _, historian := range historians {
		quad := historian.getQuadrant()
		if quad == -1 {
			continue
		}
		quadCounts[quad]++
	}

	ans := 1
	for _, quad := range quadCounts {
		ans *= quad
	}

	fmt.Printf("Part 1 Ans: %v\n", ans)
}

func part2() {
	historians := getContents()

	seconds := 0
	for {
		seconds += 1

		grid := [][]rune{}
		for _ = range maxY {
			row := []rune{}
			for _ = range maxX {
				row = append(row, '.')
			}
			grid = append(grid, row)
		}

		for i := range historians {
			historians[i].move()
			grid[historians[i].p.y][historians[i].p.x] = '#'
		}

		maxConsecutive := 0
		for y := 0; y < maxY; y++ {
			consecutive := 0
			for x := 0; x < maxX; x++ {
				if grid[y][x] == '#' {
					consecutive++
				} else {
					maxConsecutive = int(math.Max(float64(maxConsecutive), float64(consecutive)))
					consecutive = 0
				}
			}
		}
		if maxConsecutive > 20 {
			break
		}
	}

	plot(historians)

	fmt.Printf("Part 2 Ans: %v\n", seconds)
}

func main() {
	part1()
	part2()
}
