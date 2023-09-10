package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func newIdentityMatrix(n int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int64) matrix {
	res := newIdentityMatrix(len(a))
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func numberOfWays(s, t string, k int64) int {
	n := len(s)
	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	kmpSearch := func(text, pattern string) (cnt int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for _, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				cnt++
				c = match[c-1]
			}
		}
		return
	}
	c := kmpSearch(s+s[:n-1], t)
	m := matrix{
		{c - 1, c},
		{n - c, n - 1 - c},
	}.pow(k)
	if s == t {
		return m[0][0]
	}
	return m[0][1]
}
