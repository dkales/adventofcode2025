package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const testinput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
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
	var modified string
	modified = strings.ReplaceAll(input, "L", "-")
	modified = strings.ReplaceAll(modified, "R", "")
	lines := strings.Split(strings.TrimSpace(modified), "\n")
	current := 50
	res := 0
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		current = (current + num + 100) % 100
		if current == 0 {
			res += 1
		}
	}
	return res
}

func stage2(input string) int {
	var modified string
	modified = strings.ReplaceAll(input, "L", "-")
	modified = strings.ReplaceAll(modified, "R", "")
	lines := strings.Split(strings.TrimSpace(modified), "\n")
	current := 50
	res := 0
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		abs := int(math.Abs(float64(num)))
		res += abs / 100
		num = num % 100
		new_unred := current + num
		new := (current + num + 100) % 100
		if (new_unred != new || new == 0) && current != 0 {
			res += 1
		}
		current = new
	}
	return res
}
