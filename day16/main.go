package main

import (
	"fmt"
	"strconv"
)

func main() {
	lookup := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	input := "D2FE28"
	base2 := ""
	for i := range input {
		c := string(input[i])
		bits := lookup[c]
		base2 = base2 + bits
	}
	fmt.Println(base2)
}

func versionFromPacket(packet string) uint64 {
	v, err := strconv.ParseUint(packet[:3], 2, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func operatorPacket(packet string) bool {
	t, err := strconv.ParseUint(packet[3:6], 2, 64)
	if err != nil {
		panic(err)
	}
	return t != 4
}

func getSubpackets(packet string) string {
	// assumes we've checked it's an operator packet
	lenTypeBit := packet[6]
	if lenTypeBit == '0' {
		// 15 bits = length of subpackets in bits
		lBits := packet[7:22]
		l, err := strconv.ParseUint(lBits, 2, 64)
		if err != nil {
			panic(err)
		}
		return packet[22 : 22+l]
	} else {
		// 11 bits = number of subpackets
		// lBits := packet[7:18]
		return packet[18:]
	}
}
