package partone

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ReadInput(t *testing.T) {
	input := ReadInput("../test-input.txt")
	pairs := NewPacketNodePairs(input)
	require.Equal(t, 8, len(pairs))
}

func Test_Compare(t *testing.T) {
	cases := []struct {
		name    string
		outcome bool
		packetA *PacketNode
		packetB *PacketNode
	}{
		{
			name:    "foo",
			packetA: NewPacketNode("[1,1,3,1,1]"),
			packetB: NewPacketNode("[1,1,5,1,1]"),
			outcome: true,
		},
		{
			name:    "bar",
			packetA: NewPacketNode("[[1],[2,3,4]]"),
			packetB: NewPacketNode("[[1],4]"),
			outcome: true,
		},
		{
			name:    "cat",
			packetA: NewPacketNode("[9]"),
			packetB: NewPacketNode("[[8,7,6]]"),
			outcome: false,
		},
		{
			name:    "delta",
			packetA: NewPacketNode("[]"),
			packetB: NewPacketNode("[3]"),
			outcome: true,
		},
		{
			name:    "echo",
			packetA: NewPacketNode("[[[]]]"),
			packetB: NewPacketNode("[[]]"),
			outcome: false,
		},
		{
			name:    "f",
			packetA: NewPacketNode("[7,7,7,7]"),
			packetB: NewPacketNode("[7,7,7]"),
			outcome: false,
		},
		{
			name:    "g",
			packetA: NewPacketNode("[[4,4],4,4]"),
			packetB: NewPacketNode("[[4,4],4,4,4]"),
			outcome: true,
		},
		{
			name:    "gross",
			packetA: NewPacketNode("[1,[2,[3,[4,[5,6,7]]]],8,9]"),
			packetB: NewPacketNode("[1,[2,[3,[4,[5,6,0]]]],8,9]"),
			outcome: false,
		},
		{
			name:    "???",
			packetA: NewPacketNode("[[[[1,7,9,2,2],[10,9,9,2]],[5,[],1]],[[],3,1,[[1,8,2],1,[]],7],[[[],3,1,1,3],[5,[9,6,7,7,0],10],1,5,[3,[6,5,7,5,10],[4,2],[10,9,8,5,9],[1,9,5]]]]"),
			packetB: NewPacketNode("[[[4,10,[0],0],5],[4,10,2,[0],[7]],[[4,[9,2]],6,5,7,8]]"),
			outcome: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			compared := Compare(tc.packetA.Children, tc.packetB.Children)
			require.Equal(t, tc.outcome, compared < 1)
		})
	}
}

func Test_NewPacketNode(t *testing.T) {
	test := NewPacketNode("[1,2,3,4,5,6]")
	fmt.Println(test)

	test = NewPacketNode("[]")
	fmt.Println(test)

	test = NewPacketNode("[1,2,3,[4,5,6]]")
	fmt.Println(test)
}
