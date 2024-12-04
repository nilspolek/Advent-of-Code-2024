package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

// 2530 - 2554
var filePath string

func main() {
	flag.StringVar(&filePath, "file", "input.txt", "file path")
	flag.Parse()

	var (
		re = regexp.MustCompile(`XMAS`)
	)

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	input := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line))
	}

	totalString := ""
	// Collect all the XMases Horizontal
	for _, line := range input {
		totalString += string(line)
		totalString += "\n"
	}
	// Collect all the XMases Vertical
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			totalString += string(input[j][i])
		}
		totalString += "\n"
	}

	// Collect all the XMases Diagonal
	m := len(input[0])
	n := len(input)
	for startCol := 0; startCol < m; startCol++ {
		var diagonal string
		x, y := 0, startCol
		for x < n && y < m {
			diagonal += string(input[x][y])
			x++
			y++
		}
		totalString += diagonal
		totalString += "\n"

	}

	// Start from each element in the first column
	for startRow := 1; startRow < n; startRow++ {
		var diagonal string
		x, y := startRow, 0
		for x < n && y < m {
			diagonal += string(input[x][y])
			x++
			y++
		}
		totalString += diagonal
		totalString += "\n"

	}

	for startCol := m - 1; startCol >= 0; startCol-- {
		var diagonal string
		x, y := 0, startCol
		for x < n && y >= 0 {
			diagonal += string(input[x][y])
			x++
			y--
		}
		totalString += diagonal
		totalString += "\n"
	}

	for startRow := 1; startRow < n; startRow++ {
		var diagonal string
		x, y := startRow, m-1
		for x < n && y >= 0 {
			diagonal += string(input[x][y])
			x++
			y--
		}
		totalString += diagonal
		totalString += "\n"
	}

	totalString += reverseString(totalString)
	occurences := re.FindAllStringSubmatch(totalString, -1)
	fmt.Printf("Occurences of XMas:\t\t%d\n", len(occurences))

	counter := 0
	for k := 0; k < 4; k++ {
		for i := 0; i < len(input[0])-2; i++ {
			for j := 0; j < len(input)-2; j++ {
				if isXMAS(input, i, j) {
					counter++
				}
			}
		}
		rotateMatrix(&input)
	}
	fmt.Printf("Occurences of X formed MAS:\t%d\n", counter)
}

func isXMAS(input [][]rune, i, j int) bool {
	return input[i][j] == 'M' && input[i+2][j] == 'S' && input[i+1][j+1] == 'A' && input[i][j+2] == 'M' && input[i+2][j+2] == 'S'
}

func reverseString(s string) (out string) {
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}
	return
}

func rotateMatrix(matrix *[][]rune) {
	n := len(*matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			(*matrix)[i][j], (*matrix)[j][i] = (*matrix)[j][i], (*matrix)[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			(*matrix)[i][j], (*matrix)[i][k] = (*matrix)[i][k], (*matrix)[i][j]
		}
	}
}
