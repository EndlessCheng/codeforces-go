package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
const mx = 1001
var np = [mx]bool{}
var primes []int

func init() {
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func maxPrimeSubset(n int, s int) (ans []int) {
	m := sort.SearchInts(primes, n+1)

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, s+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2 // 避免加法溢出
		}
	}
	f[m][0] = 0

	// 倒着 DP，这样可以正序输出方案，且保证越靠前的质数越小
	for i := m - 1; i >= 0; i-- {
		x := primes[i]
		for j := range s + 1 {
			if j < x {
				f[i][j] = f[i+1][j]
			} else {
				f[i][j] = min(f[i+1][j], f[i+1][j-x]+1)
			}
		}
	}

	j := s
	for j > 0 && f[0][j] == math.MaxInt/2 {
		j--
	}

	for i := 0; j > 0; i++ {
		x := primes[i]
		// 靠前的质数越小越好，f[i+1][j-x]+1 = f[i+1][j] 的时候要选
		if f[i+1][j-x]+1 <= f[i+1][j] {
			ans = append(ans, x)
			j -= x
		}
	}

	return
}
