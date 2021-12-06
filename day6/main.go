package main

import (
    "strings"
    "fmt"
    "bufio"
    "os"
    "strconv"
)

// naieve iterative-append version works just fine for 80 iterations
// but at 256 iterations you're likely to go OOM -- issue is space complexity
// however you can rewrite with O(1) space complexity (I think? it's been a while since school)
// by just unrolling it

func main() {
    f, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }

    school := map[int]int{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0,}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        list := scanner.Text()
        tok := strings.Split(list, ",")
        for i := range tok {
            n, err := strconv.Atoi(tok[i])
            if err != nil {
                panic(err)
            }
            school[n]++
        }
    }

    for i := 0; i < 256; i++ {
        new8 := school[0]
        new7 := school[8]
        new6 := school[0] + school[7]
        new5 := school[6]
        new4 := school[5]
        new3 := school[4]
        new2 := school[3]
        new1 := school[2]
        new0 := school[1]

        school[8] = new8
        school[7] = new7
        school[6] = new6
        school[5] = new5
        school[4] = new4
        school[3] = new3
        school[2] = new2
        school[1] = new1
        school[0] = new0
    }
    /*
    for j := 0; j < 256; j++ {
        count := len(lfs)
        for i := 0; i < count; i++ {
            if lfs[i] == 0 {
                lfs[i] = 6
                lfs = append(lfs, 8)
            } else {
                lfs[i]--
            }
        }
       // fmt.Printf("After day %d: %v\n", j, lfs)
    }
    */
    
    count := 0
    for i := 0; i < 9; i++ {
        count += school[i]
    }
    fmt.Printf("Count: %d\n", count)
}

func offspringCount(state, days int) int {
    count := 0
    daysAfterFirst := days - state - 1
    if daysAfterFirst > 0 {
        count += 1 
        count += offspringCount(8, daysAfterFirst)
    }
    for i := daysAfterFirst; i > 0; i -= 7 {
        if i > 0 {
            count += 1 
            count += offspringCount(8, i)
        }
    }
    return count 
}
