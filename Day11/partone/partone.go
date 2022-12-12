package partone

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"log"
)

// Part 2
const DIVISOR int64 = 5 * 17 * 7 * 13 * 19 * 3 * 11 * 2

type OperationFunc func(old int64) int64
type TestFunc func(val int64) bool

type Monkey struct {
	itemQueue       *queue.Queue
	inspectionCount int64
	operation       OperationFunc
	testFunc        TestFunc
	falseMonkey     *Monkey
	trueMonkey      *Monkey
	divisor         int
}

func NewMonkey(
	startingItems []int64, operation OperationFunc, testFunc TestFunc,
) *Monkey {
	newMonkey := &Monkey{
		itemQueue: queue.New(),
		operation: operation,
		testFunc:  testFunc,
	}
	for _, item := range startingItems {
		newMonkey.itemQueue.Enqueue(item)
	}
	return newMonkey
}

// AssignTargets Once we've estashed the group of monkeys,
// we need to assign them to their respective targets
func (m *Monkey) AssignTargets(tm *Monkey, fm *Monkey) {
	m.trueMonkey = tm
	m.falseMonkey = fm
}

func (m *Monkey) Receive(item int64) {
	m.itemQueue.Enqueue(item)
}

// ProcessItems
// Each monkey eventually processes the items:
// Keep processing items off the queue until nothing.
// (part1) Run the operation, divide the value by 3
// (part 2) modulo the divisor to keep the numbers down,
// then based on the test function, send it to the respective monkey.
func (m *Monkey) ProcessItems() {
	nextItem := m.itemQueue.Dequeue()
	for nextItem != nil {
		m.inspectionCount++
		nextInt, ok := nextItem.(int64)
		if !ok {
			log.Fatal("Could not parse value from item queue")
		}
		newVal := m.operation(nextInt)
		// partone
		//newVal = int(float64(newVal) / 3)
		// parttwo
		newVal = newVal % DIVISOR
		if m.testFunc(newVal) {
			m.trueMonkey.Receive(newVal)
		} else {
			m.falseMonkey.Receive(newVal)
		}
		// this counts how many items pass through the monkey
		//m.inspectionCount++
		nextItem = m.itemQueue.Dequeue()
	}
}

func SetupMonkies() []*Monkey {
	// first initialize the monkies
	monkies := []*Monkey{
		NewMonkey(
			[]int64{74, 64, 74, 63, 53},
			func(old int64) int64 { return old * 7 },
			func(val int64) bool { return val%5 == 0 },
		),
		NewMonkey(
			[]int64{69, 99, 95, 62},
			func(old int64) int64 { return old * old },
			func(val int64) bool { return val%17 == 0 },
		),
		NewMonkey(
			[]int64{59, 81},
			func(old int64) int64 { return old + 8 },
			func(val int64) bool { return val%7 == 0 },
		),
		NewMonkey(
			[]int64{50, 67, 63, 57, 63, 83, 97},
			func(old int64) int64 { return old + 4 },
			func(val int64) bool { return val%13 == 0 },
		),
		NewMonkey(
			[]int64{61, 94, 85, 52, 81, 90, 94, 70},
			func(old int64) int64 { return old + 3 },
			func(val int64) bool { return val%19 == 0 },
		),
		NewMonkey(
			[]int64{69},
			func(old int64) int64 { return old + 5 },
			func(val int64) bool { return val%3 == 0 },
		),
		NewMonkey(
			[]int64{54, 55, 58},
			func(old int64) int64 { return old + 7 },
			func(val int64) bool { return val%11 == 0 },
		),
		NewMonkey(
			[]int64{79, 51, 83, 88, 93, 76},
			func(old int64) int64 { return old * 3 },
			func(val int64) bool { return val%2 == 0 },
		),
	}
	// now we need to assign their targets
	monkies[0].AssignTargets(monkies[1], monkies[6])
	monkies[1].AssignTargets(monkies[2], monkies[5])
	monkies[2].AssignTargets(monkies[4], monkies[3])
	monkies[3].AssignTargets(monkies[0], monkies[7])
	monkies[4].AssignTargets(monkies[7], monkies[3])
	monkies[5].AssignTargets(monkies[4], monkies[2])
	monkies[6].AssignTargets(monkies[1], monkies[5])
	monkies[7].AssignTargets(monkies[0], monkies[6])
	return monkies
}

func Simulate(monkies []*Monkey, rounds int) int64 {
	for i := 1; i <= rounds; i++ {
		if (i < 1000 && i%100 == 0) || i%1000 == 0 {
			fmt.Println("Round: ", i)
		}
		for _, monkey := range monkies {
			monkey.ProcessItems()
		}
	}
	first := monkies[0]
	second := monkies[1]
	for _, monkey := range monkies {
		if monkey.inspectionCount > first.inspectionCount {
			second = first
			first = monkey
		} else if monkey.inspectionCount > second.inspectionCount {
			second = monkey
		}
	}
	monkeyBusiness := first.inspectionCount * second.inspectionCount
	return monkeyBusiness
}
