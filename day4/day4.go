package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	data     string
	sectorid int
	checksum string
}

// Pair is a simple key-value-struct used for sorting
type Pair struct {
	Key   string
	Value int
}

// PairList is the list
type PairList []Pair

func (p PairList) Len() int { return len(p) }

func (p PairList) Less(i, j int) bool {
	// sort per key if values are the same
	if p[i].Value == p[j].Value {
		return p[i].Key < p[j].Key
	}
	return p[i].Value > p[j].Value
}

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func loadData() ([]item, error) {
	dat, err := ioutil.ReadFile("day4/data4.txt")
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(dat))
	scanner := bufio.NewScanner(reader)

	orders := make([]item, 0)
	for scanner.Scan() {
		re := regexp.MustCompile(`([a-z\-]*)(\d*)\[([a-z]*)\]`)
		res := re.FindAllStringSubmatch(scanner.Text(), -1)

		data := strings.Replace(res[0][1], "-", "", -1)
		sectorid, _ := strconv.Atoi(res[0][2])
		checksum := res[0][3]

		i := item{data, sectorid, checksum}
		orders = append(orders, i)
	}

	return orders, nil
}

func (i *item) verify() bool {
	m := make(map[string]int)
	for _, o := range i.data {
		m[string(o)]++
	}
	p := make(PairList, len(m))
	j := 0
	for k, v := range m {
		p[j] = Pair{k, v}
		j++
	}
	sort.Sort(p)

	for h, c := range i.checksum {
		if p[h].Key != string(c) {
			return false
		}
	}

	return true
}

func main() {
	dat, err := loadData()
	if err != nil {
		panic("file not found")
	}

	sum := 0

	for _, i := range dat {
		if i.verify() {
			fmt.Printf("checksum correct %s\n", i.checksum)
			sum += i.sectorid
		} else {
			fmt.Printf("checksum incorrect %s\n", i.checksum)
		}
	}
	fmt.Printf("total sectorid: %d", sum)
}
