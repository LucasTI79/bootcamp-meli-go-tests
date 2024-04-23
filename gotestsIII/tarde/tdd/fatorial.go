package main

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	// implementação com complexidade de o(n)
	// return number * factorial(number-1)

	// implementação com complexidade de o(n)
	f := 1
	for i := 1; i <= number; i++ {
		f *= i
	}

	return f
}
