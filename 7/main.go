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
	flag.StringVar(&filePath, "file", "input.txt", "File path to read")
	flag.Parse()

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		lhs, rhs := parseLine(line)
		total += solve(lhs, rhs)
	}
	fmt.Println(total)
}

func solve(lhs int, rhs []int) int {
	for _, v := range parcioalSolve(rhs) {
		if v == lhs {
			return v
		}
	}
	return 0
}

func parcioalSolve(nums []int) []int {
	results := []int{0}
	for i := 0; i < len(nums); i++ {
		temp := make([]int, len(results))
		copy(temp, results)
		temp2 := make([]int, len(results))
		copy(temp2, results)
		for j := 0; j < len(results); j++ {
			results[j] += nums[i]
			temp[j] *= nums[i] * 1
			temp2[j], _ = strconv.Atoi(strconv.Itoa(temp2[j]) + strconv.Itoa(nums[i]))

		}
		results = append(results, temp...)
		results = append(results, temp2...)
	}
	return results
}
func parseLine(line string) (int, []int) {
	eq := strings.Split(line, ": ")
	res := eq[0]
	lhs, _ := strconv.Atoi(res)
	rhs := make([]int, 0)
	for _, v := range strings.Split(eq[1], " ") {
		n, _ := strconv.Atoi(v)
		rhs = append(rhs, n)
	}
	return lhs, rhs
}
