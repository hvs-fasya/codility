package binary_gap

func Solution(n int) int {
	cur := n & 1
	for n >= 1 {
		if cur == 1 {
			break
		}
		n = n >> 1
		cur = n & 1
	}
	maxCount := 0
	count := 0
	for n >= 1 {
		if cur == 1 {
			count = 0
		} else {
			count++
		}
		if count >= maxCount {
			maxCount = count
		}
		n = n >> 1
		cur = n & 1
	}
	return maxCount
}
