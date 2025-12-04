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
	adjacents := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	total := 0

	for y, row := range inpSplit {
		for x, r := range row {
			if r != '@' {
				continue
			}

			count := 0
			for _, adjacent := range adjacents {
				newX := x + adjacent[0]
				newY := y + adjacent[1]

				if newX < 0 || newX >= len(row) {
					continue
				}
				if newY < 0 || newY >= len(inpSplit) {
					continue
				}

				if inpSplit[newY][newX] == '@' {
					count++
				}
			}

			if count < 4 {
				total++
			}
		}
	}

	fmt.Println(total)
}

func task02() {
	inp, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inpSplit := strings.Split(strings.TrimSpace(string(inp)), "\n")
	runeGrid := make([][]rune, 0)
	for _, row := range inpSplit {
		runeGrid = append(runeGrid, []rune(row))
	}
	adjacents := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	total := 0

	for true {
		totalRun := 0
		for y, row := range runeGrid {
			for x, r := range row {
				if r != '@' {
					continue
				}

				count := 0
				for _, adjacent := range adjacents {
					newX := x + adjacent[0]
					newY := y + adjacent[1]

					if newX < 0 || newX >= len(row) {
						continue
					}
					if newY < 0 || newY >= len(inpSplit) {
						continue
					}

					if runeGrid[newY][newX] == '@' {
						count++
					}
				}

				if count < 4 {
					totalRun++
					runeGrid[y][x] = '.'
				}
			}
		}

		if totalRun == 0 {
			break
		}

		total += totalRun
	}

	fmt.Println(total)
}
