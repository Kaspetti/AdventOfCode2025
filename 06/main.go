package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	inpSplit := strings.Split(strings.TrimSuffix(string(inp), "\n"), "\n")
	operatorLine := inpSplit[len(inpSplit)-1]

	total := 0
	for i := 0; i < len(operatorLine); {
		operator := operatorLine[i]
		colLength := 1
		for j := i + 1; j < len(operatorLine) && operatorLine[j] == ' '; j++ {
			colLength++
		}

		var colTotal int
		switch operator {
			case '+':
				colTotal = 0
			case '*':
				colTotal = 1
		}
		for _, line := range inpSplit[:len(inpSplit) - 1] {
			numStr := line[i:i+colLength]
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				log.Fatal(err)
			}

			switch operator {
				case '+':
					colTotal += num
				case '*':
					colTotal *= num
			}
		}

		total += colTotal

		i += colLength
	}

	fmt.Println(total)
}


func task02() {
	inp, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inpSplit := strings.Split(strings.TrimSuffix(string(inp), "\n"), "\n")
	operatorLine := inpSplit[len(inpSplit)-1]

	total := 0
	for i := 0; i < len(operatorLine); {
		operator := operatorLine[i]
		colLength := 0
		for j := i + 1; j < len(operatorLine) && operatorLine[j] == ' '; j++ {
			colLength++
			if j == len(operatorLine) - 1 {
				colLength++
			}
		}

		var colTotal int
		switch operator {
			case '+':
				colTotal = 0
			case '*':
				colTotal = 1
		}

		for j := 0; j < colLength; j++ {
			numStr := ""
			for _, line := range inpSplit[:len(inpSplit) - 1] {
				char := line[i + j]
				if char != ' ' {
					numStr += string(line[i + j])
				}
			}

			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}

			switch operator {
				case '+':
					colTotal += num
				case '*':
					colTotal *= num
			}
		}

		total += colTotal

		i += colLength + 1
	}

	fmt.Println(total)
}
