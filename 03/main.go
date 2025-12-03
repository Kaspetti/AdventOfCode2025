package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	task01()	
	task02()
}

func task01() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "")

		joltageStr := ""
		largest := 0
		largestPos := 0
		for i, numStr := range lineSplit[:len(lineSplit) - 1] {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			if num > largest {
				largest = num
				largestPos = i
				joltageStr = numStr
			}
		}

		secondLargest := 0
		for _, numStr := range lineSplit[largestPos+1:] {	
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			if num > secondLargest {
				secondLargest = num
			}
		}

		joltageStr += strconv.Itoa(secondLargest)
		joltage, err := strconv.Atoi(joltageStr)
		if err != nil {
			panic(err)
		}

		total += joltage
	}

	fmt.Println(total)
}

func task02() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var total int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "")

		batteries := make([]int, 0)
		for _, numStr := range lineSplit {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			batteries = append(batteries, num)
		}

		joltageStr := ""
		batteryCandidates := batteries[:len(batteries) - 11]
		iCur := 0
		for len(joltageStr) < 12 {
			largestNum := 0
			largestPos := 0
			for i, num := range batteryCandidates {
				if num > largestNum {
					largestNum = num
					largestPos = i
				}
			}

			joltageStr += strconv.Itoa(largestNum)
			iCur += largestPos + 1
			batteryCandidates = batteries[iCur:len(batteries) - (11 - len(joltageStr))]
		}

		num, err := strconv.ParseInt(joltageStr, 10, 64)
		if err != nil {
			panic(err)
		}

		total += num
	}

	fmt.Println(total)
}
