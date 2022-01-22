package main

import (
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var vals, f [1 << len(primes)]int
var validSubs [1 << len(primes)][]int
var p2 = [1e5 + 1]int{1}

func init() {
	f[0] = 1
	for i := 1; i < len(vals); i++ {
		vals[i] = 1
		for s := uint(i); s > 0; s &= s - 1 {
			vals[i] *= primes[bits.TrailingZeros(s)]
			if vals[i] > 30 {
				vals[i] = 0
				break
			}
		}
		top := 1 << (bits.Len(uint(i)) - 1)
		for sub := i; sub&top > 0; sub = (sub - 1) & i {
			if v := vals[sub]; v > 0 {
				validSubs[i] = append(validSubs[i], sub)
			}
		}
	}
	for i := 1; i <= 1e5; i++ {
		p2[i] = p2[i-1] * 2 % mod
	}
}

func numberOfGoodSubsets(a []int) (ans int) {
	cnt := [31]int{}
	for _, v := range a {
		cnt[v]++
	}
	for s := 1; s < len(f); s++ {
		f[s] = 0
		for _, sub := range validSubs[s] {
			f[s] = (f[s] + cnt[vals[sub]]*f[s^sub]) % mod
		}
		ans += f[s]
	}
	return ans % mod * p2[cnt[1]] % mod
}
