package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/cgeorgiades27/aoc23/utils"
)

func buildMatrix(r io.Reader) [][]rune {
	var mtx [][]rune
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		mtx = append(mtx, row)

	}

	return mtx
}

func computeSchematicSum(schematic [][]rune) int {
	sum := 0
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			var buff bytes.Buffer
			symbol := false
			for j < len(schematic[i]) && utils.IsDigit(schematic[i][j]) {
				symbol = symbol || hasSymbol360(i, j, schematic)
				buff.WriteRune(schematic[i][j])
				j++
			}
			if symbol {
				sum += utils.Atoi(buff.String())
			}
		}
	}
	return sum
}

func hasSymbol360(i, j int, schematic [][]rune) bool {
	up := schematic[max(i-1, 0)][j]
	down := schematic[min(i+1, len(schematic)-1)][j]
	left := schematic[i][max(j-1, 0)]
	right := schematic[i][min(j+1, len(schematic[i])-1)]
	upLeft := schematic[max(i-1, 0)][max(j-1, 0)]
	upRight := schematic[min(i+1, len(schematic)-1)][max(j-1, 0)]
	bottomLeft := schematic[max(i-1, 0)][min(j+1, len(schematic[i])-1)]
	bottomRight := schematic[min(i+1, len(schematic)-1)][min(j+1, len(schematic[i])-1)]

	runes := []rune{up, down, left, right, upLeft, upRight, bottomLeft, bottomRight}
	for _, r := range runes {
		if !utils.IsDigit(r) && r != '.' {
			return true
		}
	}

	return false
}

func TestComputeSchematicSum(t *testing.T) {
	inEx, _ := os.Open("infiles/day3-ex.in")
	defer inEx.Close()
	in, _ := os.Open("infiles/day3.in")
	defer in.Close()

	var testCases = []struct {
		mtx      [][]rune
		expected int
	}{
		{
			mtx:      buildMatrix(bufio.NewReader(inEx)),
			expected: 4361,
		},
		{
			mtx:      buildMatrix(bufio.NewReader(in)),
			expected: 527446,
		},
	}

	for _, tc := range testCases {
		actual := computeSchematicSum(tc.mtx)
		if actual != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, actual)
		}
	}
}
