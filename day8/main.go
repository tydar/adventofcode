package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"
)

func main () {
    f, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(f)
    charVals := map[byte]uint8{'a': 64, 'b': 32, 'c': 16, 'd': 8, 'e': 4, 'f': 2, 'g': 1}

    acc := 0
    for scanner.Scan() {
        tok := strings.Split(scanner.Text(), "|")
        uniques := strings.Fields(tok[0])
        output := strings.Fields(tok[1])

        // set up our placeholders
        zero, one, two, three, four, five, six, seven, eight, nine := uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0)
        a, b, c, d, e, f, g := uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0)

        // set up sortation based on length
        uniquesSorted := map[int][]string{2: []string{}, 3: []string{}, 4: []string{}, 5: []string{}, 6: []string{}, 7: []string{}}

        for i := range uniques {
            seg := uniques[i]
            l := len(seg)
            uniquesSorted[l] = append(uniquesSorted[l], seg)
            ui := stringToUint(uniques[i], charVals)
            if l == 2 {
                // this is a one
                one = ui
            } else if l == 3 {
                seven = ui
            } else if l == 4 {
                four = ui
            } else if l == 7 {
                eight = ui
            }
        }

        a = seven ^ one
        for i := range uniquesSorted[6] {
            ui := stringToUint(uniquesSorted[6][i], charVals)
            if countBits(one ^ ui) == 6 {
                six = ui
            } else if countBits(four ^ ui) == 2 {
                nine = ui
            } else if countBits(four ^ ui) == 4 {
                zero = ui
            }
        }

        c = eight ^ six
        e = eight ^ nine
        d = eight ^ zero

        five = (eight ^ c) ^ e
        f = five & one
        b = ((four ^ d) ^ c) ^ f
        g = (((((eight ^ a) ^ b) ^ c) ^ d) ^ e) ^ f

        //have: zero, one, four, five, six, seven, eight, nine
        //need: two, three
        two = a | c | d | e | g
        three = a | c | d | f | g
        
        //now parse and add output
        for i := 1; i <= len(output); i++ {
            ui := stringToUint(output[len(output) - i], charVals)
            mult := intPow(10, i - 1)
            switch ui {
            case one:
                acc += 1 * mult
            case two:
                acc += 2 * mult
            case three:
                acc += 3 * mult
            case four:
                acc += 4 * mult
            case five:
                acc += 5 * mult
            case six:
                acc += 6 * mult
            case seven:
                acc += 7 * mult
            case eight:
                acc += 8 * mult
            case nine:
                acc += 9 * mult
            }
        }
    }


    // can immediately ID a by doing 7 ^ 1
    // ided: a

    // first we ID 6 by doing 1 ^ x
    // for x == each of the 6-bit options
    // the one that gives us 5 set bits is 6

    // 6 is 8 - c
    // so in binary rep if we do 8 ^ 6 we get c
    // ided: ac

    // 9 ^ 4 is 2 bits
    // 0 ^ 4 is 4 set bits
    // 9 is 8 - e
    // so in binary rep if we do 8 ^ 9 we get e
    // ided: ace

    // 0 is 8 - d
    // so in binary rep if we do 8 ^ 0 we get d
    // ided: acde

    // 5 is 8 - e - d
    // so we can do 8 ^ c ^ e
    // and ID f by doing 5 && 1
    // now have acedf so if we do 4 ^ d ^ c ^ f we get b
    // which gives us abcedf
    // and we can do 8 ^ abcedf to get g

    fmt.Println(acc)
}

func stringToUint(s string, valMap map[byte]uint8) uint8 {
    final := uint8(0)
    for i := range s {
        v, prs := valMap[s[i]]
        if !prs {
            log.Fatal("something didn't work in parsing")
        }
        final += v
    }
    return final
}

// algo from K&R, which I have a copy of somewhere
// but I got it from https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
func countBits(n uint8) int {
    count := 0
    for count = 0; n > 0; count++ {
        n = n & ( n -1 )
    }
    return count
}

func intPow(n, x int) int {
    if x == 0 {
        return 1
    } else {
        return n * intPow(n, x-1)
    }
}
