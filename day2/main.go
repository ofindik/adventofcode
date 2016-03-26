package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	l := parseArg(os.Args[1])
	w := parseArg(os.Args[2])
	h := parseArg(os.Args[3])
	area := (2 * l * w) + (2 * l * h) + (2 * h * w)
	fmt.Println("area:", area)
	dimension := []int{
		l,
		w,
		h,
	}
	sort.Ints(dimension)
	fmt.Println("dimension:", dimension)
	extra := dimension[0] * dimension[1]
	fmt.Println("extra:", extra)
	total := area + extra
	fmt.Println("total:", total)
}

func parseArg(arg string) int {
	result, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	return result
}
