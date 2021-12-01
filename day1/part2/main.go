package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
)


func sumFrame(fr []int) int {
    // we assume frame length is 3
    // because it better be
    return fr[0] + fr[1] + fr[2]
}

func main() {
    f, err := os.Open("input.full")
    if err != nil {
        panic(err)
    }

    input := make([]int, 0)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        val, cErr := strconv.Atoi(scanner.Text())
        if cErr != nil {
            panic(cErr)
        }
        input = append(input, val)
    }

    incCount := 0
    length := len(input)
    var lastFrameSum int
    if length >= 3 {
        firstFrame := input[0:3]
        lastFrameSum = sumFrame(firstFrame)
    } else {
        return
    }
    for i := 1; i < length && length >= 4; i++  {
        // check if we have enough for a frame
        if i + 2 < length {
            newFrame := input[i:i+3]
            newFrSum := sumFrame(newFrame)
            if newFrSum > lastFrameSum {
                incCount++
            }
            lastFrameSum = newFrSum
        }
    }

    fmt.Printf("%d increases\n", incCount)
}
