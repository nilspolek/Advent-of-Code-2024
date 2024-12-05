package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Lhs int
	Rhs int
}

var filePath string
var rules []Rule

func main() {
	flag.StringVar(&filePath, "file", "input.txt", "path to the rules file")
	flag.Parse()
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	updates := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rule := parseRule(line)
		rules = append(rules, rule)
	}
	for scanner.Scan() {
		line := scanner.Text()
		updates = append(updates, parseLine(line))
	}

	result := 0
	for _, update := range updates {
		if isCorrectOrder(update, rules) {
			result += int(update[len(update)/2])
		}
	}
	fmt.Printf("Middle of updates:\t%d\n", result)
}

func parseLine(line string) (out []int) {
	numbers := strings.Split(line, ",")

	for _, num := range numbers {
		if res, err := strconv.Atoi(num); err == nil {
			out = append(out, res)
		} else {
			panic(err)
		}
	}
	return
}

func parseRule(line string) Rule {
	nums := strings.Split(line, "|")
	lhs, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}
	rhs, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}
	return Rule{
		Lhs: lhs,
		Rhs: rhs,
	}
}

func isCorrectOrder(update []int, rules []Rule) bool {
	for _, rule := range rules {
		x, y := rule.Lhs, rule.Rhs
		xIndex, yIndex := -1, -1
		for i, page := range update {
			if page == x {
				xIndex = i
			}
			if page == y {
				yIndex = i
			}
		}
		if xIndex != -1 && yIndex != -1 && xIndex > yIndex {
			return false
		}
	}
	return true
}
