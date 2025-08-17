package main

import "math"

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func xorAfterQueries(nums []int, queries [][]int) (ans int) {
	n := len(nums)
	B := int(math.Sqrt(float64(len(queries))))
	type tuple struct{ l, r, v int }
	groups := make([][]tuple, B)

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		if k < B {
			groups[k] = append(groups[k], tuple{l, r, v})
		} else {
			for i := l; i <= r; i += k {
				nums[i] = nums[i] * v % mod
			}
		}
	}

	diff := make([]int, n+1)
	for k, g := range groups {
		if g == nil {
			continue
		}
		buckets := make([][]tuple, k)
		for _, t := range g {
			buckets[t.l%k] = append(buckets[t.l%k], t)
		}
		for start, bucket := range buckets {
			if bucket == nil {
				continue
			}
			if len(bucket) == 1 {
				// 只有一个询问，直接暴力
				t := bucket[0]
				for i := t.l; i <= t.r; i += k {
					nums[i] = nums[i] * t.v % mod
				}
				continue
			}

			for i := range (n-start-1)/k + 1 {
				diff[i] = 1
			}
			for _, t := range bucket {
				diff[t.l/k] = diff[t.l/k] * t.v % mod
				r := (t.r-start)/k + 1
				diff[r] = diff[r] * pow(t.v, mod-2) % mod
			}

			mulD := 1
			for i := range (n-start-1)/k + 1 {
				mulD = mulD * diff[i] % mod
				j := start + i*k
				nums[j] = nums[j] * mulD % mod
			}
		}
	}

	for _, x := range nums {
		ans ^= x
	}
	return
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
