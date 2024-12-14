package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type block struct {
	id               int
	startIndex, size int64
}

func (v *block) isFree() bool {
	return -1 == v.id
}

func (v block) String() string {
	output := ""
	for _ = range v.size {
		if v.id == -1 {
			output += "."
		} else {
			output += strconv.Itoa(v.id)
		}
	}
	return output
}

type memory []block

func (v memory) String() string {
	output := ""
	for _, block := range v {
		output += block.String()
	}
	return output
}

func getContents() []rune {
	file, _ := os.Open("input.txt")
	defer file.Close()

	output := []rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineRune := []rune(line)
		output = append(output, lineRune...)
	}

	return output
}

func flatten(denseStr *[]rune, contiguous bool) memory {

	flatStr := memory{}
	id := 0
	isFile := true
	index := int64(0)

	for _, char := range *denseStr {

		num, _ := strconv.ParseInt(string(char), 10, 64)
		if isFile {
			if contiguous {
				flatStr = append(flatStr, block{id: id, startIndex: index, size: num})
				id++
				isFile = false
				index += num
			} else {
				for _ = range num {
					flatStr = append(flatStr, block{id: id, startIndex: index, size: 1})
					index++
				}
				id++
				isFile = false
			}
		} else {
			if contiguous {
				flatStr = append(flatStr, block{id: -1, startIndex: index, size: num})
				isFile = true
				index += num
			} else {
				for _ = range num {
					flatStr = append(flatStr, block{id: -1, startIndex: index, size: 1})
					index++
				}
				isFile = true
			}
		}
	}

	return flatStr
}

func getNextFreeIndex(startIndex int, mem *memory) int {
	n := len(*mem)
	for i := startIndex; i < n; i++ {
		if (*mem)[i].isFree() {
			return i
		}
	}
	return n
}

func getNextFreeBlockOfSizeN(n int64, mem *memory) int {
	for index, block := range *mem {
		if block.isFree() && n <= block.size {
			return index
		}
	}
	return -1
}

func makeCompact(mem *memory) {

	n := len(*mem)
	freeIndex := getNextFreeIndex(0, mem)

	for i := n - 1; i >= 0; i-- {
		if freeIndex >= i {
			break
		}

		if (*mem)[i].isFree() {
			continue
		}

		temp := (*mem)[freeIndex]
		(*mem)[freeIndex] = (*mem)[i]
		(*mem)[i] = temp
		freeIndex = getNextFreeIndex(freeIndex, mem)
	}
}

func moveContiguousBlock(source, target int, mem memory) (memory, int) {

	sourceStartindex := mem[source].startIndex
	targetStartIndex := mem[target].startIndex
	diff := mem[target].size - mem[source].size

	shift := 0
	mem[target] = mem[source]
	mem[target].startIndex = targetStartIndex
	mem[source] = block{id: -1, startIndex: sourceStartindex, size: mem[source].size}
	if diff > 0 {
		mem = append(mem[:target+2], mem[target+1:]...)
		mem[target+1] = block{id: -1, startIndex: targetStartIndex + mem[target].size, size: diff}
		shift++

		// merge to right if possible
		if mem[target+2].isFree() {
			mem[target+1].size += mem[target+2].size
			mem = append(mem[:target+2], mem[target+3:]...)
		}
	}

	// merge around source if possible
	if source+shift < len(mem)-1 && mem[source+shift+1].isFree() {
		mem[source+shift].size += mem[source+shift+1].size
		mem = append(mem[:source+shift+1], mem[source+shift+2:]...)
	}
	if mem[source+shift-1].isFree() {
		farLeftStartIndex := mem[source+shift-1].startIndex
		mem[source+shift].size += mem[source+shift-1].size
		mem = append(mem[:source+shift-1], mem[source+shift:]...)
		mem[source+shift-1].startIndex = farLeftStartIndex
		shift--
	}

	return mem, shift

}

func makeContigiousCompact(mem memory) memory {

	n := len(mem)
	currIndex := n - 1
	for currIndex >= 0 {

		block := mem[currIndex]
		if block.isFree() {
			currIndex--
			continue
		}

		freeIndex := getNextFreeBlockOfSizeN(int64(block.size), &mem)
		if freeIndex == -1 {
			currIndex--
			continue
		}
		if freeIndex >= currIndex {
			currIndex--
			continue
		}

		newMem, shift := moveContiguousBlock(currIndex, freeIndex, mem)
		mem = newMem
		currIndex = currIndex - (1 - shift)

		n = len(mem)
	}

	return mem
}

func calculateChecksumPart1(mem *memory) int64 {

	output := int64(0)
	for index, block := range *mem {
		if block.isFree() {
			continue
		}
		output += int64(index) * int64(block.id)
	}
	return output
}

func calculateChecksumPart2(mem *memory) int64 {

	output := int64(0)
	for _, block := range *mem {
		if block.isFree() {
			continue
		}
		for i := block.startIndex; i < block.startIndex+block.size; i++ {
			output += int64(i) * int64(block.id)
		}
	}
	return output
}

func part1(denseStr *[]rune) {

	flatStr := flatten(denseStr, false)
	makeCompact(&flatStr)

	fmt.Printf("Part 1 Ans: %v\n", calculateChecksumPart1(&flatStr))
}

func part2(denseStr *[]rune) {

	mem := flatten(denseStr, true)
	mem = makeContigiousCompact(mem)

	fmt.Printf("Part 2 Ans: %v\n", calculateChecksumPart2(&mem))
}

func main() {

	denseStr := getContents()
	part1(&denseStr)
	part2(&denseStr)
}
