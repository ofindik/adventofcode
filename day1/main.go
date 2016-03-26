package main

import (
	"fmt"
	"os"
)

func main() {
	direction := os.Args[1]
	curFloor := 0
	pos := 0
	firstBasementPos := 0
	for i := 0; i < len(direction); i++ {
		if direction[i] == '(' {
			curFloor += 1
		} else if direction[i] == ')' {
			curFloor += -1
		}
		pos += 1
		if curFloor == -1 && firstBasementPos == 0 {
			firstBasementPos = pos
		}

	}
	fmt.Println("Floor:", curFloor)
	fmt.Println("firstBasementPos:", firstBasementPos)

}
