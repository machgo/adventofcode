package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

func loadData() ([]string, error) {
	dat, err := ioutil.ReadFile("day2/data2.txt")
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(dat))
	scanner := bufio.NewScanner(reader)

	orders := make([]string, 0)
	for scanner.Scan() {
		orders = append(orders, scanner.Text())
	}

	return orders, nil
}

func printNumber(x int, y int, command string) (int, int) {
	numbers := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, r := range command {
		direction := string(r)
		if direction == "U" && y > 0 {
			y--
		}
		if direction == "D" && y < 2 {
			y++
		}
		if direction == "L" && x > 0 {
			x--
		}
		if direction == "R" && x < 2 {
			x++
		}
		//fmt.Printf("command: %s new number: %d\n", direction, numbers[y][x])
	}
	fmt.Printf("%d", numbers[y][x])
	return x, y
}

func main() {
	dat, err := loadData()
	if err != nil {
		panic(0)
	}
	x, y := 1, 1
	for i := range dat {
		x, y = printNumber(x, y, dat[i])
	}
}
