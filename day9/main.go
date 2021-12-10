package main

// two approaches:
// 1) load everyone in and do naive calc spot by spot
// 2) attempt to calc as loading by having an accumulator per cell

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)


func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	values := make([][]int, len(lines))
	lineLen := len(lines[0])
	for i := range values {
		values[i] = make([]int, lineLen)
		for j := 0; j < lineLen; j++ {
			values[i][j], err = strconv.Atoi(string(lines[i][j]))
		}
	}

	minima := make([]Point, 0)

	for i := range values {
		row := values[i]
		for j := range row {
			n := values[i][j]
			//corner cases
			if i == 0 && j == 0 {
				if n < values[0][1] && n < values[1][0] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else if i == 0 && j == (lineLen - 1) {
				if n < values[0][j-1] && n < values[1][j] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else if i == len(values) - 1 && j == 0 {
				if n < values[i-1][0] && n < values[i][1] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else if i == len(values) - 1 && j == lineLen - 1 {
				if n < values[i-1][j] && n < values[i][j - 1] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else if i == 0 {
				// top
				if n < values[0][j - 1] && n < values[0][j + 1] && n < values[1][j] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else if i == len(values) - 1 {
				// bottom row
				if n < values[i][j-1] && n < values[i][j+1] && n < values[i-1][j] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else if j == 0 {
				// left column
				if n < values[i-1][0] && n < values[i][1] && n < values[i+1][0] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else if j == lineLen - 1 {
				// right column
				if n < values[i-1][j] && n < values[i][j-1] && n < values [i+1][j] {
					minima = append(minima, Point{x: i, y: j})
				}
			} else {
				if n < values[i-1][j] && n < values[i+1][j] && n < values[i][j-1] && n < values[i][j+1] {
					minima = append(minima, Point{x: i, y: j})
				}
			}
		}
	}
}

type Point struct {
	x, y int
}

func basinCalc(pt Point, values [][]int) int {
	// find the basin associated with a minima
	// and return the size

}

func getNeighbors(pt Point) {
//
}
