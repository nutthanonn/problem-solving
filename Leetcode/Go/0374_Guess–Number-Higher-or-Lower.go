/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left := 1
	right := n

	if guess(n) == 0 {
		return n
	}

	for left != right {
		mid := (left + right) / 2
		g := guess(mid)
		if g == -1 {
			right = mid
		} else if g == 1 {
			left = mid
		} else {
			return mid
		}
	}

	return left
}