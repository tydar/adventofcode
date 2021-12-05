package main

import (
    "bufio"
    "strconv"
    "os"
    "strings"
    "fmt"
    "math"
)

// this is the solution to part 2
// comment out diagonal point gen for part 1
// this is also very slow. There are quicker algorithms
// but many assume working in a continuous problem space
// instead of discrete, and others assume things I don't think are true
// e.g. no line segments sharing endpoints.
// This one seems very interesting however: https://en.wikipedia.org/wiki/Bentley%E2%80%93Ottmann_algorithm

func main() {
    f, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(f)
    lss := make([]LineSegment, 0)
    for scanner.Scan() {
        line := scanner.Text()
        ptTokens := strings.Split(line, " -> ")

        startTok := strings.Split(ptTokens[0], ",")
        startX, err := strconv.Atoi(startTok[0])
        if err != nil {
            panic(err)
        }
        startY, err := strconv.Atoi(startTok[1])
        if err != nil {
            panic(err)
        }

        endTok := strings.Split(ptTokens[1], ",")
        endX, err := strconv.Atoi(endTok[0])
        if err != nil {
            panic(err)
        }
        endY, err := strconv.Atoi(endTok[1])
        if err != nil {
            panic(err)
        }

        start := Point{x: startX, y: startY,}
        end := Point{x: endX, y: endY,}
        ls := LineSegment{
            start: start,
            end:   end,
            points: getPoints(start, end),
        }
        lss = append(lss, ls)
    }

    results := GetAllIntersections(lss)

    count := 0
    for _, v := range results {
        if v > 1 {
            count++
        }
    }
    fmt.Printf("%d\n", count)
}

func GetAllIntersections(lss []LineSegment) map[Point]int {
    results := make(map[Point]int)
    for i := 0; i < len(lss) - 1; i++ {
        others := lss[i+1:]
        this := lss[i]
        for j := range others {
            other := others[j]
            if this.Intersects(other) {
                for _, pt := range this.Intersections(other) {
                    _, prs := results[pt]
                    if prs {
                        results[pt] += 1
                    } else {
                        results[pt] = 2
                    }
                }
            }
        }
    }
    return results
}

type Point struct {
    x, y int
}

func (p *Point) Measure() float64 {
    return math.Sqrt(float64((p.x * p.x) + (p.y * p.y)))
}

func SmallerBigger(p1, p2 Point) (Point, Point) {
    if p1.Measure() < p2.Measure() {
        return p1, p2
    } else {
        return p2, p1
    }
}

type LineSegment struct {
    start, end Point
    points     []Point
}

func getPoints(start, end Point) []Point {
    acc := make([]Point, 0)
    if start.x == end.x {
        // vertical line case
        if start.y > end.y {
            for i := end.y; i <= start.y; i++ {
                acc = append(acc, Point{x: start.x, y: i,})
            }
        } else {
            for i := start.y; i <= end.y; i++ {
                acc = append(acc, Point{x: start.x, y: i,})
            }
        }
    } else if start.y == end.y {
        //horizontal line case
        if start.x > end.x {
            for i := end.x; i <= start.x; i++ {
                acc = append(acc, Point{x: i, y: start.y,})
            }
        } else {
            for i := start.x; i <= end.x; i++ {
                acc = append(acc, Point{x: i, y: start.y,})
            }
        }
    } else {
        // diagonal line case
        if start.x < end.x && start.y < end.y {
            for i := 0; i + start.x <= end.x; i++ {
                acc = append(acc, Point{x: start.x + i, y: start.y + i,})
            }
        } else if start.x > end.x && start.y > end.y {
            for i := 0; start.x - i >= end.x; i++ {
                acc = append(acc, Point{x: start.x - i, y: start.y -i,})
            }
        } else if start.x > end.x && start.y < end.y {
            // sub from x add to y
            for i := 0; start.x - i >= end.x; i++ {
                acc = append(acc, Point{x: start.x - i, y: start.y + i,})
            }
        } else if start.x < end.x && start.y > end.y {
            // add to x and sub from y
            for i := 0; start.x + i <= end.x; i++ {
                acc = append(acc, Point{x: start.x + i, y: start.y - i},)
            }
        }
    }
    return acc
}

func (ls *LineSegment) Intersects(other LineSegment) bool {
    // I don't think the input is large enough that it will have a meaningful impact to filter out
    // before getting all the intersections
    return true
}

func (ls *LineSegment) Intersections(other LineSegment) []Point {
    acc := make([]Point, 0)
    for i := range ls.points {
        pt := ls.points[i]
        for j := range other.points {
            if pt == other.points[j] {
                acc = append(acc, pt)
            }
        }
    }
    return acc
}
