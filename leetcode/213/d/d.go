package main

import "math/big"

// github.com/EndlessCheng/codeforces-go
func kthSmallestPath(destination []int, k int) (ans string) {
	n, m := destination[0], destination[1]
	i, j := 0, 0
	for i < n || j < m {
		if i == n {
			ans += "H"
			j++
		} else if j == m {
			ans += "V"
			i++
		} else {
			c := int(new(big.Int).Binomial(int64(n-i+m-j-1), int64(n-i)).Int64())
			if k > c {
				k -= c
				ans += "V"
				i++
			} else {
				ans += "H"
				j++
			}
		}
	}
	return
}
