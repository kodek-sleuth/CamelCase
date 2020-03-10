package main

import (
	"fmt"
	_"regexp"
	"unicode"
	_"unicode"
)

func upperCaseOrLowerCase(r rune) bool {
	if unicode.IsUpper(r) == true{
		return true
	}
	return false
}

func wordCounter(s string) int {
	wordCount := 0
	for index, word := range s {
		if upperCaseOrLowerCase(int32(s[0])) == false {
			if upperCaseOrLowerCase(word) == true && upperCaseOrLowerCase(int32(s[index+1])) == false {
				wordCount++
			}
			continue
		}
		return 0
	}
	if upperCaseOrLowerCase(int32(s[0])) == false {
		wordCount++
	}
	return wordCount
}



func main(){
	count := wordCounter("camelCaseMolly")
	fmt.Println(count)
}
