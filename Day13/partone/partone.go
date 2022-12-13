package partone

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type PacketNode struct {
	value    int
	Children []*PacketNode
}

type PacketPair struct {
	Left  *PacketNode
	Right *PacketNode
}

// Compare returns 1 when the right is larger, 0 when they are the same, and -1 when the left is smaller.
func Compare(left []*PacketNode, right []*PacketNode) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		leftVal := left[i]
		rightVal := right[i]
		if leftVal.Children == nil && rightVal.Children == nil {
			if leftVal.value < rightVal.value {
				return -1
			}
			if leftVal.value > rightVal.value {
				return 1
			}
		} else {
			leftChildren := leftVal.Children
			rightChildren := rightVal.Children

			if leftChildren == nil {
				leftChildren = []*PacketNode{leftVal}
			}
			if rightChildren == nil {
				rightChildren = []*PacketNode{rightVal}
			}
			diff := Compare(leftChildren, rightChildren)
			if diff != 0 {
				return diff
			}
		}
	}

	if len(left) > len(right) {
		return 1
	}
	if len(left) < len(right) {
		return -1
	}

	return 0
}

func ReadInput(filepath string) string {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Failed to parse file")
	}
	return string(bytes)
}

func NewPacketNode(input interface{}) *PacketNode {
	// this should only ever fire the first time
	if text, ok := input.(string); ok {
		var something interface{}
		json.Unmarshal([]byte(text), &something)
		return NewPacketNode(something)
	}

	iterable, ok := input.([]interface{})
	children := make([]*PacketNode, 0)
	if ok {
		for _, val := range iterable {
			if floatVal, ok := val.(float64); ok {
				children = append(children, &PacketNode{value: int(floatVal)})
			} else {
				children = append(children, NewPacketNode(val))
			}
		}
	}

	return &PacketNode{Children: children}
}

func NewPacketNodePairs(input string) []*PacketPair {
	allPairText := strings.Split(input, "\n\n")
	packetPairs := make([]*PacketPair, len(allPairText))
	for i, pairText := range allPairText {
		pairs := strings.Split(pairText, "\n")
		leftPair := NewPacketNode(pairs[0])
		rightPair := NewPacketNode(pairs[1])
		newPacketPair := &PacketPair{Left: leftPair, Right: rightPair}
		packetPairs[i] = newPacketPair
	}
	return packetPairs
}
