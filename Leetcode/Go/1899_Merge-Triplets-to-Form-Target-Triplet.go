package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func mergeTriplets(triplets [][]int, target []int) bool {

	temp := [3]int{}

	for _, v := range triplets {
		if v[0] <= target[0] && v[1] <= target[1] && v[2] <= target[2] {
			temp = [3]int{max(v[0], temp[0]), max(v[1], temp[1]), max(v[2], temp[2])}
		}
	}

	for i := 0; i < 3; i++ {
		if target[i] != temp[i] {
			return false
		}
	}

	return true
}
