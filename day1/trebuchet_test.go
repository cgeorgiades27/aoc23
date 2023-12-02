package day1

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"testing"
)

type pair struct {
	index int
	word  string
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func trebuchet(r io.Reader) int {
	m := map[string]rune{
		"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9',
	}

	sum := 0
	s := bufio.NewScanner(r)

	for s.Scan() {
		first, last := 0, 0
		line := s.Text()

		min, max := pair{index: len(line), word: ""}, pair{index: 0, word: ""}
		for k := range m {
			firstIndex := strings.Index(line, k)
			lastIndex := strings.LastIndex(line, k)

			if firstIndex < min.index && firstIndex >= 0 {
				min.index = firstIndex
				min.word = k
			}

			if lastIndex > max.index {
				max.index = lastIndex
				max.word = k
			}
		}

		line = strings.Replace(line, min.word, string(m[min.word]), 1)
		line = strings.Replace(line, max.word, string(m[max.word]), -1)

		set := false
		for i, rn := range line {
			if isDigit(rn) {
				if !set {
					first = i
					last = i
					set = true
				} else {
					last = i
				}
			}
		}

		calVal, err := strconv.Atoi(string(line[first]) + string(line[last]))
		if err != nil {
			calVal = 0
		}

		sum += calVal
	}

	return sum
}

func TestTrebuchet(t *testing.T) {
	testCases := []struct {
		reader   io.Reader
		expected int
	}{
		{
			reader: strings.NewReader(`1abc2
			pqr3stu8vwx
			a1b2c3d4e5f
			treb7uchet`),
			expected: 142,
		},
		{
			reader: strings.NewReader(`two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen`),
			expected: 281,
		},
	}

	for _, test := range testCases {
		actual := trebuchet(test.reader)
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}

}

func TestIsDigit(t *testing.T) {
	testCases := []struct {
		value    rune
		expected bool
	}{
		{
			value:    '1',
			expected: true,
		},
		{
			value:    '2',
			expected: true,
		},
		{
			value:    '9',
			expected: true,
		},
		{
			value:    'a',
			expected: false,
		},
		{
			value:    'A',
			expected: false,
		},
		{
			value:    'z',
			expected: false,
		},
		{
			value:    'Z',
			expected: false,
		},
	}

	for _, test := range testCases {
		actual := isDigit(test.value)
		if actual != test.expected {
			t.Errorf("expected %t, got %t", test.expected, actual)
		}
	}
}
