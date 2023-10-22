package main

// https://space.bilibili.com/206214
// 预处理每个数的不包括自己的因子，时间复杂度 O(mx*log(mx))
const mx = 200

var divisors [mx + 1][]int

func init() {
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func calc(s string) int {
	n := len(s)
	res := n
	for _, d := range divisors[n] {
		cnt := 0
		for i0 := 0; i0 < d; i0++ {
			for i, j := i0, n-d+i0; i < j; i, j = i+d, j-d {
				if s[i] != s[j] {
					cnt++
				}
			}
		}
		res = min(res, cnt)
	}
	return res
}

func minimumChanges(s string, k int) (ans int) {
	n := len(s)
	modify := make([][]int, n-1)
	for l := range modify {
		modify[l] = make([]int, n)
		for r := l + 1; r < n; r++ { // 半回文串长度至少为 2
			modify[l][r] = calc(s[l : r+1])
		}
	}

	f := modify[0]
	for i := 1; i < k; i++ {
		for j := n - 1 - (k-1-i)*2; j > i*2; j-- {
			f[j] = n
			for L := i * 2; L < j; L++ {
				f[j] = min(f[j], f[L-1]+modify[L][j])
			}
		}
	}
	return f[n-1]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
