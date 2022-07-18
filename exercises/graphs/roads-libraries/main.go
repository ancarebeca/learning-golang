package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getConnectedComponents(adjacent map[int32][]int32) [][]int32 {
	var connectedComponents [][]int32
	visited := make(map[int32]bool)

	for city, _ := range adjacent {

		if !visited[city] {
			connectedComponents = append(connectedComponents, dfs(city, visited, adjacent))
		}
	}
	return connectedComponents
}

func dfs(city int32, visited map[int32]bool, adjacentList map[int32][]int32) []int32 {
	visited[city] = true
	result := []int32{city}

	for _, v := range adjacentList[city] {
		if !visited[v] {
			result = append(result, dfs(v, visited, adjacentList)...)
		}
	}

	return result
}

func createAdjacentList(cities [][]int32) map[int32][]int32 {
	adjacent := make(map[int32][]int32)
	for i := 0; i < len(cities); i++ {
		city := cities[i]
		adjacent[city[0]] = append(adjacent[city[0]], city[1])
		adjacent[city[1]] = append(adjacent[city[1]], city[0])
	}
	return adjacent
}

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	// In case of building a library is cheaper than fixing a road
	if c_lib <= c_road {
		return int64(n) * int64(c_lib)
	}

	adjacent := createAdjacentList(cities)
	var ret int64

	sizeOfAdj := int32(len(adjacent))
	if n > sizeOfAdj {
		// It means there are "n-sizeOfAdj" cities that are not connect with other cities
		// in that case the only solution is to build a library
		ret = int64(n-sizeOfAdj) * int64(c_lib)
	}

	connectedCities := getConnectedComponents(adjacent)
	for i := 0; i < len(connectedCities); i++ {
		numRoads := len(connectedCities[i]) - 1                 // The number of edges to connect N nodes is N-1. In our case the edges are the roads
		ret += int64(c_lib) + (int64(numRoads) * int64(c_road)) // We need at least 1 library and N number of roads
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
