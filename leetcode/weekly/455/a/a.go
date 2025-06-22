package main

// https://space.bilibili.com/206214
const mx = 101

var np = [mx]bool{1: true}

func init() {
	// 质数=false 非质数=true
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func checkPrimeFrequency(nums []int) bool {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for _, c := range cnt {
		if !np[c] {
			return true
		}
	}
	return false
}
