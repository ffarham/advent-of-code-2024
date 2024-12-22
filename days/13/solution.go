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
	buttonACost = 3
	buttonBCost = 1
)

type location struct {
	x, y int
}

func (v *location) isValid(prize location) bool {
	return v.x <= prize.x && v.y <= prize.y
}

type button struct {
	xDelta, yDelta int
}

func (v *button) press(loc location) location {
	return location{loc.x + v.xDelta, loc.y + v.yDelta}
}

type machine struct {
	buttonA, buttonB button
	prize            location
}

func (v *machine) atPrize(loc location) bool {
	return v.prize.x == loc.x && v.prize.y == loc.y
}

func getButton(buttonLine string) button {
	buttonValue := buttonLine[10:]
	buttonXY := strings.Split(buttonValue, ", ")
	xDelta, _ := strconv.Atoi(buttonXY[0][2:])
	yDelta, _ := strconv.Atoi(buttonXY[1][2:])
	return button{xDelta, yDelta}
}

func getPrizeLoc(prizeLine string, increaseby int) location {
	prizeValue := prizeLine[7:]
	prizeXY := strings.Split(prizeValue, ", ")
	x, _ := strconv.Atoi(prizeXY[0][2:])
	y, _ := strconv.Atoi(prizeXY[1][2:])
	return location{x + increaseby, y + increaseby}
}

func getContents(increaseBy int) []machine {
	file, _ := os.Open("input.txt")
	defer file.Close()

	machines := []machine{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buttonALine := scanner.Text()
		scanner.Scan()
		buttonBLine := scanner.Text()
		scanner.Scan()
		prizeLine := scanner.Text()

		machines = append(machines, machine{
			getButton(buttonALine),
			getButton(buttonBLine),
			getPrizeLoc(prizeLine, increaseBy),
		})
		scanner.Scan()
	}
	return machines
}

func search(loc location, m *machine, memo *map[location]int64) int64 {

	if m.atPrize(loc) {
		return 0
	}

	visited, ok := (*memo)[loc]
	if ok {
		return visited
	}

	if !loc.isValid(m.prize) {
		return -1
	}

	buttonARoute := search(m.buttonA.press(loc), m, memo)
	buttonBRoute := search(m.buttonB.press(loc), m, memo)
	if buttonARoute == -1 && buttonBRoute == -1 {
		(*memo)[loc] = -1
		return -1
	} else if buttonARoute == -1 {
		(*memo)[loc] = buttonBRoute + buttonBCost
		return buttonBRoute + buttonBCost
	} else if buttonBRoute == -1 {
		(*memo)[loc] = buttonARoute + buttonACost
		return buttonARoute + buttonACost
	} else {
		minCost := int64(math.Min(float64(buttonARoute+buttonACost), float64(buttonBRoute+buttonBCost)))
		(*memo)[loc] = minCost
		return minCost
	}

}

func part1() {
	machines := getContents(0)

	totalTokens := int64(0)
	for _, machine := range machines {
		initialLoc := location{0, 0}

		memo := make(map[location]int64)
		tokens := search(initialLoc, &machine, &memo)
		if tokens == -1 {
			continue
		}
		totalTokens += tokens
	}

	fmt.Printf("Part 1 Ans: %v\n", totalTokens)
}

func solver(m machine) int64 {
	bNumerator := m.buttonA.xDelta*m.prize.y - m.prize.x*m.buttonA.yDelta
	bDenominator := m.buttonA.xDelta*m.buttonB.yDelta - m.buttonB.xDelta*m.buttonA.yDelta
	if bNumerator%bDenominator != 0 {
		return 0
	}

	b := bNumerator / bDenominator
	aNumberator := m.prize.x - b*m.buttonB.xDelta
	aDenominator := m.buttonA.xDelta
	if aNumberator%aDenominator != 0 {
		return 0
	}

	a := aNumberator / aDenominator

	return int64(3*a + b)

}

func part2() {
	machines := getContents(10000000000000)

	ans := int64(0)
	for _, m := range machines {

		ans += solver(m)
	}
	fmt.Println("Part 2 Ans:", ans)
}

func main() {

	part1()
	part2()

}
