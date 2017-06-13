//
// Package primenumber implements sieve of Eratosthenes and sieve of Atkin algorithms.
//
package primenumber

import "math"

// IsPrime checks the number and returns true in number is prime number.
// Method implements trial division algorithm.
func IsPrime(number int) bool {
	// 0 and 1 - are not prime numbers.
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

// IsPrimeWithSieveOfEratosthenes checks the number and returns true in number
// is prime number. Method implements the sieve of Eratosthenes algorithm.
// Too slow to use
/*func IsPrimeWithSieveOfEratosthenes(number int) bool {
	// 0 and 1 - are not prime numbers.
	if number <= 1 {
		return false
	}

	b := make([]bool, number+1)

	for i := 2; i <= number; i++ {
		if b[i] == true {
			continue
		}

		if i == number {
			return true
		}

		for k := i * i; k <= number; k += i {
			b[k] = true
		}
	}

	return false
}*/

// GetSequenceWithTrialDivision finds and returns the prime numbers up to
// a given integer limit with trial division algorithm.
// Expected asymptotic complexity:
//
// Data structure              Array
// Best-case performance       O(n^(1/2))
// Average performance         O(n^(1/2))
// Worst-case performance      O(n^(1/2))
// Worst-case space complexity O(n)
func GetSequenceWithTrialDivision(limit int) (sequence []int) {
	for i := 2; i <= limit; i++ {
		if IsPrime(i) {
			sequence = append(sequence, i)
		}
	}

	return
}

// GetSequenceWithSieveOfEratosthenes finds and returns the prime numbers
// up to a given integer limit with the sieve of Eratosthenes algorithm.
// Expected asymptotic complexity:
//
// Data structure              Array
// Best-case performance       O(n log log n)
// Average performance         O(n log log n)
// Worst-case performance      O(n log log n)
// Worst-case space complexity O((n^(1/2)) / log n) bits
func GetSequenceWithSieveOfEratosthenes(limit int) (sequence []int) {
	b := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		if b[i] == true {
			continue
		}
		sequence = append(sequence, i)
		for k := i * i; k <= limit; k += i {
			b[k] = true
		}
	}

	return
}

// GetSequenceWithSieveOfAtkin finds and returns the prime numbers up to
// a given integer limit with the sieve of Atkin algorithm. In theory it
// should be faster than the Eratosthenes sieve. But the implementation
// is actually slower.
// Expected asymptotic complexity:
//
// Data structure              Array
// Best-case performance       O(n / (log log n))
// Average performance         O(n / (log log n))
// Worst-case performance      O(n / (log log n))
// Worst-case space complexity O((n^(1/2)) / log n) bits
func GetSequenceWithSieveOfAtkin(limit int) (sequence []int) {
	limitSqrt := int(math.Sqrt(float64(limit)))
	isPrime := make([]bool, limit+1)

	for x := 1; x <= limitSqrt; x++ {
		for y := 1; y <= limitSqrt; y++ {
			k := 4*(x*x) + y*y
			if k <= limit && (k%12 == 1 || k%12 == 5) {
				isPrime[k] = !isPrime[k]
			}

			k = 3*(x*x) + y*y
			if k <= limit && k%12 == 7 {
				isPrime[k] = !isPrime[k]
			}

			k = 3*(x*x) - y*y
			if x > y && k <= limit && k%12 == 11 {
				isPrime[k] = !isPrime[k]
			}
		}
	}

	for k := 5; k <= limitSqrt; k++ {
		if isPrime[k] {
			for y := k * k; y < limit; y += k * k {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	for x := 2; x <= len(isPrime)-1; x++ {
		if isPrime[x] {
			sequence = append(sequence, x)
		}
	}

	return
}
