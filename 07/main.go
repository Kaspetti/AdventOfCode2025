package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)


func main() {
	task01()
	task02()
}


func task01() {
	inp, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inpSplit := strings.Split(strings.TrimSpace(string(inp)), "\n")
	startIndex := strings.Index(inpSplit[0], "S")

	beamX := map[int]bool{startIndex: true}
	totalSplit := 0

	for _, line := range inpSplit[1:] {
		for x := range beamX {
			if !beamX[x] {
				continue
			}
			if line[x] == '^' {
				beamX[x] = false
				beamX[x - 1] = true
				beamX[x + 1] = true
				totalSplit++
			}
		}
	}

	fmt.Println(totalSplit)
}


func task02() {
	inp, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inpSplit := strings.Split(strings.TrimSpace(string(inp)), "\n")
	startIndex := strings.Index(inpSplit[0], "S")

	activeBeams := make([]int, len(inpSplit[0]))
	activeBeams[startIndex] = 1

	for _, line := range inpSplit[1:] {
		for i, x := range activeBeams {
			if x == 0 {
				continue
			}
			if line[i] == '^' {
				activeBeams[i - 1] += x
				activeBeams[i + 1] += x
				activeBeams[i] = 0
			}
		}
	}

	total := 0
	for _, num := range activeBeams {
		total += num
	}

	fmt.Println(total)
}
