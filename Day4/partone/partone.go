package partone

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
	pairAOverlapPairB := p.pairA.min <= p.pairB.min && p.pairA.max >= p.pairB.max
	pairBOverlapPairA := p.pairB.min <= p.pairA.min && p.pairB.max >= p.pairA.max
	return pairAOverlapPairB || pairBOverlapPairA
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
