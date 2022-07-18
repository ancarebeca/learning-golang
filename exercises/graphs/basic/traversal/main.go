package main

import (
	"fmt"
)

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

func (g *graph) addEdge(origin, destination int) *graph {
	g.AdjMatrix[origin][destination] = 1
	return g
}

func (g *graph) printMatrix() {
	for i, row := range g.AdjMatrix {
		fmt.Printf("%d: ", i)
		for j := range row {
			fmt.Printf("%d ", g.AdjMatrix[i][j])
		}
		fmt.Printf("\n")
	}

}

func (g *graph) bfs() {
	visited := make(map[int]bool, len(g.AdjMatrix[0]))
	totalNodes := len(g.AdjMatrix[0])
	var q []int

	for i := 0; i < totalNodes; i++ {
		if !visited[i] {
			q = append(q, i)
			fmt.Printf("%v", g.bfsRec(q, visited))
			fmt.Println()
		}
	}

}

func (g *graph) bfsRec(queue []int, visited map[int]bool) []int {
	result := []int{}
	if len(queue) == 0 {
		return []int{}
	}

	node := queue[0]
	queue = queue[1:]

	if !visited[node] {
		//fmt.Printf("%d ", node)
		result = append(result, node)
	}

	visited[node] = true
	totalNodes := len(g.AdjMatrix[0])

	for j := 0; j < totalNodes; j++ {
		if g.AdjMatrix[node][j] == 1 && !visited[j] {
			queue = append(queue, j)
		}
	}

	result = append(result, g.bfsRec(queue, visited)...)
	return result
}

func (g *graph) dfs() {
	visited := make(map[int]bool, len(g.AdjMatrix[0]))
	totalNodes := len(g.AdjMatrix[0])
	var stack []int

	for i := 0; i < totalNodes; i++ {
		if !visited[i] {
			stack = append(stack, i)
			fmt.Printf("%v", g.dfsRec(stack, visited))
			fmt.Println()
		}
	}

}

func (g *graph) dfsRec(stack []int, visited map[int]bool) []int {
	result := []int{}
	if len(stack) == 0 {
		return []int{}
	}

	node := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	if !visited[node] {
		//fmt.Printf("%d ", node)
		result = append(result, node)
	}

	visited[node] = true
	totalNodes := len(g.AdjMatrix[0])

	for j := 0; j < totalNodes; j++ {
		if g.AdjMatrix[node][j] == 1 && !visited[j] {
			stack = append(stack, j)
		}
	}

	result = append(result, g.dfsRec(stack, visited)...)
	return result
}

func main() {
	g := newGraph(7)
	g.addEdge(0, 1)
	g.addEdge(1, 4)
	g.addEdge(0, 2)
	g.addEdge(2, 3)

	fmt.Println("Adjacency Matrix")
	g.printMatrix()
	fmt.Println("breadth-first search (BFS) ")
	g.bfs()
	fmt.Println("depth-first search (DFS) ")
	g.dfs()
}
