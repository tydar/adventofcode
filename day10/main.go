package main

import (
	"os"
	"bufio"
	"fmt"
	"sort"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	counters := map[string]int{")": 0, "]": 0, "}": 0, ">": 0,}
	scores := []int{0,}
	scIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		ok, stack, bad := validate("", line)
		if !ok {
			counters[bad]++
		} else {
			newScore := 0
			scIdx++
			for i := range stack {
				c := string(stack[i])
				newScore = newScore * 5
				switch c {
				case "(":
					newScore += 1
				case "[":
					newScore += 2
				case "{":
					newScore += 3
				case "<":
					newScore += 4
				}
			}
			scores = append(scores, newScore)
		}
	}
	sort.Ints(scores)
	middleScore := scores[(len(scores) / 2)]
	fmt.Printf(") %d ] %d } %d > %d\n", counters[")"], counters["]"], counters["}"], counters[">"])
	fmt.Printf("%d\n", (counters[")"] * 3) + (counters["]"] * 57) + (counters["}"] * 1197) + (counters[">"] * 25137))
	fmt.Printf("len(scores) = %d, scores = %v, middle = %d\n", len(scores), scores, middleScore)
}

func validate(stack, rest string) (bool, string, string) {
	if len(rest) == 0 {
		return true, stack, ""
	}

	next := string(rest[0])
	if isCloser(next) {
		if len(stack) == 0 {
			return false, stack, next
		} else if string(stack[0]) == matchingBracket(next) {
			return validate(stack[1:], rest[1:])
		} else {
			return false, stack, next
		}
	} else {
		return validate(next + stack, rest[1:])
	}
}

func isCloser(s string) bool {
	return s == "]" || s == ")" || s == "}" || s == ">"
}

func matchingBracket(s string) string {
	switch s {
	case "(":
		return ")"
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return "<"
	case ")":
		return "("
	case "]":
		return "["
	case "}":
		return "{"
	case ">":
		return "<"
	}
	return ""
}
