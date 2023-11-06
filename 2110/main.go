func getDescentPeriods(prices []int) int64 {
	left := 0
	right := 1
	result := 0
	for left < len(prices) {
		for right < len(prices) && prices[right - 1] == prices[right] + 1 {
			right++
		}
		count := right - left
		result += count * (1 + count) / 2
		left = right
		right = left + 1
	}
	return int64(result)
}