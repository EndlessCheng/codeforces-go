package main

// https://space.bilibili.com/206214
const mx = 500_000

var primes []int
var np [mx + 1]bool
var ans [mx + 1]int

func init() {
	for i := 2; i <= mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}

	sum, j := 0, 0
	for i := 2; i <= mx; i++ {
		if sum+primes[j] <= i {
			sum += primes[j]
			j++
		}
		if !np[sum] {
			ans[i] = sum
		} else {
			ans[i] = ans[i-1]
		}
	}
}

func largestPrime(n int) int {
	return ans[n]
}
