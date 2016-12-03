package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func loadData() ([]string, error) {
	dat, err := ioutil.ReadFile("day3/data3.txt")
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

func checkTriangle(data string) bool {
	re := regexp.MustCompile(`\s*(\d*)\s*(\d*)\s*(\d*)`)
	res := re.FindAllStringSubmatch(data, -1)
	a, _ := strconv.Atoi(res[0][1])
	b, _ := strconv.Atoi(res[0][2])
	c, _ := strconv.Atoi(res[0][3])

	if a < (b+c) && b < (a+c) && c < (a+b) {
		fmt.Printf("possible: %d %d %d\n", a, b, c)
		return true
	}

	fmt.Printf("not possible: %d %d %d\n", a, b, c)
	return false
}

func main() {
	dat, err := loadData()
	if err != nil {
		panic("file not found")
	}
	possible := 0
	for _, o := range dat {
		if checkTriangle(o) {
			possible++
		}
	}
	fmt.Printf("possible: %d\nnotpossible: %d\ntotal: %d\n", possible, len(dat)-possible, len(dat))
}
