package main

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	j := 0
	for i := 1; i < l; i++ {
		if nums[j] != nums[i] {
			nums[j] = nums[i]
			j++
		}
	}

	return j + 1
}

func main() {

}
