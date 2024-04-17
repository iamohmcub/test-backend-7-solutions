package test1

import (
	"encoding/json"
	"fmt"
	"os"
)

func findMaxPath(input [][]int) int {
	for row := len(input) - 1; row > 0; row-- {
		for index := 0; index < len(input[row-1]); index++ {
			current := input[row][index]
			next := input[row][index+1]
			if current > next {
				input[row-1][index] += current
				continue
			}
			input[row-1][index] += next
		}
	}
	return input[0][0]
}

func GetGreatestPath(testCase string) string {
	var input [][]int
	switch testCase {
	case "2":
		file, _ := os.Open("test1/hard.json")
		decoder := json.NewDecoder(file)
		decoder.Decode(&input)
	default:
		input = [][]int{
			{59},
			{73, 41},
			{52, 40, 53},
			{26, 53, 6, 34},
		}
	}

	result := findMaxPath(input)
	str := fmt.Sprintf("%d", result)
	return str
}
