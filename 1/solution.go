package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getContents() ([]int, []int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	leftArr := []int{}
	rightArr := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "   ")

		leftInt, leftErr := strconv.Atoi(words[0])
		if leftErr != nil {
			panic(leftErr)
		}
		rightInt, rightErr := strconv.Atoi(words[1])
		if rightErr != nil {
			panic(rightErr)
		}

		leftArr = append(leftArr, leftInt)
		rightArr = append(rightArr, rightInt)

	}

	if len(leftArr) != len(rightArr) {
		panic(errors.New("the left and right arrays are of different sizes"))
	}

	return leftArr, rightArr
}

func part1() {

	leftArr, rightArr := getContents()

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	ans := 0.0
	for i := 0; i < len(leftArr); i++ {
		ans += math.Abs(float64(leftArr[i] - rightArr[i]))
	}

	fmt.Printf("Part 1 Ans: %v\n", int(ans))
}

func part2() {

	leftArr, rightArr := getContents()

	rightStore := make(map[int]int)
	for i := 0; i < len(leftArr); i++ {
		rightStore[rightArr[i]] += 1
	}

	ans := 0
	for _, val := range leftArr {
		ans += val * rightStore[val]
	}

	fmt.Printf("Part 2 Ans: %v\n", ans)

}

func main() {
	fmt.Println("Solution for day 1 of advent of code 2024")

	part1()
	part2()

}
