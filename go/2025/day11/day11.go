package day11

import (
	"fmt"
	"strings"

	"github.com/mobanhawi/aoc/2025/util"
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
	nodeIdx  int
	pathMask uint64
}

// solvePt2 solves part 2 of the puzzle
// finds all paths from 'you' to 'out' that visit 'fft' or 'dac'
func solvePt2(lines []string) int {
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

	// Build node ID -> bit index mapping
	nodeToIdx := make(map[string]int)
	idxToNode := make([]string, 0, len(g.Nodes))
	idx := 0
	for nodeId := range g.Nodes {
		nodeToIdx[nodeId] = idx
		idxToNode = append(idxToNode, nodeId)
		idx++
	}

	// Special nodes we need to check
	fftIdx := nodeToIdx["fft"]
	dacIdx := nodeToIdx["dac"]
	fftMask := uint64(1) << fftIdx
	dacMask := uint64(1) << dacIdx

	visited := make(map[string]bool)
	cache := make(map[cacheKey]int)

	var dfs func(nodeId string, pathMask uint64) int
	dfs = func(nodeId string, pathMask uint64) int {
		nodeIdx := nodeToIdx[nodeId]
		key := cacheKey{nodeIdx: nodeIdx, pathMask: pathMask}

		if cached, ok := cache[key]; ok {
			fmt.Println("Using cached value for", key, ":", cached)
			return cached
		}

		nodeMask := uint64(1) << nodeIdx
		pathMask |= nodeMask

		if nodeId == "out" {
			result := 0
			if (pathMask&fftMask != 0) && (pathMask&dacMask != 0) {
				result = 1
			}
			cache[key] = result
			return result
		}

		visited[nodeId] = true
		totalPaths := 0
		for _, neighborId := range g.Nodes[nodeId].Edges {
			if !visited[neighborId] {
				totalPaths += dfs(neighborId, pathMask)
			}
		}
		visited[nodeId] = false

		cache[key] = totalPaths
		return totalPaths
	}

	return dfs("svr", 0)
}
