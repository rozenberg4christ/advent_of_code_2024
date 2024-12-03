package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	if len(leftList) != len(rightList) {
		fmt.Println("Error: The two lists have different lengths.")
		return
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	fmt.Println("Total distance:", totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
