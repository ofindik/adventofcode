package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMD5Hex(input string) [md5.Size]byte {
	data := []byte(input)
	return md5.Sum(data)
}

func checkFirst5Bytes(input [md5.Size]byte) bool {
	hexStr := hex.EncodeToString(input[:])
	if strings.HasPrefix(hexStr, "000000") {
		return true
	}
	return false
}

func main() {
	prefix := os.Args[1]
	index := 1
	for {
		md5Hex := getMD5Hex(prefix + strconv.Itoa(index))
		if checkFirst5Bytes(md5Hex) {
			fmt.Println("Found: ", index)
			break
		}
		index += 1
		if index%1000000 == 0 {
			fmt.Print("+")
		}
	}
}
