package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    f, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(f)
    starts := make(map[int]int)

    scanner.Scan()
    line := scanner.Text()
    tok := strings.Split(line, ",")
    min, max := -1, -1
    for i := range tok {
        n, err := strconv.Atoi(tok[i])
        if err != nil {
            panic(err)
        }
        if min < 0 || n < min {
            min = n
        } else if n > max {
            max = n
        }

        _, prs := starts[n]
        if prs {
            starts[n] += 1
        } else {
            starts[n] = 1
        }
    }

    minCost := CalcCost(starts, min)
    for i := min+1; i <= max; i++ {
        cost := CalcCost(starts, i)
        if cost < minCost {
            minCost = cost
        }
    }
    fmt.Printf("%d\n", minCost)
}

// CalcCost takes a map of positions -> number of crabs starting there
// for part 2, cost is no longer |pos - start| but instead T_|pos - start|
// the triangular number
func CalcCost(starts map[int]int, position int) int {
    acc := 0
    for k, v := range starts {
        if k >= position {
            acc += Combinations(k - position + 1, 2) * v
        } else {
            acc += Combinations(position - k + 1, 2) * v
        }
    }
    return acc
}

// function to calculate n C k
// taken from stack overflow
// which says it was taken from Knuth
func Combinations(n, k int) int {
    if (k > n) {
        return 0
    }

    r := 1
    for d := 1; d <= k; d++ {
        r *= n
        n--
        r /= d
    }
    return r
}
