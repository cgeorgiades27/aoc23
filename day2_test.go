package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var colorQuantities = map[string]int{
	"red": 12, "green": 13, "blue": 14,
}

func checkGame(r io.Reader) int {
	idSum := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.Split(s.Text(), ":")
		if len(line) < 1 {
			continue
		}

		id, err := strconv.Atoi(strings.Split(line[0], " ")[1])
		if err != nil {
			continue
		}

		isUnder := true
		for _, cube := range strings.Split(line[1], ";") {
			for _, color := range strings.Split(cube, ",") {
				set := strings.Split(strings.TrimLeft(color, " "), " ")
				val, err := strconv.Atoi(set[0])
				if err != nil {
					continue
				}
				if val > colorQuantities[set[1]] {
					isUnder = false
					break
				}
			}
		}
		if isUnder {
			idSum += id
		}

	}

	return idSum
}

func powerSet(r io.Reader) int {
	idSum := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.Split(s.Text(), ":")
		if len(line) < 1 {
			continue
		}

		m := make(map[string]int, 3)
		for _, cube := range strings.Split(line[1], ";") {
			for _, color := range strings.Split(cube, ",") {
				set := strings.Split(strings.TrimLeft(color, " "), " ")
				val, err := strconv.Atoi(set[0])
				if err != nil {
					continue
				}
				if val > m[set[1]] {
					m[set[1]] = val
				}
			}
		}

		pSet := 1
		for _, v := range m {
			pSet *= v
		}

		idSum += pSet
	}

	return idSum
}

func TestPowerSet(t *testing.T) {
	inFile, err := os.Open("./day2.in")
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}
	defer inFile.Close()

	testCases := []struct {
		reader   io.Reader
		expected int
	}{
		{
			reader: strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
			Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
			Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
			Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
			Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`),
			expected: 2286,
		},
		{
			reader:   bufio.NewReader(inFile),
			expected: 66909,
		},
	}

	for _, test := range testCases {
		actual := powerSet(test.reader)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestCheckGame(t *testing.T) {
	inFile, err := os.Open("./day2.in")
	if err != nil {
		t.Error(err)
	}
	defer inFile.Close()

	testCases := []struct {
		reader   io.Reader
		expected int
	}{
		{
			reader: strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
			Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
			Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
			Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
			Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`),
			expected: 8,
		},
		{
			reader:   bufio.NewReader(inFile),
			expected: 2156,
		},
	}

	for _, test := range testCases {
		actual := checkGame(test.reader)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
