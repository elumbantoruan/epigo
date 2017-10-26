package ch12searching

func ComputeSquareRoot(n int) int {
	floor := 0
	ceiling := n

	for floor <= ceiling {
		mid := floor + (ceiling - floor) / 2
		product := mid * mid
		if product <= n {
			floor = mid + 1
		} else {
			ceiling = mid - 1
		}
	}
	return floor - 1
}
