package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getContents() [][]int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	lines := [][]int{}

	for scanner.Scan() {

		line := scanner.Text()
		words := strings.Split(line, " ")
		ints := []int{}
		for _, word := range words {
			intValue, convError := strconv.Atoi(word)
			if convError != nil {
				panic(convError)
			}
			ints = append(ints, intValue)
		}
		lines = append(lines, ints)

	}

	return lines

}

func removeIndex(arr []int, index int) []int {
	newArr := make([]int, len(arr)-1)
	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, arr[index+1:]...)
	return newArr
}

func validator(line []int) bool {
	safeFlag := true
	prevFlag := true
	increasing := true
	prev := 0
	for i, val := range line {
		if prevFlag {
			prev = val
			prevFlag = false
			continue
		}

		// equal check
		if val == prev {
			safeFlag = false
			break
		}

		// determine order
		if i == 1 {
			if val < prev {
				increasing = false
			}
		}

		// order check
		if increasing {
			if val < prev {
				safeFlag = false
				break
			}
		} else {
			if val > prev {
				safeFlag = false
				break
			}
		}

		// diff check
		if math.Abs(float64(val-prev)) > 3 {
			safeFlag = false
			break
		}

		prev = val

	}

	return safeFlag
}

func part1(line []int) bool {
	return validator(line)
}

func part2(line []int) bool {
	flag := false
	for i := 0; i < len(line); i++ {
		flag = validator(removeIndex(line, i))
		if flag {
			break
		}
	}
	return flag
}

func solver() {
	lines := getContents()

	part1Count := 0
	part2Count := 0

	for _, line := range lines {

		part1Flag := part1(line)

		if part1Flag {
			part1Count++
		}

		part2Flag := part2(line)
		if part2Flag {
			part2Count++
		}

	}

	fmt.Printf("Part 1 Ans: %v\n", part1Count)
	fmt.Printf("Part 2 Ans: %v\n", part2Count)
}

func main() {
	solver()
}
