package main

func TruckTour(petrolpumps [][]int32) int32 {
	var start int32
	end := int32(1)
	size := int32(len(petrolpumps))
	balance := petrolpumps[0][0] - petrolpumps[0][1]

	for start != end {

		for start != end && balance < 0 {
			balance -= petrolpumps[start][0] - petrolpumps[start][1]
			start = (start + 1) % size
			if start == 0 {
				return -1
			}
		}

		balance += petrolpumps[end][0] - petrolpumps[end][1]
		end = (end + 1) % size
	}
	return start

}
