package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"time"
)

type Word struct {
	value        string
	startingWord bool
	punctuation  bool
}

var dictionary = make(map[string][]Word)
var startingWords = make([]string, 0)

func main() {
	rand.Seed(time.Now().Unix())
	fileBytes, err := ioutil.ReadFile("input.txt")
	re := regexp.MustCompile(`\r?\n`)
	input := re.ReplaceAllString(string(fileBytes), " ")

	var file []Word

	if err == nil {
		word := ""
		startingWord := true
		for _, char := range input {
			//If we hit a space push the word to the array
			if char == ' ' {
				if word != "" {
					file = append(file, Word{value: word, startingWord: startingWord})
					startingWord = false
					word = ""
				}
				continue
			}

			//If we hit punctuation push the word to the array
			if (char == '.') || (char == '?') || (char == '!') {
				file = append(file, Word{value: word, startingWord: startingWord})
				word = string(char)
				file = append(file, Word{value: word, punctuation: true})
				startingWord = true
				word = ""
				continue
			}

			word += string(char)
		}
	}

	for i, word := range file {
		if word.startingWord {
			startingWords = append(startingWords, word.value)
		}
		if len(dictionary[word.value]) == 0 {
			dictionary[word.value] = make([]Word, 0)
		}

		if i < len(file)-1 {
			dictionary[word.value] = append(dictionary[word.value], file[i+1])
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i, " : ", generateSentence())
	}

}

func generateSentence() string {
	//create a random sentence
	done := false
	sentence := ""
	currentWord := ""
	punctuation := false
	for !done {
		if currentWord == "" && len(sentence) == 0 {
			currentWord = startingWords[rand.Intn(len(startingWords))]
		} else {
			if len(dictionary[currentWord]) == 0 {
				done = true
			} else {
				word := dictionary[currentWord][rand.Intn(len(dictionary[currentWord]))]
				currentWord = word.value
				punctuation = word.punctuation
			}
		}

		if punctuation || len(sentence) == 0 {
			sentence += currentWord
			if punctuation {
				done = true
			}
		} else {
			sentence += " " + currentWord
		}
	}

	return sentence
}
