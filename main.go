package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1answer, err := SolvePart1("./input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("P1 answer = %d\n", p1answer)
}

type mySet map[string]struct{}

func (s mySet) has(value string) bool {
	_, ok := s[value]
	return ok
}
func (s mySet) add(value string) {
	s[value] = struct{}{}
}

func SolvePart1(path string) (int, error) {
	banks, err := readInput(path)
	if err != nil {
		return 0, err
	}
	iterations := 0
	history := mySet{}

	for {
		// create entry for repeat check and insert to set/map
		key := keyFromSlice(&banks)
		if history.has(key) {
			return iterations, nil
		}
		history.add(key)

		// distribute value
		distribute(&banks)
		iterations += 1
	}
}

func distribute(values *[]int) {
	index, largestValue := findFirstLargest(values)
	(*values)[index] = 0
	for i := 0; i < largestValue; i++ {
		// i iterator, index offset, +1 start at next memory bank slot for distribution
		(*values)[(i+index+1)%len(*values)] += 1
	}
}
func findFirstLargest(slice *[]int) (int, int) {
	largestValue := 0
	largestPos := 0 // hmmm

	for idx, value := range *slice {
		if value > largestValue {
			largestPos = idx
			largestValue = value
		}
	}
	return largestPos, largestValue
}

func keyFromSlice(slice *[]int) string {
	var str strings.Builder
	for i, v := range *slice {
		str.WriteString(strconv.Itoa(v))
		if i < len(*slice)-1 {
			str.WriteString(" ")
		}
	}
	return str.String()
}

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		text := scanner.Text()
		var values []int

		for _, v := range strings.Split(text, "\t") {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			values = append(values, i)
		}
		return values, nil
	}
	return nil, errors.New("unknown error")
}
