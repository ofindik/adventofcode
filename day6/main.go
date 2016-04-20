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

func processTurnCommand(commandArr []string) {
	start := commandArr[2]
	startArr := strings.Split(start, ",")
	startX := convertToInt(startArr[0])
	startY := convertToInt(startArr[1])

	end := commandArr[4]
	endArr := strings.Split(end, ",")
	endX := convertToInt(endArr[0])
	endY := convertToInt(endArr[1])

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			switch commandArr[1] {
			case "on":
				grid[i][j] = 1
			case "off":
				grid[i][j] = 0
			}
		}
	}
}

func processToggle(commandArr []string) {
	start := commandArr[1]
	startArr := strings.Split(start, ",")
	startX := convertToInt(startArr[0])
	startY := convertToInt(startArr[1])

	end := commandArr[3]
	endArr := strings.Split(end, ",")
	endX := convertToInt(endArr[0])
	endY := convertToInt(endArr[1])

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			if grid[i][j] == 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
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
	litLights := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] == 1 {
				litLights++
			}
		}
	}
	fmt.Println("litLights:", litLights)
}
