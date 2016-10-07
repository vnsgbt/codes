package exercise

// Sum ...
func Sum(n int) int {
	if n < 1 {
		return 0
	}
	return n + Sum(n-1)
}

// PairSumSeq ...
func PairSumSeq(n int) int {
	pairSum := func(a int, b int) int {
		return a + b
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += pairSum(i, i+1)
	}
	return sum
}

// TreeInt ...
func TreeInt(n int) int {
	if n <= 1 {
		return 1
	}
	return 1 + TreeInt(n-1) + TreeInt(n-1)
}
