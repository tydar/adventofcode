package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// did have to check out r/adventofcode bc I was stuck. Just seeing the words "only care about the pairs" helped

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	start := scanner.Text()
	scanner.Scan() // should give us the empty line

	charCounts := make(map[string]int)
	for i := range start {
		charCounts[string(start[i])] = 0
	}

	pairs := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		tok := strings.Split(line, " -> ")
		pairs[tok[0]] = tok[1]
		c1, c2, c3 := string(tok[0][0]), string(tok[0][1]), tok[1]
		charCounts[c1] = 0
		charCounts[c2] = 0
		charCounts[c3] = 0
	}

	counts := countPairs(start)
	for i := 0; i < 40; i++ {
		counts = Step(counts, pairs)
	}

	for k, v := range counts {
		c1, c2 := string(k[0]), string(k[1])
		charCounts[c1] += v
		charCounts[c2] += v
	}

	for k := range charCounts {
		charCounts[k] = (charCounts[k] / 2) + (charCounts[k] % 2)
	}

	fmt.Println(charCounts)
	fmt.Println(minmaxDiff(charCounts))
}

func countPairs(start string) map[string]int {
	acc := make(map[string]int)
	for i := 0; i < len(start)-1; i++ {
		pair := start[i : i+2]
		_, ok := acc[pair]
		if ok {
			acc[pair] += 1
		} else {
			acc[pair] = 1
		}
	}
	return acc
}

func minmaxDiff(m map[string]int) int {
	min, max := -1, -1
	for _, v := range m {
		if min == -1 {
			min = v
		}

		if max == -1 {
			max = v
		}

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}
	return max - min
}

func Step(counts map[string]int, rules map[string]string) map[string]int {
	newCounts := make(map[string]int)
	for k := range counts {
		newCounts[k] = 0
	}
	for k, v := range counts {
		c := rules[k]
		p1, p2 := string(k[0])+c, c+string(k[1])
		newCounts[p1] += v
		newCounts[p2] += v
	}
	return newCounts
}
