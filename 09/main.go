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
}

type Dist struct {
	c1 		Coord
	c2 		Coord
	dist 	float64
}

type Area struct {
	c1		Coord
	c2		Coord
	area 	int
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
	maxArea := 0
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			maxArea = max(maxArea, calculateArea(c1, c2))
		}
	}

	fmt.Println(maxArea)
}


func task02() {
	inp, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(inp)), "\n")
	coords := parseCoords(lines)
	areas := make([]Area, 0)
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			areas = append(areas, Area{c1, c2, calculateArea(c1, c2)})
		}
	}
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].area > areas[j].area	
	})

	insidePoints := make(map[Coord]bool)

	for _, area := range areas {
		// Check if each inner point is within the concave hull made by the points
		minX := min(area.c1.x, area.c2.x)
		maxX := max(area.c1.x, area.c2.x)
		minY := min(area.c1.y, area.c2.y)
		maxY := max(area.c1.y, area.c2.y)

		leftEdgeValid := true
		for y := minY; y <= maxY; y++ {
			p := Coord{minX, y}
			inside, ok := insidePoints[p]
			if !ok {
				inside = isInsideHull(p, coords)
				insidePoints[p] = inside
			}

			if !inside {
				leftEdgeValid = false
				break
			}
		}
		if !leftEdgeValid { continue }
		rightEdgeValid := true
		for y := minY; y <= maxY; y++ {
			p := Coord{maxX, y}
			inside, ok := insidePoints[p]
			if !ok {
				inside = isInsideHull(p, coords)
				insidePoints[p] = inside
			}

			if !inside {
				rightEdgeValid = false
				break
			}
		}
		if !rightEdgeValid { continue }
		bottomEdgeValid := true
		for x := minX; x <= maxX; x++ {
			p := Coord{x, minY}
			inside, ok := insidePoints[p]
			if !ok {
				inside = isInsideHull(p, coords)
				insidePoints[p] = inside
			}

			if !inside {
				bottomEdgeValid = false
				break
			}
		}
		if !bottomEdgeValid { continue }
		topEdgeValid := true
		for x := minX; x <= maxX; x++ {
			p := Coord{x, maxY}
			inside, ok := insidePoints[p]
			if !ok {
				inside = isInsideHull(p, coords)
				insidePoints[p] = inside
			}

			if !inside {
				topEdgeValid = false
				break
			}
		}
		if !topEdgeValid { continue }

		fmt.Println(area.area)
		break
	}

}


func parseCoords(lines []string) []Coord {
	coords := make([]Coord, 0)
	for _, coordStr := range lines {
		comps := strings.Split(coordStr, ",")
		x, err := strconv.Atoi(comps[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(comps[1])
		if err != nil {
			log.Fatal(err)
		}

		coords = append(coords, Coord{x, y})
	}

	return coords
}


func calculateArea(c1, c2 Coord) int {
	xLen := int(math.Abs(float64(c2.x - c1.x)) + 1)
	yLen := int(math.Abs(float64(c2.y - c1.y)) + 1)

	return xLen * yLen
}


func isInsideHull(p Coord, hull []Coord) bool {
    inside := false
    j := len(hull) - 1

    for i := range hull {
        xi, yi := hull[i].x, hull[i].y
        xj, yj := hull[j].x, hull[j].y

        if isOnSegment(p, hull[j], hull[i]) {
            return true
        }

        if ((yi > p.y) != (yj > p.y)) &&
            (p.x < (xj-xi)*(p.y-yi)/(yj-yi)+xi) {
            inside = !inside
        }
        j = i
    }

    return inside
}

func isOnSegment(p, a, b Coord) bool {
    cross := (p.x-a.x)*(b.y-a.y) - (p.y-a.y)*(b.x-a.x)
    if cross != 0 {
        return false
    }

    return p.x >= min(a.x, b.x) && p.x <= max(a.x, b.x) &&
           p.y >= min(a.y, b.y) && p.y <= max(a.y, b.y)
}
