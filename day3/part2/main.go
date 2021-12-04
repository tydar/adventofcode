package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// The bit manipulation code in this is uglier than I would like
// I hard coded some things because I was convinced I had mis-written some things
// but turns out I had mis-read the problem ¯\_(ツ)_/¯

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	input := Input{
		nums: make([]uint64, 0),
	}
	indicies := make([]int, 0)
	count := 0
	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 2, 0)
		if err != nil {
			panic(err)
		}
		input.nums = append(input.nums, num)
		indicies = append(indicies, count)
		count++
	}

	mc := input.calcMostCommon(indicies, 11)
	ox := input.findRating(11, indicies, mc, 0, 1)
	co2 := input.findRating(11, indicies, mc, 1, 0)
	fmt.Printf("%d\n", ox*co2)
}

type Input struct {
	nums []uint64
}

func (in *Input) calcMostCommon(indicies []int, bitIdx int) uint64 {
	if bitIdx < 0 {
		panic("Shouldn't be here")
	}
	acc := uint64(0)
	count := 0
	for j := range indicies {
		num := in.nums[indicies[j]]
		shifted := num >> bitIdx
		masked := shifted &^ 0b111111111110
		acc += masked
		count += 1
	}
	if acc == (uint64(count) - acc) {
		fmt.Println("mc returning 2")
		return 2
	} else if acc > uint64(count/2) {
		return 1
	} else {
		return 0
	}
}

func (in *Input) findRating(bitIdx int, indicies []int, mostCommon uint64, xorVal uint64, keep uint64) uint64 {
	if len(indicies) == 1 {
		// if we've narrowed to one value, return it
		fmt.Printf("rating is %d at index %d\n", in.nums[indicies[0]], indicies[0])
		return in.nums[indicies[0]]
	} else if len(indicies) == 0 || bitIdx < 0 {
		log.Fatal(fmt.Sprintf("Code broke: index length: match  bitIdx %d\n", len(indicies), bitIdx))
	} else if len(indicies) > 0 && bitIdx >= 0 {
		// do the iteration
		indexAccumulator := make([]int, 0)
		for i := range indicies {
			num := in.nums[indicies[i]]
			shifted := num >> uint64(bitIdx)
			masked := shifted &^ 0b111111111110
			fmt.Printf("original: %b shifted: %b masked: %b mostCommon: %b xord: %b xorval: %b\n", num, shifted, masked, mostCommon, masked^mostCommon, xorVal)
			if mostCommon > 1 {
				if masked == keep {
					indexAccumulator = append(indexAccumulator, indicies[i])
				}
			} else if masked^mostCommon == xorVal {
				indexAccumulator = append(indexAccumulator, indicies[i])
			}
		}
		for i := range indexAccumulator {
			fmt.Printf("index: %d val: %b\n", indexAccumulator[i], in.nums[indexAccumulator[i]])
		}
		mc := uint64(0)
		if len(indexAccumulator) > 1 {
			mc = in.calcMostCommon(indexAccumulator, bitIdx-1)
		}
		return in.findRating(bitIdx-1, indexAccumulator, mc, xorVal, keep)
	}
	fmt.Println("Shouldn't be here")
	return 0
}
