package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", "input.txt", "file path")
	flag.Parse()

	data, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	var (
		lhs, rhs []int
	)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		values := strings.Fields(line)
		if len(values) == 2 {
			num1, err := strconv.Atoi(values[0])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			num2, err := strconv.Atoi(values[1])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			lhs = append(lhs, num1)
			rhs = append(rhs, num2)
		}
	}
	fmt.Printf("Distance:\t%d\nSimularity:\t%d\n", getDistance(lhs, rhs), getSimularity(lhs, rhs))
}

func getDistance(arr1, arr2 []int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	totalDistance := 0
	for i := 0; i < len(arr1); i++ {
		totalDistance += abs(arr1[i] - arr2[i])
	}
	return totalDistance
}

func getSimularity(arr1, arr2 []int) (simularity int) {
	counter := make(map[int]int)
	for _, v := range arr2 {
		counter[v]++
	}
	for _, v := range arr1 {
		simularity += v * counter[v]
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
