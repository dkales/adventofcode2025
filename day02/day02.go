package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testinput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

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
	return stage(input, ismirrornum)
}

func stage2(input string) int {
	return stage(input, issupermirrornum)
}

func stage(input string, classifier func(int) int) int {
	segments := strings.Split(strings.TrimSpace(input), ",")
	sum := 0
	for _, segment := range segments {
		firstlast := strings.SplitN(segment, "-", 2)
		first, err := strconv.Atoi(firstlast[0])
		if err != nil {
			panic(err)
		}
		last, err := strconv.Atoi(firstlast[1])
		if err != nil {
			panic(err)
		}
		for i := first; i <= last; i++ {
			sum += classifier(i)
		}

	}
	return sum
}

func ismirrornum(number int) int {
	input := strconv.Itoa(number)
	n := len(input)
	if n%2 != 0 {
		return 0
	}
	if input[:n/2] == input[n/2:] {
		num, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		return num
	} else {
		return 0
	}
}
func issupermirrornum(number int) int {
	input := strconv.Itoa(number)
	n := len(input)
	for sublen := 1; sublen <= n/2; sublen++ {
		if n%sublen != 0 {
			continue
		}
		count := n / sublen
		matched := true
		for i := 1; i < count; i++ {
			if input[0:sublen] != input[i*sublen:(i+1)*sublen] {
				matched = false
				break
			}
		}
		if matched {
			num, err := strconv.Atoi(input)
			if err != nil {
				panic(err)
			}
			return num
		}
	}
	return 0
}
