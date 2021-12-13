package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func main () {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := make([]point, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] == 'f' {
			tok := strings.Fields(line)
			fold := strings.Split(tok[2], "=")
			if fold[0] == "y" {
				for i := range points {
					y, err := strconv.Atoi(fold[1])
					if err != nil {
						panic(err)
					}

					points[i] = foldUp(points[i], y)
				}
			} else if fold[0] == "x" {
				for i:= range points {
					x, err := strconv.Atoi(fold[1])
					if err != nil {
						panic(err)
					}

					points[i] = foldLeft(points[i], x)
				}
			}
		} else if len(line) > 0 {
			tok := strings.Split(line, ",")
			x, err := strconv.Atoi(tok[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(tok[1])
			if err != nil {
				panic(err)
			}

			points = append(points, point{x: x, y: y})
		}
	}

	w, h := xyMax(points)
	grid := make([][]string, h+1)
	for i := 0; i <= h; i++ {
		row := make([]string, w+1)
		for j := 0; j <= w; j ++ {
			row[j] = " "
		}
		grid[i] = row
	}

	for i := range points {
		p := points[i]
		grid[p.y][p.x] = "X"
	}

	printGrid(grid)
}

type point struct {
	x, y int
}

func foldUp(p point, line int) point {
	if p.y < line {
		return p
	} else {
		diff := p.y - line
		newY := line - diff
		return point{x: p.x, y: newY}
	}
}

func foldLeft(p point, line int) point {
	if p.x < line {
		return p
	} else {
		diff := p.x - line
		newX := line - diff
		return point{x: newX, y: p.y}
	}
}

func pointSearch(points []point, p point) bool {
	for i := range points {
		if p.x == points[i].x && p.y == points[i].y {
			return true
		}
	}
	return false
}

func xyMax(points []point) (int, int) {
	xMax, yMax := 0, 0
	for i := range points {
		if points[i].x > xMax {
			xMax = points[i].x
		}

		if points[i].y > yMax {
			yMax = points[i].y
		}
	}
	return xMax, yMax
}

func printGrid(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%s", grid[i][j])
		}
		fmt.Println("")
	}
}
