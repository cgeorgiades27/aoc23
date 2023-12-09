package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

type NetworkNode struct {
	Val   string
	Right string
	Left  string
}

func buildNodeGraph(r io.Reader) (string, string, map[string]NetworkNode) {
	s := bufio.NewScanner(r)
	network := make(map[string]NetworkNode)

	regex := regexp.MustCompile(`([A-Z]{3})`)

	s.Scan()
	directions := s.Text()
	count := 0
	first := ""

	// AAA = (BBB, CCC)
	for s.Scan() {
		line := s.Text()
		matches := regex.FindAllStringSubmatch(line, 3)
		if len(matches) == 0 || len(matches[0]) < 1 {
			continue
		}

		// find inital index
		if count == 0 {
			first = matches[0][0]
		}

		network[matches[0][0]] = NetworkNode{Val: matches[0][0], Left: matches[1][0], Right: matches[2][0]}
		count++
	}

	return first, directions, network
}

func findZZZ(first, directions string, network map[string]NetworkNode) int {
	directionIndex := 0
	var recurser func(int, string, *int)
	recurser = func(dirIndex int, index string, sum *int) {
		if index == "ZZZ" {
			return
		}

		currDir := directions[dirIndex]
		*sum++
		dirIndex++
		if dirIndex == len(directions) {
			dirIndex = 0
		}

		var nextDir string
		if string(currDir) == "L" {
			nextDir = network[index].Left
		} else {
			nextDir = network[index].Right
		}

		recurser(dirIndex, nextDir, sum)
	}

	sum := 0
	recurser(directionIndex, first, &sum)

	return sum
}

func TestFindZZZ(t *testing.T) {
	inFile, _ := os.Open("infiles/day8.in")
	r := bufio.NewReader(inFile)
	defer inFile.Close()

	testCases := []struct {
		reader   io.Reader
		expected int
	}{
		{
			reader: strings.NewReader(`RL

			AAA = (BBB, CCC)
			BBB = (DDD, EEE)
			CCC = (ZZZ, GGG)
			DDD = (DDD, DDD)
			EEE = (EEE, EEE)
			GGG = (GGG, GGG)
			ZZZ = (ZZZ, ZZZ)
			`),
			expected: 2,
		},
		{
			reader: strings.NewReader(`LLR

			AAA = (BBB, BBB)
			BBB = (AAA, ZZZ)
			ZZZ = (ZZZ, ZZZ)`),
			expected: 6,
		},
		{
			reader:   r,
			expected: 100,
		},
	}

	for _, test := range testCases {
		first, directions, network := buildNodeGraph(test.reader)
		actual := findZZZ(first, directions, network)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
