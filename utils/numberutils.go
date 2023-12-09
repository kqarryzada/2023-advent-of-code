package fileutils

// FindPrimeFactors computes the set of prime factors of a number and returns
// the factors as a slice of unsigned integers.
func FindPrimeFactors(number int) []uint64 {
	factors := make([]uint64, 0)
	for (number % 2) == 0 {
		factors = append(factors, 2)
		number /= 2
	}

	for i := 3; (i * i) <= number; i += 2 {
		// For each number that is divisible by 'i', check multiple times in
		// case the factor appears twice (e.g., 36 == 3 * 3 * 2)
		for (number % i) == 0 {
			factors = append(factors, uint64(i))
			number /= i
		}
	}

	// If no common factors were found (i.e., 'number' is prime), it will be
	// greater than 2.
	if number > 2 {
		factors = append(factors, uint64(number))
	}

	return factors
}

// FindLCM computes the least common multiple among a set of numbers.
func FindLCM(numbers []int) uint64 {
	factorMap := make(map[uint64]int, 0)
	for _, number := range numbers {
		factors := FindPrimeFactors(number)
		for _, n := range factors {
			factorMap[n] = 1
		}
	}
	lcm := uint64(1)
	for primeFactor, _ := range factorMap {
		lcm *= primeFactor
	}

	return lcm
}
