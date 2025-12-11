package main

import (
	"fmt"
	"os"
	"strings"
)

const testinput = `987654321111111
811111111111119
234234234234278
818181911112111
`

func main() {
	fmt.Printf("Stage1(test): %d\n", stage1(testinput))
	fmt.Printf("Stage2(test): %d\n", stage2(testinput))
	// read file into string

	inbytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(inbytes)
	fmt.Printf("Stage1: %d\n", stage1(input))
	fmt.Printf("Stage2: %d\n", stage2(input))
}

func stage1(input string) int {
	return stage(input, 2)
}

func stage2(input string) int {
	return stage(input, 12)
}

func stage(input string, count int) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	res := 0
	for _, line := range lines {
		sum := 0
		start := 0
		end := len(line) - count + 1
		for j := 0; j < count; j += 1 {
			largest := -1
			idx := -1
			for i := start; i < end; i += 1 {
				ch := int(line[i] - '0')
				if ch > largest {
					largest = ch
					idx = i
				}
			}
			sum *= 10
			sum += largest
			start = idx + 1
			end += 1
		}
		res += sum
	}
	return res
}
