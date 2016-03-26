package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInt(arg string) int {
	result, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	return result
}

func getPaperAndRibbon(dimension string) (int, int) {
	if dimension == "" {
		return 0, 0
	}
	dimensions := strings.Split(dimension, "x")
	fmt.Println(dimensions)
	l := convertToInt(dimensions[0])
	w := convertToInt(dimensions[1])
	h := convertToInt(dimensions[2])
	area := (2 * l * w) + (2 * l * h) + (2 * h * w)
	fmt.Println("area:", area)
	dimensionSlice := []int{
		l,
		w,
		h,
	}
	sort.Ints(dimensionSlice)
	extra := dimensionSlice[0] * dimensionSlice[1]
	paper := area + extra
	ribbon := dimensionSlice[0] + dimensionSlice[0] + dimensionSlice[1] + dimensionSlice[1]
	bow := l * w * h
	ribbon = ribbon + bow
	return paper, ribbon
}

func main() {
	dat, err := ioutil.ReadFile("./test.txt")
	check(err)
	dataStr := string(dat)
	dataArr := strings.Split(dataStr, "\n")
	totalPaper, totalRibbon := 0, 0
	for i := 0; i < len(dataArr); i++ {
		paper, ribbon := getPaperAndRibbon(dataArr[i])
		totalPaper += paper
		totalRibbon += ribbon
	}
	fmt.Println("totalPaper:", totalPaper, " totalRibbon:", totalRibbon)
}
