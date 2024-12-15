package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stone struct {
	value string
}

func (v *stone) getInt() int {
	num, err := strconv.Atoi(v.value)
	if err != nil {
		panic(err)
	}
	return num
}

func (v *stone) blink() []stone {
	output := []stone{}

	num := v.getInt()
	if num == 0 {
		output = append(output, stone{"1"})
	} else if len(v.value)%2 == 0 {
		mid := len(v.value) / 2
		output = append(output, stone{v.value[:mid]})
		output = append(output, stone{trimZeros(v.value[mid:])})
	} else {
		newStr := strconv.Itoa(v.getInt() * 2024)
		output = append(output, stone{newStr})
	}
	return output
}

func (v stone) String() string {
	return v.value
}

func trimZeros(str string) string {
	n := len(str)
	index := n - 1
	for i := 0; i < n-1; i++ {
		if str[i] != '0' {
			index = i
			break
		}
	}
	return str[index:]
}

func getContents() []stone {

	file, _ := os.Open("input.txt")
	defer file.Close()

	output := []stone{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		for _, word := range words {
			output = append(output, stone{word})
		}
	}
	return output
}

func part1() {
	stones := getContents()
	n := 25

	for _ = range n {
		newStones := []stone{}
		for _, stone := range stones {
			result := stone.blink()
			newStones = append(newStones, result...)
		}
		stones = newStones
	}

	fmt.Printf("Part 1 Ans: %v\n", len(stones))
}

func increment(value int, key string, store *map[string]int) {
	_, ok := (*store)[key]
	if !ok {
		(*store)[key] = 0
	}
	(*store)[key] += value
}

func part2() {
	stones := getContents()
	store := make(map[string]int)
	for _, stone := range stones {
		increment(1, stone.value, &store)
	}

	n := 75
	for _ = range n {

		newStore := make(map[string]int)

		for stone, stoneCount := range store {
			if stone == "0" {
				increment(stoneCount, "1", &newStore)
			} else if len(stone)%2 == 0 {
				left := stone[:len(stone)/2]
				increment(stoneCount, left, &newStore)

				right := trimZeros(stone[len(stone)/2:])
				increment(stoneCount, right, &newStore)
			} else {
				num, _ := strconv.Atoi(stone)
				newNum := strconv.Itoa(num * 2024)
				increment(stoneCount, newNum, &newStore)
			}
		}
		store = newStore
	}

	ans := 0
	for _, count := range store {
		ans += count
	}

	fmt.Printf("Part 2 Ans: %v\n", ans)
}

func main() {
	part1()
	part2()
}
