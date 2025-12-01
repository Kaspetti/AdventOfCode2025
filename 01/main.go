package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	pointer := 50
	zeros := 0

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == 'L' {
			pointer -= steps
		} else {
			pointer += steps
		}

		pointer = (pointer % 100 + 100) % 100

		if pointer == 0 {
			zeros++
		}
	}

	fmt.Println(zeros)
}

func task02() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)	
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pointer := 50
	zeros := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		
		if direction == 'L' {
			if pointer - steps <= 0 {
				zeros += (steps - pointer) / 100

				if pointer != 0 {
					zeros++
				}
			}

			pointer -= steps
		} else {
			zeros += (pointer + steps) / 100
			pointer += steps
		}

		pointer = (pointer % 100 + 100) % 100
	}

	fmt.Println(zeros)
}
