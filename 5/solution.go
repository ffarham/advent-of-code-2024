package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getContents() (map[int][]int, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	rules := make(map[int][]int)
	updates := [][]int{}

	isRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			isRules = false
			continue
		}

		if isRules {
			words := strings.Split(line, "|")
			leftInt, _ := strconv.Atoi(words[0])
			rightInt, _ := strconv.Atoi(words[1])

			_, ok := rules[rightInt]
			if ok {
				rules[rightInt] = append(rules[rightInt], leftInt)
			} else {
				rules[rightInt] = []int{leftInt}
			}

		} else {
			words := strings.Split(line, ",")
			lineUpdate := []int{}
			for _, word := range words {
				wordInt, _ := strconv.Atoi(word)
				lineUpdate = append(lineUpdate, wordInt)
			}
			updates = append(updates, lineUpdate)
		}

	}

	return rules, updates
}

func part1() {
	rules, updates := getContents()

	ans := 0

	for _, update := range updates {
		inOrder := true
	update_list:
		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				if slices.Contains(rules[update[i]], update[j]) {
					inOrder = false
					break update_list
				}
			}
		}

		if !inOrder {
			continue
		}

		if len(update)%2 == 0 {
			panic(errors.New("even length update"))
		}

		mid := len(update) / 2
		ans += update[mid]

	}

	fmt.Printf("Part 1 Ans: %v\n", ans)
}

func part2() {
	rules, updates := getContents()

	ans := 0

	for _, update := range updates {

		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {

				if slices.Contains(rules[update[i]], update[j]) {
					slices.SortFunc(update, func(a, b int) int {
						aValue, aOk := rules[a]
						if aOk {
							if slices.Contains(aValue, b) {
								return 1
							}
						}

						bValue, bOk := rules[b]
						if bOk {
							if slices.Contains(bValue, a) {
								return -1
							}
						}

						return 0

					})

					if len(update)%2 == 0 {
						panic(errors.New("even length update"))
					}

					mid := len(update) / 2
					ans += update[mid]
				}

			}
		}

	}

	fmt.Printf("Part 2 Ans: %v\n", ans)
}

func main() {
	fmt.Println("Advent of code day 5 solution")

	part1()
	part2()
}
