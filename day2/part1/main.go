package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"
)

func main() {
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    
    depth, position := 0, 0
    for scanner.Scan() {
        tokens := strings.Fields(scanner.Text())
        value, err := strconv.Atoi(tokens[1])
        if err != nil {
            log.Fatal(err)
        }
        
        switch tokens[0] {
        case "forward":
            position += value
        case "down":
            depth += value
        case "up":
            depth -= value
        }
    }
    fmt.Printf("position %d * depth %d = %d\n", position, depth, position * depth)
}
