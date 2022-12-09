package parttwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SectionRange struct {
	min int
	max int
}

func NewSectionRange(val string) *SectionRange {
	rangeVals := strings.Split(val, "-")
	if len(rangeVals) != 2 {
		log.Fatal("Range does not contain two elements")
	}
	valA, err := strconv.ParseInt(rangeVals[0], 10, 32)
	valB, err := strconv.ParseInt(rangeVals[1], 10, 32)
	if err != nil {
		log.Fatal("failed to parse range val")
	}
	if valA > valB {
		log.Fatal("Invalid range value")
	}
	return &SectionRange{int(valA), int(valB)}
}

type SectionRangePair struct {
	pairA *SectionRange
	pairB *SectionRange
}

func (p *SectionRangePair) Overlaps() bool {
	// 6-7, 3-8 => 6 is greater than 3 and 6 is less than 8
	pairAMinOverlapsB := p.pairA.min >= p.pairB.min && p.pairA.min <= p.pairB.max
	// 6-7, 6-8 =>  7  is less than 8, and 7 is greater than 6
	pairAMaxOverlapsB := p.pairA.max <= p.pairB.max && p.pairA.max >= p.pairB.min
	// 3-8, 6-7 => 6 is greater than 3, and 6 is less than 8
	pairBMinOverlapsA := p.pairB.min >= p.pairA.min && p.pairB.min <= p.pairA.max
	// 3-8, 6-7 =>  7 is less than 8 and 7 is greater than 3
	pairBMaxOverlapsA := p.pairB.max <= p.pairA.max && p.pairB.max >= p.pairA.min
	return pairAMinOverlapsB || pairAMaxOverlapsB || pairBMinOverlapsA || pairBMaxOverlapsA
}

func NewSectionRangePair(val string) *SectionRangePair {
	// Split by comma
	pairs := strings.Split(val, ",")
	if len(pairs) != 2 {
		log.Fatal("Pair not split into two ranges")
	}
	return &SectionRangePair{NewSectionRange(pairs[0]), NewSectionRange(pairs[1])}
}

func ReadInputIntoSectionRangePairs(filepath string) []*SectionRangePair {
	fmt.Println(fmt.Sprintf("Reading file into section range pairs: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	var sectionRangePairs []*SectionRangePair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowText := scanner.Text()
		sectionRangePair := NewSectionRangePair(rowText)
		sectionRangePairs = append(sectionRangePairs, sectionRangePair)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		panic(err)
	}

	return sectionRangePairs
}
