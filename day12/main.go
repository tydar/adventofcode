package main

// by following the state machine rules, we want to create an exhaustive list of valid paths
// invalid: does not get to end OR traverses the same "small cave" e.g. lower case state twice

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	pairs := make([]string, 0)
	for scanner.Scan() {
		pairs = append(pairs, scanner.Text())
	}

	transitions := transitionsFromPairs(pairs)
	startState := []Path{
		Path{p: []string{"end"}, small: true,},
	}
	paths := buildPaths(startState, transitions)
	fmt.Println(len(paths))
}

type Path struct {
	p []string
	small bool
}

func transitionsFromPairs(pairs []string) [][]string {
	acc := make([][]string, 0)
	for i := range pairs {
		states := strings.Split(pairs[i], "-")
		acc = append(acc, states)
		acc = append(acc, []string{states[1], states[0]})
	}
	return acc
}

func buildPaths(soFar []Path, transitions [][]string) []Path {
	acc := make ([]Path, 0)
	done := true
	for i := range soFar {
		path := soFar[i]
		mostRecent := path.p[0]
		if mostRecent == "start" {
			acc = append(acc, path)
		} else {
			possible := validFrom(mostRecent, transitions)
			for j := range possible {
				p := possible[j]
				if ok(p, path.p) {
					newPath := Path{p: append([]string{p,}, path.p...), small: path.small,}
					acc = append(acc, newPath)
					done = false
				} else if path.small && p != "end" && p != "start" {
					newPath := Path{p: append([]string{p,}, path.p...), small: false,}
					acc = append(acc, newPath)
					done = false
				}
			}
		}
	}
	if done {
		return acc
	} else {
		return buildPaths(acc, transitions)
	}
}

func validFrom(start string, transitions [][]string) []string {
	acc := make([]string, 0)
	for i := range transitions {
		if transitions[i][0] == start {
			acc = append(acc, transitions[i][1])
		}
	}
	return acc
}

func ok(next string, path []string) bool {
	if strings.ToLower(next) == next {
		return strings.Count(strings.Join(path,""), next) == 0
	} else {
		return true
	}
}
