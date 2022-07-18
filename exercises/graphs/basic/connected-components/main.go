package main

import "fmt"

func getConnectedComponents(adjacent map[int32][]int32) [][]int32 {
	var connectedComponents [][]int32
	visited := make(map[int32]bool)

	for node, _ := range adjacent {

		if !visited[node] {
			connectedComponents = append(connectedComponents, dfs(node, visited, adjacent))
		}
	}
	return connectedComponents
}

func dfs(node int32, visited map[int32]bool, adjacentList map[int32][]int32) []int32 {
	visited[node] = true
	result := []int32{node}

	for _, v := range adjacentList[node] {
		if !visited[v] {
			result = append(result, dfs(v, visited, adjacentList)...)
		}
	}

	return result
}

func createAdjacentList(nodes [][]int32) map[int32][]int32 {
	adjacent := make(map[int32][]int32)
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		adjacent[node[0]] = append(adjacent[node[0]], node[1])
		adjacent[node[1]] = append(adjacent[node[1]], node[0])
	}
	return adjacent
}

func main() {
	nodes := [][]int32{{1, 2}, {3, 1}, {2, 3}, {4, 5}}

	adjcent := createAdjacentList(nodes)
	fmt.Println("Adjacency List")
	fmt.Println(adjcent)

	fmt.Println("Connected Components ")
	fmt.Println(getConnectedComponents(adjcent))
}
