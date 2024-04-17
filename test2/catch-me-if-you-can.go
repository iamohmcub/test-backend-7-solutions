package test2

import (
	"fmt"
	"math"
)

func CatchMeIfYouCan(input string) string {
	minimumSummary := math.MaxInt64
	result := ""
	inputLength := len(input)
	loopLength := int(math.Pow(10, float64(inputLength+1)))

	for i := 0; i < loopLength; i++ {
		possibleCharacter := fmt.Sprintf("%0*d", inputLength+1, i)
		isMatch := isMatchCondition(possibleCharacter, input)
		if isMatch {
			currentCharacterSum := calculateSum(possibleCharacter)
			if currentCharacterSum < minimumSummary {
				minimumSummary = currentCharacterSum
				result = possibleCharacter
			}
		}
	}
	return result
}

func isMatchCondition(possibleCharacter string, input string) bool {
	for index := 0; index < len(input); index++ {
		left := int(possibleCharacter[index] - '0')
		right := int(possibleCharacter[index+1] - '0')
		switch input[index] {
		case 'L':
			if left <= right {
				return false
			}
		case 'R':
			if left >= right {
				return false
			}
		case '=':
			if left != right {
				return false
			}
		}

	}

	return true
}

func calculateSum(sequence string) int {
	sum := 0
	for _, c := range sequence {
		sum += int(c - '0')
	}
	return sum
}
