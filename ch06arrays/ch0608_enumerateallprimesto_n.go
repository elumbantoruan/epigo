package ch06arrays

func GeneratePrimes(n int) []int {
	var results []int
	var isprimes = initializeIsPrimes(n)
	for i := 2; i <= n; i++ {
		if isprimes[i] {
			results = append(results, i)
		}
		for j := i; j <= n; j += i {
			isprimes[j] = false
		}
	}
	return results
}

func initializeIsPrimes(n int) []bool {
	var isPrimes []bool
	for i := 0; i <= n; i++ {
		isPrimes = append(isPrimes, true)
	}
	return isPrimes
}
