package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var grid [1000][1000]int

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

func parseXY(xyString string) (int, int) {
	xyArr := strings.Split(xyString, ",")

	return convertToInt(xyArr[0]), convertToInt(xyArr[1])
}

func processTurnCommand(commandArr []string) {
	startX, startY := parseXY(commandArr[2])
	endX, endY := parseXY(commandArr[4])

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			switch commandArr[1] {
			case "on":
				grid[i][j] += 1
			case "off":
				if grid[i][j] > 0 {
					grid[i][j] -= 1
				}
			}
		}
	}
}

func processToggle(commandArr []string) {
	startX, startY := parseXY(commandArr[1])
	endX, endY := parseXY(commandArr[3])

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			grid[i][j] += 2
		}
	}
}

func processCommand(command string) {
	commandArr := strings.Split(command, " ")
	switch commandArr[0] {
	case "turn":
		processTurnCommand(commandArr)
	case "toggle":
		processToggle(commandArr)
	default:
		panic("Invalid input data")
	}
}
func main() {
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}

	dat, err := ioutil.ReadFile("./test.txt")
	check(err)
	dataStr := string(dat)
	dataArr := strings.Split(dataStr, "\n")
	for i := 0; i < len(dataArr); i++ {
		processCommand(dataArr[i])
	}
	totalBrightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			totalBrightness += grid[i][j]
		}
	}
	fmt.Println("totalBrightness:", totalBrightness)
}
