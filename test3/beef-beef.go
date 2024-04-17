package test3

import (
	"io"
	"net/http"
	"strings"
)

type BeefTypeCount struct {
	Beef map[string]int `json:"beef"`
}

func countBeef(text string) BeefTypeCount {
	meatCounts := make(map[string]int)

	words := strings.FieldsFunc(text, func(salt rune) bool {
		return salt == ',' || salt == '.' || salt == ' ' || salt == '\n'
	})

	for _, word := range words {
		word = strings.ToLower(word)
		if word != "" {
			meatCounts[word]++
		}
	}
	json := BeefTypeCount{
		Beef: meatCounts,
	}
	return json
}

func GetResponseBeef() BeefTypeCount {
	client := &http.Client{}
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	serverResponse, _ := client.Get(url)
	parsedBody, _ := io.ReadAll(serverResponse.Body)

	return countBeef(string(parsedBody))
}
