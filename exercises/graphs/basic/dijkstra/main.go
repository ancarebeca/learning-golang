package main

import (
	"fmt"
	"math"
)

// Dijkstra's algorithm allows us to find the shortest path between any two vertices of a graph.
// Time Complexity: O(E Log V) where, E is the number of edges and V is the number of vertices.
// Space Complexity: O(V)

func (g *graph) Dijkstra(S int) []int {
	totalNodes := len(g.AdjMatrix)
	distance := make([]int, totalNodes)
	visitedVertex := make([]bool, totalNodes)
	prevVertex := make([]int, totalNodes)

	for node := 0; node < totalNodes; node++ {
		distance[node] = math.MaxInt
	}

	// Distance of self loop is zero
	distance[S] = 0

	for i := 0; i < totalNodes; i++ {
		// Update the distance between neighbouring vertex and source vertex
		u := findMinDistance(distance, visitedVertex)
		visitedVertex[u] = true

		// Update all the neighbouring vertex distances
		for v := 0; v < totalNodes; v++ {
			// for each unvisited node v adjacent to u, let's calculate the distance v, u and update it if it is minimal
			if !visitedVertex[v] && g.AdjMatrix[u][v] != 0 {
				if distance[u]+g.AdjMatrix[u][v] < distance[v] {
					distance[v] = distance[u] + g.AdjMatrix[u][v]
					prevVertex[v] = u
				}
			}
		}
	}
	fmt.Println(prevVertex)

	return distance
}

// Finding the minimum distance
func findMinDistance(distance []int, visitedVertex []bool) int {
	minDistance := math.MaxInt
	minDistanceVertex := -1

	for i := 0; i < len(distance); i++ {
		if !visitedVertex[i] && distance[i] < minDistance {
			minDistance = distance[i]
			minDistanceVertex = i
		}
	}
	return minDistanceVertex
}

type graph struct {
	AdjMatrix [][]int
}

func newGraph(totalNodes int) graph {
	matrix := make([][]int, totalNodes)
	for i := range matrix {
		matrix[i] = make([]int, totalNodes)
	}
	return graph{
		AdjMatrix: matrix,
	}
}

func main() {
	// Providing the graph
	g := newGraph(5)
	g.AdjMatrix = [][]int{
		{0, 1, 6, 0, 0},
		{1, 0, 2, 1, 0},
		{6, 2, 0, 2, 5},
		{0, 1, 2, 0, 5},
		{0, 0, 5, 5, 0},
	}

	g.Dijkstra(0)

}
