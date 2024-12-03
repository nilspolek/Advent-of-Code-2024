package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", "input.txt", "file path")
	flag.Parse()
	var (
		result = 0

		doRe  = regexp.MustCompile(`do\(\).*?don't\(\)`)
		mulRe = regexp.MustCompile(`mul\(\d+,\d+\)`)
		numRe = regexp.MustCompile(`\d+`)
	)

	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	input := string(file)
	input = "do()" + input + "don't()"
	input = strings.ReplaceAll(input, "\n", "")

	for _, part := range doRe.FindAllString(input, -1) {
		for _, mul := range mulRe.FindAllString(part, -1) {
			nums := numRe.FindAllString(mul, -1)
			lhs, _ := strconv.Atoi(nums[0])
			rhs, _ := strconv.Atoi(nums[1])
			result += lhs * rhs
		}
	}
	fmt.Printf("Result:\t%d\n", result)
}
