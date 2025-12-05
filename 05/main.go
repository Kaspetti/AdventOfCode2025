package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	min int64
	max	int64
}

func main() {
	task01()
	task02()
}


func task01() {
	inpBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inpSplit := strings.Split(strings.TrimSpace(string(inpBytes)), "\n\n")
	ranges := createRanges(inpSplit[0])
	numFresh := 0

	for line := range strings.SplitSeq(inpSplit[1], "\n") {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		for _, r := range ranges {
			if num >= r.min && num <= r.max {
				numFresh++
				break
			}
		}
	}

	fmt.Println(numFresh)
}

func task02() {
	inpBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inpSplit := strings.Split(strings.TrimSpace(string(inpBytes)), "\n\n")
	ranges := createRanges(inpSplit[0])
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	var totalFresh int64 = 0
	var maxNum int64 = 0

	for _, r := range ranges {
		if r.max <= maxNum {
			continue
		}

		rangeStart := max(r.min, maxNum + 1)
		totalFresh += r.max - rangeStart + 1
		maxNum = r.max
	}

	fmt.Println(totalFresh)
}

func createRanges(input string) []Range {
	ranges := make([]Range, 0)
	for r := range strings.SplitSeq(input, "\n") {
		minMaxStr := strings.Split(r, "-")

		minRange, err := strconv.ParseInt(minMaxStr[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		maxRange, err := strconv.ParseInt(minMaxStr[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, Range{minRange, maxRange})
	}

	return ranges
}
