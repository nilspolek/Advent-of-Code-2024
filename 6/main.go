package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", "input.txt", "File path")
	flag.Parse()

	var grid [][]rune
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		temp := []rune{}
		for _, c := range scanner.Text() {
			temp = append(temp, c)
		}
		grid = append(grid, temp)
	}
	guard, currentDir := getPosition(grid)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '#' {
				for k := j; k < len(grid[0]); k++ {
					if grid[i][k] == '#' {

					}
				}
			}
		}
	}

	for inGrid(grid, []int{guard[0] + currentDir[0], guard[1] + currentDir[1]}) {
		next := grid[guard[0]+currentDir[0]][guard[1]+currentDir[1]]
		if next != '#' {
			guard = []int{guard[0] + currentDir[0], guard[1] + currentDir[1]}
			grid[guard[0]][guard[1]] = 'X'
		} else {
			currentDir = newDir(currentDir)
		}
	}
	printGrid(grid)
}

func getPosition(grid [][]rune) ([]int, [2]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '>' {
				return []int{i, j}, [2]int{0, 1}
			}
			if grid[i][j] == '<' {
				return []int{i, j}, [2]int{0, -1}
			}
			if grid[i][j] == '^' {
				return []int{i, j}, [2]int{-1, 0}
			}
			if grid[i][j] == 'v' {
				return []int{i, j}, [2]int{1, 0}
			}
		}
	}
	return []int{0, 0}, [2]int{0, 0}
}

func printGrid(grid [][]rune) {
	totalVisited := 0
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(string(c))
			if c != '.' && c != '#' {
				totalVisited++
			}
		}
		fmt.Println()
	}

	fmt.Println(totalVisited)
}

func inGrid(grid [][]rune, position []int) bool {
	if position[0] < 0 || position[0] >= len(grid) || position[1] < 0 || position[1] >= len(grid[0]) {
		return false
	}
	return true
}

func newDir(dir [2]int) [2]int {
	switch dir {
	case [2]int{1, 0}: // Down
		return [2]int{0, -1}
	case [2]int{0, 1}: // Right
		return [2]int{1, 0}
	case [2]int{-1, 0}: // UP
		return [2]int{0, 1}
	case [2]int{0, -1}: // Left
		return [2]int{-1, 0}
	}
	return [2]int{0, 0}
}
