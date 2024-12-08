package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type calibration struct {
	result   int64
	operands []int64
}

func (v *calibration) reproducable() bool {
	return reproducableHelper(v.result, v.operands)
}

func reproducableHelper(target int64, operands []int64) bool {

	// fmt.Println(target, operands)

	if len(operands) == 1 {
		return target == operands[0]
	}

	n := len(operands)

	plusArr := make([]int64, n-1)
	copy(plusArr, operands[:n-1])
	plusArr[n-2] = operands[n-1] + operands[n-2]

	mulArr := make([]int64, n-1)
	copy(mulArr, operands[:n-1])
	mulArr[n-2] = operands[n-1] * operands[n-2]

	concatArr := make([]int64, n-1)
	copy(concatArr, operands[:n-1])
	concatStr := strconv.FormatUint(uint64(operands[n-1]), 10) + strconv.FormatUint(uint64(operands[n-2]), 10)
	concatInt, _ := strconv.ParseInt(concatStr, 10, 64)
	concatArr[n-2] = concatInt

	return reproducableHelper(target, plusArr) || reproducableHelper(target, mulArr) || reproducableHelper(target, concatArr)
}

func getContents() []calibration {
	file, _ := os.Open("input.txt")
	defer file.Close()

	output := []calibration{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, ":")
		resultInt, _ := strconv.ParseInt(lineArr[0], 10, 64)
		operandsInt := []int64{}
		for _, operandStr := range strings.Split(lineArr[1], " ") {
			value, _ := strconv.ParseInt(operandStr, 10, 64)
			operandsInt = append(operandsInt, value)
		}
		slices.Reverse(operandsInt)
		output = append(output, calibration{result: resultInt, operands: operandsInt[:len(operandsInt)-1]})
	}

	return output
}

func part1() {
	data := getContents()

	ans := int64(0)
	for _, cal := range data {
		if cal.reproducable() {
			ans += cal.result
		}
	}

	fmt.Printf("Part 1 Ans: %v\n", ans)
}

func main() {
	part1()
}
