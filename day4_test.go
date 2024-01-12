package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/cgeorgiades27/aoc23/utils"
)

func computeScratchValue(r io.Reader) int {
	scratchSum := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		cardSum := 0
		line := strings.Split(s.Text(), ":")[1]
		nums := strings.Split(line, "|")
		list, winners := strings.Split(strings.TrimSpace(nums[0]), " "), strings.Split(strings.TrimSpace(nums[1]), " ")

		m := make(map[int]int, len(list)+len(winners))
		for _, num := range list {
			v := utils.Atoi(num)
			if v != 0 {
				m[v]++
			}
		}
		for _, num := range winners {
			v := utils.Atoi(num)
			if v != 0 {
				m[v]--
			}
		}

		for _, v := range m {
			if v == 0 {
				if cardSum == 0 {
					cardSum = 1
				} else {
					cardSum *= 2
				}
			}
		}

		scratchSum += cardSum
	}

	return scratchSum
}

func computeScratchCopies(r io.Reader) int {
	totalCards, winningCards := buildTable(r)

	m := make(map[int]int)
	var recurser func(*int, int)
	recurser = func(sum *int, index int) {
		card := winningCards[index]
		if len(card) == 0 {
			return
		}

		*sum += len(card)
		for _, v := range card {
			m[v]++
			recurser(sum, v)
		}
	}

	sum := 0
	for i := 1; i <= totalCards; i++ {
		recurser(&sum, i)
	}

	return sum + totalCards
}

func buildTable(r io.Reader) (int, map[int][]int) {
	cards := make(map[int][]int)

	lineCount := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		lineCount++
		splt := strings.Split(s.Text(), ":")
		line := splt[1]
		cardNumber := utils.Atoi(strings.Split(splt[0], " ")[1])
		nums := strings.Split(line, "|")
		list, winners := strings.Split(strings.TrimSpace(nums[0]), " "), strings.Split(strings.TrimSpace(nums[1]), " ")

		m := make(map[int]int, len(list)+len(winners))
		for _, num := range list {
			v := utils.Atoi(num)
			if v != 0 {
				m[v]++
			}
		}
		for _, num := range winners {
			v := utils.Atoi(num)
			if v != 0 {
				m[v]--
			}
		}

		count := 1
		for _, v := range m {
			if v == 0 {
				cards[cardNumber] = append(cards[cardNumber], count+cardNumber)
				count++
			}
		}
	}

	return lineCount, cards
}

func TestComputeScratchCopies(t *testing.T) {
	inFile, _ := os.Open("infiles/day4.in")
	defer inFile.Close()

	testCases := []struct {
		reader   io.Reader
		expected int
	}{
		{
			reader: strings.NewReader(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53 
			Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
			Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
			Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
			Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
			Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`),
			expected: 30,
		},
		{
			reader:   bufio.NewReader(inFile),
			expected: 7386339,
		},
	}

	for _, test := range testCases {
		actual := computeScratchCopies(test.reader)
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}
func TestComputeScratchValue(t *testing.T) {
	inFile, _ := os.Open("infiles/day4.in")
	defer inFile.Close()

	testCases := []struct {
		reader   io.Reader
		expected int
	}{
		{
			reader: strings.NewReader(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53 
			Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
			Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
			Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
			Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
			Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`),
			expected: 13,
		},
		{
			reader:   bufio.NewReader(inFile),
			expected: 27845,
		},
	}

	for _, test := range testCases {
		actual := computeScratchValue(test.reader)
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}
