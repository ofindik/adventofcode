package main

import (
	"fmt"
	"os"
)

func main() {
	direction := os.Args[1]
	curFloor := 0
	for i := 0; i < len(direction); i++ {
		if direction[i] == '(' {
			curFloor += 1
		} else if direction[i] == ')' {
			curFloor += -1
		}
	}
	fmt.Println("Floor:", curFloor)

}
