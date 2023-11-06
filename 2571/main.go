func minOperations(n int) int {
	opCount := 0
	for n != 0 {
		for n & 1 == 0 {
			n >>= 1
		}
		oneCount := 0
		for n & 1 == 1 {
			oneCount++
			n >>= 1
		}
		if oneCount >= 2 {
			opCount += 1
			n++
		} else {
			opCount += 1
		}
	}
	return opCount
}

// Magic:
// func minOperations(n int) int {
// 	return bits.OnesCount(uint(3*n ^ n))
// }