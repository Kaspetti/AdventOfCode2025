package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)


type Coord struct {
	x int
	y int
	z int
}

type Dist struct {
	c1 		Coord
	c2 		Coord
	dist 	float64
}


func main() {
	task01()
	task02()
}


func task01() {
	inp, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(inp)), "\n")
	coords := parseCoords(lines)
	
	dists := getDistances(coords)
	
	pathIDs := make(map[Coord]int)
	for _, coord := range coords {
		pathIDs[coord] = -1
	}
	i := 0

	for _, dist := range dists[:1000] {
		c1ID := pathIDs[dist.c1]
		c2ID := pathIDs[dist.c2]

		if c1ID == -1 && c2ID == -1 {
			pathIDs[dist.c1] = i
			pathIDs[dist.c2] = i
			i++
		} else if c1ID != -1 && c2ID != -1 {
			for c, id := range pathIDs {
				if id == c2ID {
					pathIDs[c] = c1ID
				}
			}
		} else if c1ID != -1 {
			pathIDs[dist.c2] = c1ID
		} else if c2ID != -1 {
			pathIDs[dist.c1] = c2ID
		}
	}

	pathSizes := make([]int, i)
	for _, id := range pathIDs {
		if id == -1 {
			continue
		}
		pathSizes[id]++
	}

	sort.Slice(pathSizes, func(i, j int) bool {
		return pathSizes[i] > pathSizes[j]		
	})

	fmt.Println(pathSizes[0] * pathSizes[1] * pathSizes[2])
}


func task02() {
	inp, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(inp)), "\n")
	coords := parseCoords(lines)
	
	dists := getDistances(coords)
	
	pathIDs := make(map[Coord]int)
	for _, coord := range coords {
		pathIDs[coord] = -1
	}
	circuits := len(coords)
	i := 0

	for _, dist := range dists {
		c1ID := pathIDs[dist.c1]
		c2ID := pathIDs[dist.c2]

		if c1ID == -1 && c2ID == -1 {
			pathIDs[dist.c1] = i
			pathIDs[dist.c2] = i
			i++
			circuits--
		} else if c1ID != -1 && c2ID != -1 {
			if c1ID == c2ID {
				continue
			}
			for c, id := range pathIDs {
				if id == c2ID {
					pathIDs[c] = c1ID
				}
			}
			circuits--
		} else if c1ID != -1 {
			pathIDs[dist.c2] = c1ID
			circuits--
		} else if c2ID != -1 {
			pathIDs[dist.c1] = c2ID
			circuits--
		}

		if circuits == 1 {
			fmt.Println(dist.c1.x * dist.c2.x)
			break
		}
	}
}


func parseCoords(lines []string) []Coord {
	coords := make([]Coord, 0)

	for _, line := range lines {
		comps := strings.Split(line, ",")
		x, err := strconv.Atoi(comps[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(comps[1])
		if err != nil {
			log.Fatal(err)
		}
		z, err := strconv.Atoi(comps[2])
		if err != nil {
			log.Fatal(err)
		}

		coords = append(coords, Coord{x, y, z})
	}

	return coords
}


func getDistances(coords []Coord) []Dist {
	dists := make([]Dist, 0)
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			dists = append(dists, Dist{c1, c2, distance(c1, c2)})
		}
	}

	sort.Slice(dists, func(i, j int) bool {
		return dists[i].dist < dists[j].dist
	})

	return dists
}


func distance(c1, c2 Coord) float64 {
	return math.Sqrt(math.Pow(float64(c2.x - c1.x), 2) + math.Pow(float64(c2.y - c1.y), 2) + math.Pow(float64(c2.z - c1.z), 2))
}
