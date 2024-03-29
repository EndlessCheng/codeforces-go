package main

// github.com/EndlessCheng/codeforces-go
const p = 2021

//func f(n, m int64) (c int64) {
//	for i := int64(1); i <= n; i++ {
//		if i%43 == 0 {
//			c += m / 47
//		} else if i%47 == 0 {
//			c += m / 43
//		} else {
//			c += m / p
//		}
//	}
//	return
//}
//
//func q(n, m int64) int64 { return n/p*(m/p)*7905 + f(n%p, m-m%p) + f(m%p, n-n%p) + f(n%p, m%p) }

// O(1) 解法
func q(a, b int64) int64 {
	return a/p*b + b/p*a - a/p*(b/p) + (a/43-a/p)*(b/47-b/p) + (a/47-a/p)*(b/43-b/p)
}

func findPairs(a, b, c, d int64) int64 { return q(b, d) - q(a-1, d) - q(b, c-1) + q(a-1, c-1) }
