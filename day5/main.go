package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkNiceForVowel(s string) bool {
	vowelCount := strings.Count(s, "a") + strings.Count(s, "e") + strings.Count(s, "i") + strings.Count(s, "o") + strings.Count(s, "u")
	if vowelCount < 3 {
		return false
	}
	return true
}

func checkNiceForTwice(s string) bool {
	var previousChar rune
	twiceFound := false
	for _, c := range s {
		if previousChar == c {
			twiceFound = true
			break
		}
		previousChar = c
	}
	if !twiceFound {
		return false
	}
	return true
}

func checkNiceForSubstring(s string, sub string) bool {
	if strings.Contains(s, sub) {
		return false
	}
	return true
}

func checkNice(s string) bool {
	if checkNiceForVowel(s) && checkNiceForTwice(s) && checkNiceForSubstring(s, "ab") && checkNiceForSubstring(s, "cd") && checkNiceForSubstring(s, "pq") && checkNiceForSubstring(s, "xy") {
		return true
	}
	return false
}

func main() {
	dat, err := ioutil.ReadFile("./test.txt")
	check(err)
	dataStr := string(dat)
	dataArr := strings.Split(dataStr, "\n")
	niceStringCount := 0
	for i := 0; i < len(dataArr); i++ {
		if checkNice(dataArr[i]) {
			niceStringCount += 1
		}
	}
	fmt.Println("niceStringCount:", niceStringCount)
}
