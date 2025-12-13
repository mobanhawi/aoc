package day11

import (
	"fmt"
	"strings"

	"github.com/mobanhawi/aoc/util"
)

/*
Solve https://adventofcode.com/2025/day/10
*/
func Solve() {
	fmt.Println("2025/day/11 pt1", solvePt1(util.ReadLines("./2025/day11/input.txt")))
	fmt.Println("2025/day/11 pt2", solvePt2(util.ReadLines("./2025/day11/input.txt")))
}

// Node is a node in the graph; it has a (unique) ID and a sequence of
// edges to other nodes.
type Node struct {
	Id    string
	Edges []string
}

// Graph contains a set of Nodes, uniquely identified by numeric IDs.
type Graph struct {
	Nodes map[string]Node
}

// solvePt1 solves part 1 of the puzzle
// finds all paths from 'you' to 'out'
func solvePt1(lines []string) int {
	// build graph
	// map from -> to
	g := Graph{Nodes: make(map[string]Node)}
	for _, line := range lines {
		before, after, _ := strings.Cut(line, ":")
		from := strings.TrimSpace(before)
		to := strings.Fields(after)
		node, exists := g.Nodes[from]
		if !exists {
			node = Node{Id: from, Edges: []string{}}
		}
		node.Edges = append(node.Edges, to...)
		g.Nodes[from] = node
	}
	// find all paths from start 'you' to end 'out'
	visited := make(map[string]bool)
	var dfs func(nodeId string) int
	dfs = func(nodeId string) int {
		if nodeId == "out" {
			return 1
		}
		visited[nodeId] = true // avoid cycles
		totalPaths := 0
		for _, neighborId := range g.Nodes[nodeId].Edges {
			if !visited[neighborId] {
				totalPaths += dfs(neighborId)
			}
		}
		visited[nodeId] = false // backtrack
		return totalPaths
	}
	return dfs("you")
}

// cacheKey represents a DFS state for memoization
type cacheKey struct {
	nodeId           string
	importantVisited uint8 // bit 0 = fft, bit 1 = dac
}

// solvePt2 solves part 2 of the puzzle
// finds all paths from 'svr' to 'out' that visit 'fft' or 'dac'
func solvePt2(lines []string) int {
	// build directed graph
	g := Graph{Nodes: make(map[string]Node)}
	for _, line := range lines {
		before, after, _ := strings.Cut(line, ":")
		from := strings.TrimSpace(before)
		to := strings.Fields(after)
		node, exists := g.Nodes[from]
		if !exists {
			node = Node{Id: from, Edges: []string{}}
		}
		node.Edges = append(node.Edges, to...)
		g.Nodes[from] = node
	}

	// Bit masks for nodes we are tracking
	fftBit := uint8(1 << 0)
	dacBit := uint8(1 << 1)

	visited := make(map[string]bool)
	cache := make(map[cacheKey]int)

	var dfs func(nodeId string, importantVisited uint8) int
	dfs = func(nodeId string, importantVisited uint8) int {
		key := cacheKey{nodeId: nodeId, importantVisited: importantVisited}
		if cached, ok := cache[key]; ok {
			return cached
		}
		if nodeId == "fft" {
			importantVisited |= fftBit
		} else if nodeId == "dac" {
			importantVisited |= dacBit
		}

		if nodeId == "out" {
			result := 0
			// Only count paths that visited both fft AND dac
			if (importantVisited&fftBit != 0) && (importantVisited&dacBit != 0) {
				result = 1
			}
			cache[key] = result
			return result
		}

		visited[nodeId] = true
		totalPaths := 0
		for _, neighborId := range g.Nodes[nodeId].Edges {
			if !visited[neighborId] {
				totalPaths += dfs(neighborId, importantVisited)
			}
		}
		visited[nodeId] = false

		cache[key] = totalPaths
		return totalPaths
	}

	return dfs("svr", 0)
}
