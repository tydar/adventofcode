package main

import (
    "strconv"
    "fmt"
    "bufio"
    "os"
)

func main() {
    // read number
    // xor against mask
    // bitshift right
    // add to accumulator, counter
    // repeat

    f, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(f)
    acc := make([]uint64, 12) // counting 1s
    count := 0
    for scanner.Scan() {
        num, err := strconv.ParseUint(scanner.Text(), 2, 0)
        if err != nil {
            panic(err)
        }
        for i := 1; i <= 12; i++ {
            masked := num &^ 0b111111111110 // 1 iff bit in question is 1
            acc[12-i] += masked
            num = num>>1
        }
        count += 1
    }

    var gamma uint64
    var epsilon uint64
    gamma, epsilon = 0, 0
    for i := 0; i < 12; i++ {
        if acc[i] > uint64(count / 2) {
            gamma += 1
        } else {
            epsilon += 1
        }
        if i < 11 {
            gamma = gamma<<1
            epsilon = epsilon<<1
        }
    }
    fmt.Printf("%d\n", gamma * epsilon)
}
