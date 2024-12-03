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

	// Pattern explanation:
	// - ^ and $ assert the start and end of the string (not used here to allow multiple matches)
	// - `mul\(` matches 'mul(' exactly
	// - `\d{1,3}` matches 1 to 3 digits
	// - `,` matches a comma
	// - `\d{1,3}` matches another 1 to 3 digits
	// - `\)` matches a closing parenthesis
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)
	totalSum := 0

	for _, match := range matches {
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

	fmt.Println("Total sum of valid multiplications:", totalSum)
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
