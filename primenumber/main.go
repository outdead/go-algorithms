package main

import (
	"fmt"
	"math"
)

func main() {
	needle := 719
	n := 137

	fmt.Printf("IsPrimeDivider ............. %v is prime == %v\n", needle, IsPrimeDivider(needle))
	fmt.Printf("IsPrimeEratosthenes ........ %v is prime == %v\n", needle, IsPrimeEratosthenes(needle))

	fmt.Printf("GeneratePrimeDivider ....... %v\n", GeneratePrimeDivider(n))
	fmt.Printf("GeneratePrimeEratosthenes .. %v\n", GeneratePrimeEratosthenes(n))
	fmt.Printf("GeneratePrimeAtkin ......... %v\n", GeneratePrimeAtkin(n))
}

// Check the number - prime or not with dividers method.
func IsPrimeDivider(n int) bool {
	// 0 and 1 - are not prime numbers
	if n <= 1 {
		return false
	}

	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

// Check the number - prime or not with the Eratosthenes sieve.
func IsPrimeEratosthenes(n int) bool {
	if n <= 1 {
		return false
	}

	sieve := make([]int, n + 1);

	// fill sieve with 1
	for k := 2; k <= n; k++ {
		sieve[k] = 1;
	}

	for k := 2; k * k <= n; k++ {
		// if k is prime (not crossed out)
		if sieve[k] == 1 {
			// then cross out k
			for l := k * k; l <= n; l += k {
				sieve[l] = 0;
			}
		}
	}

	return sieve[n] == 1
}

// Generate prime numbers to n with dividers method
func GeneratePrimeDivider(n int) (primes []int) {
	for i := 2; i <= n; i++ {
		if IsPrimeDivider(i) {
			primes = append(primes, i)
		}
	}

	return
}

// Generate prime numbers to n with the Eratosthenes sieve.
func GeneratePrimeEratosthenes(n int) (primes []int) {
	b := make([]bool, n + 1)

	for i := 2; i <= n; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k <= n; k += i {
			b[k] = true
		}
	}

	return
}

// Generate prime numbers to n with the Atkin sieve.
//
// In theory it should be faster than the Eratosthenes sieve.
// But coefficient of depreciation turned out to be very large,
// and the algorithm is actually slower
func GeneratePrimeAtkin(n int) (primes []int) {
	var x, y, k int
	nsqrt := math.Sqrt(float64(n))

	is_prime := make([]bool, n + 1)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			k = 4 * (x * x) + y * y
			if k <= n && (k % 12 == 1 || k % 12 == 5) {
				is_prime[k] = !is_prime[k]
			}
			k = 3 * (x * x) + y * y
			if k <= n && k % 12 == 7 {
				is_prime[k] = !is_prime[k]
			}
			k = 3 * (x * x) - y * y
			if x > y && k <= n && k % 12 == 11 {
				is_prime[k] = !is_prime[k]
			}
		}
	}

	for k = 5; float64(k) <= nsqrt; k++ {
		if is_prime[k] {
			for y = k * k; y < n; y += k * k {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	for x = 0; x <= len(is_prime) - 1; x++ {
		if is_prime[x] {
			primes = append(primes, x)
		}
	}

	return
}
