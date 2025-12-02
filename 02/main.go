package main

import (
	"fmt"
	"os"
	"github.com/dlclark/regexp2"

	"strings"

	"strconv"
)
func main() {
	task01()
	task02()
}

func task01() {
	inp, err := os.ReadFile("input")
	if err != nil {
		panic(err)	
	}

	idRanges := strings.Split(strings.TrimSpace(string(inp)), ",")
	var sumInvalidIds int64 = 0
	for _, idRange := range idRanges {
		splitRange := strings.Split(idRange, "-")

		lowerInt, err := strconv.ParseInt(splitRange[0], 10, 64)
		if err != nil {
			panic(err)	
		}

		upperInt, err := strconv.ParseInt(splitRange[1], 10, 64)
		if err != nil {
			panic(err)	
		}

		for id := lowerInt; id <= upperInt; id++ {
			stringId := strconv.FormatInt(id, 10)
			if stringId[:len(stringId)/2] == stringId[len(stringId)/2:] {
				sumInvalidIds += id
			}
		}
	}

	fmt.Println(sumInvalidIds)

}

func task02() {
	inp, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	idRanges := strings.Split(strings.TrimSpace(string(inp)), ",")
	var sumInvalidIds int64 = 0

	re := regexp2.MustCompile(`^(\w+)\1+$`, 0) // Slow, men fuck it. We ball (gidd ikkje sjekke individuelt :)) Regex <3)

	for _, idRange := range idRanges {
		splitRange := strings.Split(idRange, "-")
		lowerInt, err := strconv.ParseInt(splitRange[0], 10, 64)
		if err != nil {
			panic(err)	
		}

		upperInt, err := strconv.ParseInt(splitRange[1], 10, 64)
		if err != nil {
			panic(err)	
		}

		for id := lowerInt; id <= upperInt; id++ {
			stringId := strconv.FormatInt(id, 10)
			match, err := re.FindStringMatch(stringId)
			if err != nil {
				panic(err)
			}
			if match != nil && len(match.String()) == len(stringId) {
				sumInvalidIds += id				
			}
		}
	}

	fmt.Println(sumInvalidIds)
}
