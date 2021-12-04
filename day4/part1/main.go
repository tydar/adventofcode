package main

import (
    "bufio"
    "os"
    "strings"
    "strconv"
    "fmt"
    "errors"
)

// Part 2 only required ensuring boards are invalidated after their win
// so same file resolves both easily

func main() {
    f, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(f)
    ok := scanner.Scan()
    if !ok {
        panic("Something bad about input")
    }

    poolStr := scanner.Text() // first line is our string of draws
    poolTok := strings.Split(poolStr, ",")
    pool := make([]int, len(poolTok))
    for i := range poolTok {
        pool[i], err = strconv.Atoi(poolTok[i])
        if err != nil {
            panic(err)
        }
    }

    // now to read in and parse the boards
    acc := make([]string, 5) // this will accumulate 5 lines at a time
    boards := make([]Board, 0)
    count := 0
    scanner.Scan() // skip newline
    for scanner.Scan() {
        if count == 5 {
            board, err := ParseBoard(acc)
            if err != nil {
                panic(err)
            }
            boards = append(boards, board)
            count = 0
        } else {
            acc[count] = scanner.Text()
            count++
        }
    }

    validBoards := make([]bool, 0)
    for i := 0; i < len(boards); i++ {
        validBoards = append(validBoards, true)
    }

    for i := range pool {
        for j := range validBoards {
            if validBoards[j] {
                if boards[j].MarkBoard(pool[i]) {
                    fmt.Printf("winning score: %d\n", boards[j].Score() * pool[i])    
                    validBoards[j] = false
                }
            }
        }
    }
}

type Board struct {
    board [][]int
    marks [][]bool
}

func ParseBoard(input []string) (Board, error) {
    if len(input) != 5 {
        return Board{}, errors.New("ParseBoard: input length must be 5 lines")
    }

    board := Board {
        board: GenerateEmptyBoard(5, 5),
        marks: GenerateFalseMarks(5, 5),
    }


    for i := range input {
        tokens := strings.Fields(input[i])
        if len(tokens) != 5 {
            return Board{}, errors.New("ParseBoard: each line must have 5 numbers")
        }
        for j := range tokens {
            var err error
            board.board[i][j], err = strconv.Atoi(tokens[j])
            if err != nil {
                return Board{}, err
            }
        }
    }

    return board, nil
}

func (b *Board) MarkBoard(draw int) bool {
    x, y, prs := b.CheckNum(draw)
    if prs {
        b.marks[x][y] = true
    }
    return b.Win() 
}

func (b *Board) CheckNum(num int) (int, int, bool) {
    for i := range b.board {
        for j := range b.board[i] {
            if b.board[i][j] == num {
                return i, j, true
            }
        }
    }
    return -1, -1, false
}

func (b *Board) Win() bool {
    // check rows
    for i := range b.marks {
        acc := true
        for j := range b.marks[i] {
            acc = acc && b.marks[i][j]
        }
        if acc {
            return true
        }
    }

    // check columns
    for j := range b.marks[0] {
        acc := true
        for i := range b.marks {
            acc = acc && b.marks[i][j]
        }
        if acc {
            return true
        }
    }

    // check diagonals

    acc1, acc2 := true, true
    for i := range b.marks {
        acc1 = acc1 && b.marks[i][i]
        acc2 = acc2 && b.marks[i][len(b.marks)-i-1]
    }
    if acc1 || acc2 {
        return true
    }

    return false
}

func (b *Board) Score() int {
    score := 0
    for i := range b.board {
        for j := range b.board[i] {
            if !b.marks[i][j] {
                score += b.board[i][j]
            }
        }
    }
    return score
}

func GenerateFalseMarks(w, h int) [][]bool {
    m := make([][]bool, h)
    for i := 0; i < h; i++ {
        m[i] = make([]bool, w)
        for j := 0; j < w; j++ {
            m[i][j] = false
        }
    }
    return m
}

func GenerateEmptyBoard(w, h int) [][]int {
    m := make([][]int, h)
    for i := 0; i < h; i++ {
        m[i] = make([]int, w)
        for j := 0; j < w; j++ {
            m[i][j] = -1
        }
    }
    return m
}
