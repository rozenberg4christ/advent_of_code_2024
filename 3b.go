package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Define the regular expression pattern
	// Pattern explanation:
	// - `mul\(\d{1,3},\d{1,3}\)` matches valid mul instructions
	// - `do\(\)` matches do() instruction
	// - `don't\(\)` matches don't() instruction
	pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringIndex(input, -1)
	totalSum := 0
	enabled := true

	for _, loc := range matches {
		start, end := loc[0], loc[1]
		match := input[start:end]

		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if len(match) >= 4 && match[:4] == "mul(" {
			if enabled {
				numbers := match[4 : len(match)-1]

				nums := splitNumbers(numbers)

				if len(nums) != 2 {
					continue
				}
				x, err1 := strconv.Atoi(nums[0])
				y, err2 := strconv.Atoi(nums[1])

				if err1 != nil || err2 != nil {
					continue
				}

				totalSum += x * y
			}
		}
	}

	fmt.Println("Total sum of enabled multiplications:", totalSum)
}

func readInput(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return input, nil
}

func splitNumbers(numbers string) []string {
	return regexp.MustCompile(`,`).Split(numbers, -1)
}
