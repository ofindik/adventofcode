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

func checkNiceForAnyTwoTwice(s string) bool {
	twiceFound := false
	for i := 0; i < len(s)-1; i++ {
		previousStr := s[i : i+2]
		if strings.Count(s, previousStr) > 1 {
			twiceFound = true
			break
		}
	}
	return twiceFound
}

func checkNiceForRepeatWithOneLetter(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func checkNice(s string) bool {
	if checkNiceForAnyTwoTwice(s) && checkNiceForRepeatWithOneLetter(s) {
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
