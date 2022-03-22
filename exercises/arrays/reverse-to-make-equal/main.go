package main

/*
Reverse to Make Equal

Given two arrays A and B of length N, determine if there is a way to make A equal to B by reversing any subarrays from array B any number of times.

Signature
	bool areTheyEqual(int[] arr_a, int[] arr_b)
Input
	All integers in array are in the range [0, 1,000,000,000].
Output
	Return true if B can be made equal to A, return false otherwise.

Example
	A = [1, 2, 3, 4]
	B = [1, 4, 3, 2]
	output = true
	After reversing the subarray of B from indices 1 to 3, array B will equal array A.
*/
func areTheyEqual(arr_1 []int, arr_b []int) bool {
	if len(arr_1) != len(arr_b) {
		return false
	}

	m := make(map[int]int)

	for _, number := range arr_1 {
		counter := m[number]
		m[number] = counter + 1
	}

	for _, key := range arr_b {
		if _, ok:= m[key]; !ok {
			return false
		}
		delete(m, key)
	}

	if len(m) > 0 {
		return false
	}
	return true
}