package day01

import (
	"bufio"
	"os"
	"strconv"
)

func readLines(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}

	return lines
}

func sumFrequencies(frequencies []int) int {
	sum := 0
	for i := 0; i < len(frequencies); i++ {
		sum += frequencies[i]
	}
	return sum
}

func firstDupe(frequencies []int) int {
	uniq := map[int]bool{
		0: true,
	}
	freqLen := len(frequencies)
	curFreq := 0
	for i := 0; ; i = (i + 1) % freqLen {
		curFreq += frequencies[i]
		_, exists := uniq[curFreq]
		if exists {
			return curFreq
		}
		uniq[curFreq] = true
	}
}
