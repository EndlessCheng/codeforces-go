package main

import (
	"bytes"
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minCostGoodCaption(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	f := make([]int, n+1)
	f[n-1], f[n-2] = math.MaxInt/2, math.MaxInt/2
	t := make([]byte, n+1)
	size := make([]uint8, n)

	for i := n - 3; i >= 0; i-- {
		sub := []byte(s[i : i+3])
		slices.Sort(sub)
		a, b, c := sub[0], sub[1], sub[2]
		s3 := int(t[i+3])
		res := f[i+3] + int(c-a)
		mask := int(b)<<24 | s3<<16 | s3<<8 | s3 // 4 个 byte 压缩成一个 int，方便比较字典序
		size[i] = 3

		if i+4 <= n {
			sub := []byte(s[i : i+4])
			slices.Sort(sub)
			a, b, c, d := sub[0], sub[1], sub[2], sub[3]
			s4 := int(t[i+4])
			res4 := f[i+4] + int(c-a+d-b)
			mask4 := int(b)<<24 | int(b)<<16 | s4<<8 | s4
			if res4 < res || res4 == res && mask4 < mask {
				res, mask = res4, mask4
				size[i] = 4
			}
		}

		if i+5 <= n {
			sub := []byte(s[i : i+5])
			slices.Sort(sub)
			a, b, c, d, e := sub[0], sub[1], sub[2], sub[3], sub[4]
			res5 := f[i+5] + int(d-a+e-b)
			mask5 := int(c)<<24 | int(c)<<16 | int(c)<<8 | int(t[i+5])
			if res5 < res || res5 == res && mask5 < mask {
				res, mask = res5, mask5
				size[i] = 5
			}
		}

		f[i] = res
		t[i] = byte(mask >> 24)
	}

	ans := make([]byte, 0, n)
	for i := 0; i < n; i += int(size[i]) {
		ans = append(ans, bytes.Repeat([]byte{t[i]}, int(size[i]))...)
	}
	return string(ans)
}

func minCostGoodCaption3(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	f := make([]int, n+1)
	f[n-1], f[n-2] = math.MaxInt/2, math.MaxInt/2
	t := make([]byte, n+1)
	size := make([]uint8, n)

	for i := n - 3; i >= 0; i-- {
		sub := []byte(s[i : i+3])
		slices.Sort(sub)
		a, b, c := sub[0], sub[1], sub[2]
		s3 := int(t[i+3])
		res := []int{f[i+3] + int(c-a), int(b), s3, s3, s3}
		size[i] = 3

		if i+4 <= n {
			sub := []byte(s[i : i+4])
			slices.Sort(sub)
			a, b, c, d := sub[0], sub[1], sub[2], sub[3]
			s4 := int(t[i+4])
			tp := []int{f[i+4] + int(c-a+d-b), int(b), int(b), s4, s4}
			if slices.Compare(tp, res) < 0 {
				res = tp
				size[i] = 4
			}
		}

		if i+5 <= n {
			sub := []byte(s[i : i+5])
			slices.Sort(sub)
			a, b, c, d, e := sub[0], sub[1], sub[2], sub[3], sub[4]
			tp := []int{f[i+5] + int(d-a+e-b), int(c), int(c), int(c), int(t[i+5])}
			if slices.Compare(tp, res) < 0 {
				res = tp
				size[i] = 5
			}
		}

		f[i] = res[0]
		t[i] = byte(res[1])
	}

	ans := make([]byte, 0, n)
	for i := 0; i < n; i += int(size[i]) {
		ans = append(ans, bytes.Repeat([]byte{t[i]}, int(size[i]))...)
	}
	return string(ans)
}

func minCostGoodCaption2(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	f := make([][26]int, n+1)
	minJ := make([]int, n+1)
	nxt := make([][26]int, n+1)
	for i := n - 1; i >= 0; i-- {
		mn := math.MaxInt
		for j := 0; j < 26; j++ {
			res := f[i+1][j] + abs(int(s[i]-'a')-j)
			res2 := math.MaxInt
			if i <= n-6 {
				res2 = f[i+3][minJ[i+3]] + abs(int(s[i]-'a')-j) + abs(int(s[i+1]-'a')-j) + abs(int(s[i+2]-'a')-j)
			}
			if res2 < res || res2 == res && minJ[i+3] < j {
				res = res2
				nxt[i][j] = minJ[i+3]
			} else {
				nxt[i][j] = j
			}
			f[i][j] = res
			if res < mn {
				mn = res
				minJ[i] = j
			}
		}
	}

	ans := make([]byte, n)
	i, j := 0, minJ[0]
	for i < n {
		ans[i] = 'a' + byte(j)
		k := nxt[i][j]
		if k == j {
			i++
		} else {
			ans[i+1] = ans[i]
			ans[i+2] = ans[i]
			i += 3
			j = k
		}
	}
	return string(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minCostGoodCaption1(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	memo := make([][27]int, n)
	from := make([][27]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == n {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		x := int(s[i] - 'a')

		if i > n-3 {
			from[i][j] = j
			return dfs(i+1, j) + abs(x-j)
		}

		res = math.MaxInt
		for k := range 26 {
			var r int
			if k == j {
				r = dfs(i+1, k) + abs(x-k)
			} else {
				r = dfs(i+3, k) + abs(x-k) + abs(int(s[i+1]-'a')-k) + abs(int(s[i+2]-'a')-k)
			}
			if r < res {
				res = r
				from[i][j] = k
			}
		}
		return
	}
	dfs(0, 26)

	ans := make([]byte, n)
	i, j := 0, 26
	for i < n {
		k := from[i][j]
		ans[i] = 'a' + byte(k)
		if k == j {
			i++
		} else {
			ans[i+1] = ans[i]
			ans[i+2] = ans[i+1]
			i += 3
			j = k
		}
	}
	return string(ans)
}
