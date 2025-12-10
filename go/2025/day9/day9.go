package day9

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/mobanhawi/aoc/2025/util"
)

/*
Solve https://adventofcode.com/2025/day/9
*/
func Solve() {
	fmt.Println("2025/day/9 pt1", solvePt1(util.ReadLines("./2025/day9/input.txt")))
	fmt.Println("2025/day/9 pt2", solvePt2(util.ReadLines("./2025/day9/input.txt")))
}

type metric struct {
	p1 int // point index
	p2 int // point index
	d  int //  metric
}

type metricHeap []*metric

func (h metricHeap) Len() int           { return len(h) }
func (h metricHeap) Less(i, j int) bool { return h[i].d > h[j].d } // max heap
func (h metricHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *metricHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*metric))
}
func (h *metricHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// dist diagonal squared distance between two points (proxy for area)
func dist(p1, p2 []int) int {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	return dx*dx + dy*dy
}

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string) int {
	tiles := make([][]int, len(lines))
	for i, line := range lines {
		tiles[i] = make([]int, 2)
		nums := strings.Split(line, ",")
		tiles[i][0], _ = strconv.Atoi(nums[0]) // x
		tiles[i][1], _ = strconv.Atoi(nums[1]) // y
	}
	// pre-compute all distances
	n := len(tiles)
	var h metricHeap
	heap.Init(&h)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := metric{
				p1: i, p2: j,
				d: dist(tiles[i], tiles[j]),
			}
			heap.Push(&h, &d)
		}
	}
	d := heap.Pop(&h).(*metric)
	return (util.Abs(tiles[d.p1][0]-tiles[d.p2][0]) + 1) * (util.Abs(tiles[d.p1][1]-tiles[d.p2][1]) + 1)
}
