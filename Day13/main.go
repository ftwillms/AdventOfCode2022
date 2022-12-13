package main

import (
	"daythirteen/partone"
	"fmt"
	"sort"
)

// runPartOne find the shortest path from starting point to end point
func runPartOne() {
	fmt.Println("PART 1")
	input := partone.ReadInput("input.txt")
	pairs := partone.NewPacketNodePairs(input)
	sumOfOrderedPairs := 0
	lastPairToTest := pairs[len(pairs)-1]
	fmt.Println(lastPairToTest)
	for idx, pair := range pairs {
		if partone.Compare(pair.Left.Children, pair.Right.Children) < 1 {
			sumOfOrderedPairs += idx + 1
		}
	}
	// 6542 too low
	fmt.Println(sumOfOrderedPairs)
}

// runPartTwo
func runPartTwo() {
	fmt.Println("PART 2")
	firstDivider := partone.NewPacketNode("[[2]]")
	secondDivider := partone.NewPacketNode("[[6]]")
	allPackets := []*partone.PacketNode{firstDivider, secondDivider}
	input := partone.ReadInput("input.txt")
	pairs := partone.NewPacketNodePairs(input)
	for _, pair := range pairs {
		allPackets = append(allPackets, pair.Left, pair.Right)
	}
	// TIL golang has stdlib sorting...
	sort.Slice(allPackets, func(i, j int) bool {
		return partone.Compare(allPackets[i].Children, allPackets[j].Children) < 1
	})
	fmt.Println("Sorted")

	decoderKey := 1
	for i, packet := range allPackets {
		if packet == firstDivider || packet == secondDivider {
			// not 0 based indexing
			decoderKey *= i + 1
		}
	}
	fmt.Println("Decoder key: ", decoderKey)

}

func main() {
	fmt.Println("DAY 13")
	//runPartOne()
	runPartTwo()
}
