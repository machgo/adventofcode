package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func check(i string) (string, bool) {
	data := []byte(i)
	sum := fmt.Sprintf("%x", md5.Sum(data))
	if string(sum[:5]) == "00000" {
		fmt.Println(sum)
		return sum, true
	}
	return sum, false
}

func main() {
	input := "reyedfim"
	found := 0
	pw := ""

	for i := 0; i < 1000000000; i++ {
		t := strconv.Itoa(i)
		sum, result := check(input + t)
		if result {
			pw += string(sum[5])
			found++
		}

		if found == 8 {
			fmt.Println(pw)
			break
		}
	}
}
