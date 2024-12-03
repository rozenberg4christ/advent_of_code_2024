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

	safeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		levels := make([]int, len(fields))
		valid := true
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Invalid number in line:", line)
				valid = false
				break
			}
			levels[i] = num
		}
		if !valid {
			continue
		}

		if isSafeReport(levels) {
			safeReports++
		} else if canBeMadeSafe(levels) {
			safeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println("Number of safe reports:", safeReports)
}

func isSafeReport(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	diffs := make([]int, len(levels)-1)
	for i := 1; i < len(levels); i++ {
		diffs[i-1] = levels[i] - levels[i-1]
	}

	var trend string
	firstDiff := diffs[0]
	if firstDiff > 0 {
		trend = "increasing"
	} else if firstDiff < 0 {
		trend = "decreasing"
	} else {
		return false
	}

	for _, diff := range diffs {
		if diff == 0 {
			return false
		}

		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}

		if trend == "increasing" && diff <= 0 {
			return false
		}
		if trend == "decreasing" && diff >= 0 {
			return false
		}
	}

	return true
}

func canBeMadeSafe(levels []int) bool {
	n := len(levels)
	for i := 0; i < n; i++ {
		newLevels := append([]int{}, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)

		if isSafeReport(newLevels) {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
