package main

import (
	"fmt"
	"os"
	"strconv"
)

func housePosition(x, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func main() {
	direction := os.Args[1]
	posX, posY := 0, 0
	posRoboX, posRoboY := 0, 0
	houseMap := make(map[string]int)
	houseMap[housePosition(posX, posY)] += 1
	houseMap[housePosition(posRoboX, posRoboY)] += 1
	for i := 0; i < len(direction); i++ {
		if i%2 == 1 {
			if direction[i] == '^' {
				posY += 1
			} else if direction[i] == '>' {
				posX += 1
			} else if direction[i] == '<' {
				posX -= 1
			} else if direction[i] == 'v' {
				posY -= 1
			}
			houseMap[housePosition(posX, posY)] += 1
		} else {
			if direction[i] == '^' {
				posRoboY += 1
			} else if direction[i] == '>' {
				posRoboX += 1
			} else if direction[i] == '<' {
				posRoboX -= 1
			} else if direction[i] == 'v' {
				posRoboY -= 1
			}
			houseMap[housePosition(posRoboX, posRoboY)] += 1
		}
	}
	fmt.Println("len:", len(houseMap))
	fmt.Println("houseMap:", houseMap)

}
