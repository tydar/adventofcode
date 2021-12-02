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
    
    aim, position, depth := 0, 0, 0
    for scanner.Scan() {
        tokens := strings.Fields(scanner.Text())
        value, err := strconv.Atoi(tokens[1])
        if err != nil {
            log.Fatal(err)
        }
        
        switch tokens[0] {
        case "forward":
            position += value
            depth += aim * value
        case "down":
            aim += value
        case "up":
            aim -= value
        }
    }
    fmt.Printf("position %d * depth %d = %d\n", position, depth, position * depth)
}
