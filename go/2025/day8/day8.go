package day8

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mobanhawi/aoc/2025/util"
)

type pos struct {
	x       int
	y       int
	z       int
	circuit int // index of circuit it belongs to
}

type circuit struct {
	points []*pos
	size   int
}

type distance struct {
	p1 *pos // point 1
	p2 *pos // point 2
	d  int  //  distance metric
}

type distanceHeap []distance

func (h distanceHeap) Len() int           { return len(h) }
func (h distanceHeap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h distanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *distanceHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(distance))
}
func (h *distanceHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// dist squared distance between two points (proxy for dist)
func (p *pos) dist(other *pos) distance {
	dx := p.x - other.x
	dy := p.y - other.y
	dz := p.z - other.z
	return distance{
		p1: p, p2: other, d: dx*dx + dy*dy + dz*dz,
	}
}

/*
Solve https://adventofcode.com/2025/day/6
*/
func Solve() {
	fmt.Println("2025/day/8 pt1", solvePt1(util.ReadLines("./2025/day8/input.txt"), 1000))
	fmt.Println("2025/day/8 pt2", solvePt2(util.ReadLines("./2025/day8/input.txt")))
}

// solvePt1 solves part 1 of the puzzle
// finds number of junctions in the 3 largest circuits formed by connecting points
func solvePt1(lines []string, conn int) int {
	points := make([]*pos, len(lines))
	// parse input O(n)
	for i, line := range lines {
		p := strings.Split(line, ",")
		points[i] = &pos{}
		points[i].x, _ = strconv.Atoi(p[0])
		points[i].y, _ = strconv.Atoi(p[1])
		points[i].z, _ = strconv.Atoi(p[2])
	}
	// pre-compute all distances
	n := len(points)
	var h distanceHeap
	heap.Init(&h)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := points[i].dist(points[j])
			heap.Push(&h, dist)
		}
	}

	// make n shortest connections
	// current circuit
	cir := 1 // zero means no circuit
	circuits := make(map[int]*circuit, 0)
	for range conn {
		d := heap.Pop(&h).(distance)
		// fmt.Println("considering distance", d.d, "between", *d.p1, "and", *d.p2)
		if d.p1.circuit == 0 && d.p2.circuit == 0 {
			// new circuit
			cir++
			circuits[cir] = &circuit{
				points: []*pos{d.p1, d.p2},
				size:   2,
			}
			d.p1.circuit = cir
			d.p2.circuit = cir
			// fmt.Println("created new circuit", cir, "with points", *d.p1, "and", *d.p2)
		} else if d.p1.circuit != 0 && d.p2.circuit == 0 {
			// add p2 to p1's circuit
			circuitIdx := d.p1.circuit
			circuits[circuitIdx].points = append(circuits[circuitIdx].points, d.p2)
			circuits[circuitIdx].size++
			d.p2.circuit = circuitIdx
			// fmt.Println("add point", *d.p2, "to circuit", circuitIdx)
		} else if d.p1.circuit == 0 && d.p2.circuit != 0 {
			// add p1 to p2's circuit
			circuitIdx := d.p2.circuit
			circuits[circuitIdx].points = append(circuits[circuitIdx].points, d.p1)
			circuits[circuitIdx].size++
			d.p1.circuit = circuitIdx
			// fmt.Println("add point", *d.p2, "to circuit", circuitIdx)
		} else if d.p1.circuit != 0 && d.p2.circuit != 0 && d.p1.circuit != d.p2.circuit {
			// merge circuits
			circuitIdx1 := d.p1.circuit
			circuitIdx2 := d.p2.circuit
			// fmt.Println("merged circuits", circuitIdx1, "and", circuitIdx2)
			// move all points from circuit2 to circuit1
			for _, p := range circuits[circuitIdx2].points {
				p.circuit = circuitIdx1
				circuits[circuitIdx1].points = append(circuits[circuitIdx1].points, p)
				circuits[circuitIdx1].size++
			}
			// delete circuit2
			delete(circuits, circuitIdx2)
		}
	}
	// sort circuits by size
	sortedCircuits := make([]*circuit, 0, len(circuits))
	for _, c := range circuits {
		sortedCircuits = append(sortedCircuits, c)
	}
	sort.Slice(sortedCircuits, func(i, j int) bool {
		return sortedCircuits[i].size > sortedCircuits[j].size
	})

	result := 1
	for i := range 3 { // always get top 3
		result *= sortedCircuits[i].size
	}

	return result
}

// solvePt2 solves part 1 of the puzzle
// finds the distance of the last connection that completes the circuit
// then multiplies the X coordinates of those two junction boxes
func solvePt2(lines []string) int {
	points := make([]*pos, len(lines))
	// parse input O(n)
	for i, line := range lines {
		p := strings.Split(line, ",")
		points[i] = &pos{}
		points[i].x, _ = strconv.Atoi(p[0])
		points[i].y, _ = strconv.Atoi(p[1])
		points[i].z, _ = strconv.Atoi(p[2])
	}
	// pre-compute all distances
	n := len(points)
	var h distanceHeap
	heap.Init(&h)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := points[i].dist(points[j])
			heap.Push(&h, dist)
		}
	}

	// make n shortest connections
	// current circuit
	cir := 1 // zero means no circuit
	circuits := make(map[int]*circuit, 0)
	var terminalConn distance
	for {
		d := heap.Pop(&h).(distance)
		// fmt.Println("considering distance", d.d, "between", *d.p1, "and", *d.p2)
		if d.p1.circuit == 0 && d.p2.circuit == 0 {
			// new circuit
			cir++
			circuits[cir] = &circuit{
				points: []*pos{d.p1, d.p2},
				size:   2,
			}
			d.p1.circuit = cir
			d.p2.circuit = cir
			// fmt.Println("created new circuit", cir, "with points", *d.p1, "and", *d.p2)
		} else if d.p1.circuit != 0 && d.p2.circuit == 0 {
			// add p2 to p1's circuit
			circuitIdx := d.p1.circuit
			circuits[circuitIdx].points = append(circuits[circuitIdx].points, d.p2)
			circuits[circuitIdx].size++
			d.p2.circuit = circuitIdx
			if circuits[circuitIdx].size == n {
				terminalConn = d
				break
			}
			// fmt.Println("add point", *d.p2, "to circuit", circuitIdx)
		} else if d.p1.circuit == 0 && d.p2.circuit != 0 {
			// add p1 to p2's circuit
			circuitIdx := d.p2.circuit
			circuits[circuitIdx].points = append(circuits[circuitIdx].points, d.p1)
			circuits[circuitIdx].size++
			d.p1.circuit = circuitIdx
			// fmt.Println("add point", *d.p2, "to circuit", circuitIdx)
			if circuits[circuitIdx].size == n {
				terminalConn = d
				break
			}
		} else if d.p1.circuit != 0 && d.p2.circuit != 0 && d.p1.circuit != d.p2.circuit {
			// merge circuits
			circuitIdx1 := d.p1.circuit
			circuitIdx2 := d.p2.circuit
			// fmt.Println("merged circuits", circuitIdx1, "and", circuitIdx2)
			// move all points from circuit2 to circuit1
			for _, p := range circuits[circuitIdx2].points {
				p.circuit = circuitIdx1
				circuits[circuitIdx1].points = append(circuits[circuitIdx1].points, p)
				circuits[circuitIdx1].size++
			}
			if circuits[circuitIdx1].size == n {
				terminalConn = d
				break
			}
			// delete circuit2
			delete(circuits, circuitIdx2)
		}
	}
	// multiplying the X coordinates of those two junction boxes
	return terminalConn.p1.x * terminalConn.p2.x
}
