package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4, 5, 8}, bubbleSort([]int{5, 1, 4, 2, 8}))
	assert.Equal(t, []int{1, 2, 5, 8, 9}, bubbleSort([]int{5, 1, 9, 2, 8}))
}

func TestInsertionSort(t *testing.T) {
	assert.Equal(t, []int{5, 6, 11, 12, 13}, insertionSort([]int{12, 11, 13, 5, 6}))
	assert.Equal(t, []int{1, 2, 5, 8, 9}, insertionSort([]int{5, 1, 9, 2, 8}))
}

func TestMergeSort(t *testing.T) {
	assert.Equal(t, []int{12}, mergeSort([]int{12}))
	assert.Equal(t, []int{5, 11, 12, 13}, mergeSort([]int{12, 11, 13, 5}))
	assert.Equal(t, []int{5, 6, 11, 12, 13}, mergeSort([]int{12, 11, 13, 5, 6}))
	assert.Equal(t, []int{1, 2, 5, 8, 9}, mergeSort([]int{5, 1, 9, 2, 8}))
}

func TestQuickSort(t *testing.T) {
	assert.Equal(t, []int{12}, quickSort([]int{12}))
	assert.Equal(t, []int{5, 11, 12, 13}, quickSort([]int{12, 11, 13, 5}))
	assert.Equal(t, []int{5, 6, 11, 12, 13}, quickSort([]int{12, 11, 13, 5, 6}))
	assert.Equal(t, []int{1, 2, 5, 8, 9}, quickSort([]int{5, 1, 9, 2, 8}))
}
