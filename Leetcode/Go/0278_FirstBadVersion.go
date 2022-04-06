package main

func isBadVersion(n int) bool

func firstBadVersion(n int) int {
	first := 1
	last := n
	for first <= last {
		mid := int((first + last) / 2)
		if !isBadVersion(mid) {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return first
}
func main() {

}
