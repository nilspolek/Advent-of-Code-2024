package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", "input.txt", "file path")
	flag.Parse()
	lines := parse(filePath)
	fmt.Printf("Total lines: %d\n", len(lines))
	totalSaveLines := 0
	totalSaveWithErrDetLines := 0
	for _, line := range lines {
		if isSafe(line) == 0 {
			totalSaveLines++
		}
		if isSafeWithErr(line) {
			totalSaveWithErrDetLines++
		}
	}
	fmt.Printf("Total save lines: %d\n", totalSaveLines)
	fmt.Printf("Save line Percentage: %d%%\n", totalSaveLines*100/len(lines))
	fmt.Printf("Total save lines with err: %d\n", totalSaveWithErrDetLines)
	fmt.Printf("Save line Percentage with err: %d%%\n", totalSaveWithErrDetLines*100/len(lines))
}

func isSafeWithErr(line []int) bool {
	if isSafe(line) == 0 {
		return true
	}
	for i := 0; i < len(line); i++ {
		newLine := append([]int(nil), line[:i]...)
		newLine = append(newLine, line[i+1:]...)

		if isSafe(newLine) == 0 {
			return true
		}
	}
	return false
}

func isSafe(line []int) int {
	isDecrease := line[0] > line[1]
	last := line[0]
	totalErrors := 0
	for i := 1; i < len(line); i++ {
		increase := abs(line[i] - last)
		if increase == 0 {
			totalErrors++
		}
		if increase > 3 {
			totalErrors++
		}
		if isDecrease && line[i] > last {
			totalErrors++
		}
		if !isDecrease && line[i] < last {
			totalErrors++
		}
		last = line[i]
	}
	return totalErrors
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func parse(fp string) (out [][]int) {
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp := []int{}
		for _, c := range strings.Fields(line) {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			temp = append(temp, num)
		}
		out = append(out, temp)
	}
	return
}
