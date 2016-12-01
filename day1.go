package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type step struct {
	direction string
	distance  int
}

func (s *step) String() string {
	return fmt.Sprintf("direction: %s   distance: %d\n", s.direction, s.distance)
}

func getDataFile() string {
	dat, err := ioutil.ReadFile("data1.txt")
	if err != nil {
		return "error"
	}
	return string(dat)
}

func parseData(data string) []step {
	result := strings.Split(data, ",")
	resArr := make([]step, 1)

	for i := range result {
		d := strings.TrimSpace(result[i])
		stepRx := regexp.MustCompile(`(R|L)(\d*)`)
		matches := stepRx.FindAllSubmatch([]byte(d), -1)
		direction := string(matches[0][1])
		distance, _ := strconv.Atoi(string(matches[0][2]))
		a := step{direction, distance}
		resArr = append(resArr, a)
	}

	return resArr
}

func calc(input []step) (int, int) {
	var traveled [4]int
	lastDirection := 0 // 0 = north, 1 = east, 2 = south, 3 = west

	for i := range input {
		if input[i].direction == "L" {
			lastDirection--
		} else {
			lastDirection++
		}

		if lastDirection == -1 {
			lastDirection = 3
		}
		if lastDirection == 4 {
			lastDirection = 0
		}
		traveled[lastDirection] += input[i].distance
	}
	return traveled[0] - traveled[2], traveled[1] - traveled[3]
}

func main() {
	d := getDataFile()
	steps := parseData(d)
	north, east := calc(steps)
	fmt.Printf("north: %d east: %d totaldistance: %d", north, east, north+east)
}
