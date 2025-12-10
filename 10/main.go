package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/mxschmitt/golang-combinations"
)

type Machine struct {
	lights []int
	buttons [][]int
	joltageReq []int
}

func main() {
	task01()
	task02()
}

func task01() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		machine := parseMachine(line)

		maxCombLen := slices.Max(machine.joltageReq)
	}

	fmt.Println(total)
}

func task02() {
	f, err := os.Open("testinput")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		machine := parseMachine(line)

		maxJoltage := slices.Max(machine.joltageReq)
		combs := AllRepeat(machine.buttons, maxJoltage, 15)
		fmt.Println(len(combs))
		sort.Slice(combs, func(i, j int) bool {
			return len(combs[i]) < len(combs[j])
		})


		for _, comb := range combs {
			curJoltage := make([]int, len(machine.joltageReq))

			for _, button := range comb {
				for _, swtch := range button {
					curJoltage[swtch]++
				}
			}

			isEqual := true
			for i, joltage := range curJoltage {
				if joltage != machine.joltageReq[i] {
					isEqual = false
					break
				}
			}

			if isEqual {
				total += len(comb)
				break
			}
		}
	}

	fmt.Println(total)
}

func parseMachine(line string) Machine {
	lineSplit := strings.Split(line, " ")
	lightsStr := lineSplit[0]
	lights := make([]int, 0)
	for _, r := range lightsStr[1:len(lightsStr)-1] {
		switch r {
		case '.':
			lights = append(lights, 0)
		case '#':
			lights = append(lights, 1)
		}	
	}

	buttonsSlice := lineSplit[1:len(lineSplit)-1]
	buttons := make([][]int, 0)
	for _, buttonStr := range buttonsSlice {
		buttonSplit := strings.Split(buttonStr[1:len(buttonStr)-1], ",")
		switches := make([]int, 0)	
		for _, switchStr := range buttonSplit {
			num, err := strconv.Atoi(switchStr)
			if err != nil {
				log.Fatal(err)
			}

			switches = append(switches, num)
		}

		buttons = append(buttons, switches)
	}

	joltageStr := lineSplit[len(lineSplit)-1]
	joltageSlice := strings.Split(joltageStr[1:len(joltageStr)-1], ",")
	joltages := make([]int, 0)
	for _, joltage := range joltageSlice {
		num, err := strconv.Atoi(joltage)
		if err != nil {
			log.Fatal(err)
		}

		joltages = append(joltages, num)
	}

	return Machine{lights, buttons, joltages}
}

func AllRepeat[T any](set []T, mn int, mx int) (subsets [][]T) {
	if mx < 1 {
		return nil
	}

	var generateCombos func([]T, int)
	generateCombos = func(current []T, depth int) {
		if depth == 0 {
			subset := make([]T, len(current))
			copy(subset, current)
			subsets = append(subsets, subset)
			return
		}

		for _, item := range set {
			generateCombos(append(current, item), depth-1)
		}
	}

	for length := mn; length <= mx; length++ {
		generateCombos([]T{}, length)
	}

	return subsets
}
