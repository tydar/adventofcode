package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
)

func main() {
    f, err := os.Open("input.full")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(f)
    scanner.Scan() // advance through first line

    var curr int
    last, scErr := strconv.Atoi(scanner.Text()) // get first line
    if scErr != nil {
        panic(scErr)
    }

    count := 0
    for scanner.Scan() {
        curr, scErr = strconv.Atoi(scanner.Text())
        if scErr != nil {
            panic(scErr)
        }
        if curr > last {
            count++
        }
        last = curr
    }
    fmt.Printf("%d increases\n", count)
}
