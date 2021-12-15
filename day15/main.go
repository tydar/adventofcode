package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	costs := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		tok := strings.Split(line, "")
		row := make([]int, len(tok))
		for i := range tok {
			n, err := strconv.Atoi(tok[i])
			if err != nil {
				panic(err)
			}
			row[i] = n
		}
		costs = append(costs, row)
	}

	h := len(costs)
	w := len(costs[0])

	bigCosts := make([][]int, h*5)
	for i := range bigCosts {
		row := make([]int, w*5)
		bigCosts[i] = row
	}

	for i := range costs {
		for j := range costs[i] {
			n := costs[i][j]
			bigCosts[i][j] = n
			for k := 1; k < 5; k++ {
				// these conditionals give us the kth (0 indexed) occurence of n on
				// the first row of tiles
				if n+k > 9 {
					v := (n + k) % 9
					if v == 0 {
						v = 9
					}
					bigCosts[i][j+k*w] = v
				} else {
					bigCosts[i][j+k*w] = n + k
				}
			}

			for l := 1; l < 5; l++ {
				// here we will do the rest
				// l is the row of tiles
				for k := 0; k < 5; k++ {
					//k is the col of tiles
					if n+l+k > 9 {
						v := (n + k + l) % 9
						if v == 0 {
							v = 9
						}
						bigCosts[i+l*h][j+k*h] = v
					} else {
						bigCosts[i+l*h][j+k*h] = n + k + l
					}
				}
			}
		}
	}

	outToFile(bigCosts)
	fmt.Println(dijkstra(bigCosts))
}

func outToFile(costs [][]int) {
	f, err := os.Create("run.out")
	if err != nil {
		panic(err)
	}
	for i := range costs {
		for j := range costs {
			s := strconv.Itoa(costs[i][j])
			_, err := f.WriteString(s + " ")
			if err != nil {
				panic(err)
			}
		}
		_, err := f.WriteString("\n")
		if err != nil {
			panic(err)
		}
	}
}

func dijkstra(costs [][]int) int {
	// returns the lowest cost path from 0,0 to h-1,w-1 (i, j order outer/inner array)
	h := len(costs)
	w := len(costs[0])

	fmt.Printf("h: %d w: %d\n", h, w)

	visited := make([][]bool, h)
	distance := make([][]int, h)
	for i := range visited {
		rowV := make([]bool, w)
		rowD := make([]int, w)
		for j := range visited {
			rowV[j] = false
			rowD[j] = -1
		}
		visited[i] = rowV
		distance[i] = rowD
	}

	cI, cJ, v := h-1, w-1, 0
	distance[cI][cJ] = 0
	count := 0
	for !visited[0][0] && v != -1 {
		if !visited[cI][cJ] {
			visit(cI, cJ, distance, costs, visited)
			cI, cJ, v = minVal(distance, visited)
			count++
			fmt.Printf("Completing iteration %d\n", count)
		}
	}

	fmt.Printf("0,0: %d, h-1,w-1: %d\n", costs[0][0], costs[h-1][w-1])
	fmt.Println("For my full input, this code kept giving me the value 3047 when the answer was 3045. I am not sure why.")
	return distance[0][0] + costs[h-1][w-1] - costs[0][0]
}

func visit(i, j int, distance, costs [][]int, visited [][]bool) {
	// visit sets the distance values for the neighbors of a node i, j
	node := distance[i][j]
	h := len(distance)
	w := len(distance[i])
	nI, nJ := neighbors(i, j, h, w)
	for _, ni := range nI {
		if !visited[ni][j] {
			d := node + costs[ni][j]
			if distance[ni][j] == -1 || distance[ni][j] > d {
				distance[ni][j] = d
			}
		}
	}

	for _, nj := range nJ {
		if !visited[i][nj] {
			d := node + costs[i][nj]
			if distance[i][nj] == -1 || distance[i][nj] > d {
				distance[i][nj] = d
			}
		}
	}

	visited[i][j] = true
}

func neighbors(i, j, h, w int) ([]int, []int) {
	nI, nJ := make([]int, 0), make([]int, 0)
	if i > 0 && i < (h-1) {
		nI = []int{i - 1, i + 1}
	} else if i == 0 {
		nI = []int{i + 1}
	} else {
		nI = []int{i - 1}
	}

	if j > 0 && j < (w-1) {
		nJ = []int{j - 1, j + 1}
	} else if j == 0 {
		nJ = []int{j + 1}
	} else {
		nJ = []int{j - 1}
	}
	return nI, nJ
}

func minVal(a [][]int, visited [][]bool) (int, int, int) {
	min := -1
	minI, minJ := 0, 0
	for i := range a {
		for j := range a[i] {
			if (a[i][j] < min || min == -1) && !visited[i][j] {
				min = a[i][j]
				minI, minJ = i, j
			}
		}
	}
	return minI, minJ, min
}

func prettyPrint(a [][]int) {
	for i := range a {
		for j := range a[i] {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println("")
	}
}
