package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Invalid line (does not contain exactly two numbers):", line)
			continue
		}

		leftNum, err1 := strconv.Atoi(fields[0])
		rightNum, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers in line:", line)
			continue
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	rightFreq := make(map[int]int)
	for _, num := range rightList {
		rightFreq[num]++
	}

	totalSimilarity := 0
	for _, num := range leftList {
		freq := rightFreq[num]
		totalSimilarity += num * freq
	}

	fmt.Println("Total similarity score:", totalSimilarity)
}
