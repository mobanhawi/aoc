package day9

import (
	"strconv"
	"strings"

	"github.com/mobanhawi/aoc/2025/util"
)

type Edge struct {
	x1, y1, x2, y2 int
}

func intersections(minX, minY, maxX, maxY int, edges []*Edge) bool {
	for _, e := range edges {
		ex1, ex2 := util.Sort(e.x1, e.x2)
		ey1, ey2 := util.Sort(e.y1, e.y2)
		if minX < ex2 && maxX > ex1 && minY < ey2 && maxY > ey1 {
			return true
		}
	}
	return false
}

func area(x1, y1, x2, y2 int) int {
	width := util.Abs(x2-x1) + 1
	height := util.Abs(y2-y1) + 1
	return width * height
}

func solvePt2(lines []string) (result int) {
	result = 0
	edges := make([]*Edge, 0)
	tiles := make([][]int, len(lines))
	for i, line := range lines {
		tiles[i] = make([]int, 2)
		nums := strings.Split(line, ",")
		tiles[i][0], _ = strconv.Atoi(nums[0]) // x
		tiles[i][1], _ = strconv.Atoi(nums[1]) // y
		if i == 0 {
			continue // create new edge at the end
		}
		// connect all edges
		edges = append(edges, &Edge{tiles[i-1][0], tiles[i-1][1],
			tiles[i][0], tiles[i][1]})
	}
	// pre-compute all distances
	n := len(tiles)
	// close the polygon
	edges = append(edges, &Edge{tiles[0][0], tiles[0][1],
		tiles[n-1][0], tiles[n-1][1]})

	// from
	for f := 0; f < len(tiles)-1; f++ {
		// to
		for t := f; t < len(tiles); t++ {
			fromTile := tiles[f]
			toTile := tiles[t]
			minX, maxX := util.Sort(fromTile[0], toTile[0])
			minY, maxY := util.Sort(fromTile[1], toTile[1])
			if dist(fromTile, toTile) > result {
				if !intersections(minX, minY, maxX, maxY, edges) {
					area := area(fromTile[0], fromTile[1], toTile[0], toTile[1])
					if area > result {
						result = area
					}
				}
			}
		}
	}
	return result
}
