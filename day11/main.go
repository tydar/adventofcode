package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"reflect"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	board := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i := range line {
			n, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			row[i] = n
		}
		board = append(board, row)
	}

	flashed := make([][]int, 10)
	for i := 0; i < 10; i++ {
		row := make([]int, 10)
		for j := 0; j <10; j++ {
			row[j] = 0
		}
		flashed[i] = row
	}

	countVal := 0
	count := &countVal

	all := false
	step := 0
	for !all {
		Step(board, count)
		step++
		if reflect.DeepEqual(board, flashed) {
			all = true
		}
	}

	fmt.Printf("%d\n", step)
}

func possNeighbors(i, j int) ([]int, []int) {
	return []int{i-1, i-1, i-1, i, i, i, i + 1, i + 1, i + 1}, []int{j - 1, j, j + 1, j - 1, j,  j + 1, j - 1, j, j + 1}
}

func pointInbounds(i, j, w, h int) bool {
	return i >= 0 && i < h && j >= 0 && j < w
}

func Step(board [][]int, count *int) {
	// 1) add 1 to each cell
	// 2) flash any cells that are 9+
	// 2a) step through if flash says any new to flash
	// 3) after flashes settle, set all flashed cells to 0
	for i := range board {
		for j := range board[i] {
			board[i][j] += 1
		}
	}

	more := true
	for more {
		more = false
		for i := range board {
			for j := range board[i] {
				if board[i][j] > 9 {
					board[i][j] = -1
					Flash(i, j, board)
					*count += 1
				}
			}
		}
		for i := range board {
			for j := range board[i] {
				if board [i][j] > 9 {
					more = true
				}
			}
		}
	}
	ResetBoard(board)
}

func Flash(i, j int, board [][]int) {
	nI, nJ := possNeighbors(i, j)
	w, h := len(board[0]), len(board)
	for k := 0; k < len(nI); k++ {
		cI, cJ := nI[k], nJ[k]
		if pointInbounds(cI, cJ, w, h) {
			if board[cI][cJ] > -1 {
				board[cI][cJ] += 1
			}
		}
	}
}

func ResetBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == -1 {
				board[i][j] = 0
			}
		}
	}
}
