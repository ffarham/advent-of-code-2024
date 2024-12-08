package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Holder struct {
	arr *[]int
}

func getContents() string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	var content strings.Builder

	for scanner.Scan() {

		line := scanner.Text()
		content.WriteString(line)

	}

	return content.String()

}

func isDigit(str string) bool {

	_, err := strconv.Atoi(str)
	return err == nil
}

func parse(data, prev, state string, digitCount int, doFlag, seenComma bool, output *Holder) (bool, string) {

	if len(data) == 0 {
		return false, ""
	}

	firstChar := string(data[0])

	if prev == "" {

		if firstChar == "m" {
			if !doFlag {
				return parse(data[1:], "", "", 0, doFlag, seenComma, output)
			}

			isValid, result := parse(data[1:], "m", state+firstChar, 0, doFlag, false, output)
			if isValid {
				betweenParenthesis := result[4 : len(result)-1]
				nums := strings.Split(betweenParenthesis, ",")

				int1, err1 := strconv.Atoi(nums[0])
				if err1 != nil {
					panic(err1)
				}
				int2, err2 := strconv.Atoi(nums[1])
				if err2 != nil {
					panic(err2)
				}

				newArr := append(*((*output).arr), int1*int2)
				(*output).arr = &newArr

			}

			return parse(data[1:], "", "", 0, doFlag, false, output)

		} else if firstChar == "d" {
			do := "do()"
			dont := "don't()"
			if len(data) > 7 {
				if do == data[:4] {
					return parse(data[4:], "", "", 0, true, false, output)
				} else if dont == data[:7] {
					return parse(data[7:], "", "", 0, false, false, output)
				} else {
					return parse(data[1:], "", "", 0, doFlag, seenComma, output)
				}
			} else {
				return false, ""
			}

		} else {
			return parse(data[1:], "", "", 0, doFlag, false, output)
		}
	} else if prev == "m" {
		if firstChar == "u" {
			return parse(data[1:], "u", state+firstChar, 0, doFlag, false, output)
		}
		return false, ""

	} else if prev == "u" {
		if firstChar == "l" {
			return parse(data[1:], "l", state+firstChar, 0, doFlag, false, output)
		}
		return false, ""
	} else if prev == "l" {
		if firstChar == "(" {
			return parse(data[1:], "(", state+firstChar, 0, doFlag, false, output)
		}
		return false, ""
	} else if prev == "(" {
		if isDigit(firstChar) {
			return parse(data[1:], firstChar, state+firstChar, 1, doFlag, false, output)
		}
		return false, ""
	} else if isDigit(prev) {
		if isDigit(firstChar) {
			if digitCount < 3 {
				return parse(data[1:], firstChar, state+firstChar, digitCount+1, doFlag, seenComma, output)
			}
			return false, ""
		} else if firstChar == "," {
			if seenComma {
				return false, ""
			}
			return parse(data[1:], firstChar, state+firstChar, 0, doFlag, true, output)
		} else if firstChar == ")" {
			if seenComma {
				return true, state + ")"
			}
			return false, ""
		} else {
			return false, ""
		}
	} else if prev == "," {
		if isDigit(firstChar) {
			return parse(data[1:], firstChar, state+firstChar, 1, doFlag, true, output)
		}
		return false, ""
	} else {
		return false, ""

	}

}

func main() {
	fmt.Println("Advent of code day 3 solution")

	data := getContents()

	holder := Holder{arr: &[]int{}}

	parse(data, "", "", 0, true, false, &holder)

	ans := 0
	for _, val := range *(holder.arr) {
		ans += val
	}

	fmt.Printf("Ans: %v\n", ans)
}
