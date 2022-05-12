package main

func pairSums(k int32, n int32, arr []int32) int32 {
	var output int32

	if len(arr) < 2 {
		return output
	}

	weights := make(map[int32]int32)
	for _, v := range arr {
		weights[v] = weights[v] + 1
	}

	for _, v := range arr {
		if v > k {
			continue
		}

		i := k - v
		w, ok := weights[i]
		if !ok {
			continue
		}

		if i == v {
			w = w - 1
		}

		output += w
	}
	return output / 2
}
func main() {}
