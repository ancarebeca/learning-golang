package main

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_roadsAndLibraries(t *testing.T) {
	citiesCaseOne := [][]int32{{1, 2}, {3, 1}, {2, 3}}
	assert.Equal(t, int64(4), roadsAndLibraries(3, 2, 1, citiesCaseOne))

	citiesCaseTwo := [][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}}
	assert.Equal(t, int64(12), roadsAndLibraries(6, 2, 5, citiesCaseTwo))

	citiesCaseThree := [][]int32{{1, 2}, {1, 3}, {1, 4}}
	assert.Equal(t, int64(15), roadsAndLibraries(5, 6, 1, citiesCaseThree))

	inputs, err := loadFile("fixture/test-cases")
	require.NoError(t, err)
	expectedOutput := getExpectedOutput()
	for i := 0; i < len(inputs); i++ {
		assert.Equal(t, expectedOutput[i], roadsAndLibraries(inputs[i].n, inputs[i].cLib, inputs[i].cRoad, inputs[i].cities))
	}
}

func getExpectedOutput() map[int]int64 {
	return map[int]int64{
		0: 8632543093,
		1: 5505097052,
		2: 903771485,
		3: 4679823300,
		4: 9567179141,
		5: 8070816272,
		6: 4812350616,
		7: 1143691734,
		8: 7760399472,
		9: 7939695500,
	}
}

type input struct {
	n      int32
	m      int32
	cLib   int32
	cRoad  int32
	cities [][]int32
}

func loadFile(path string) ([]input, error) {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var inputs []input
	var times int
	var inputData input
	var cities [][]int32

	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")

		if len(values) == 4 {
			if times > 0 {
				inputData.cities = cities
				inputs = append(inputs, inputData)
				inputData = input{}
				cities = [][]int32{}
			}
			times++
			inputData.n, inputData.m, inputData.cLib, inputData.cRoad, err = getConstants(values)
			if err != nil {
				return inputs, err
			}
		} else if len(values) == 2 {
			city, err := getCity(values)
			if err != nil {
				return inputs, err
			}
			cities = append(cities, city)
		}
	}
	inputData.cities = cities
	inputs = append(inputs, inputData)
	readFile.Close()

	return inputs, nil
}

func getCity(values []string) ([]int32, error) {
	sourceString, err := strconv.Atoi(values[0])
	if err != nil {
		return []int32{}, err
	}
	source := int32(sourceString)

	destinationString, err := strconv.Atoi(values[1])
	if err != nil {
		return []int32{}, err
	}
	destination := int32(destinationString)

	return []int32{source, destination}, nil
}

func getConstants(values []string) (int32, int32, int32, int32, error) {
	n, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	m, err := strconv.Atoi(values[1])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	cLib, err := strconv.Atoi(values[2])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	cRoad, err := strconv.Atoi(values[3])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return int32(n), int32(m), int32(cLib), int32(cRoad), nil
}
