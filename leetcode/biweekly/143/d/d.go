package main

import (
	"bytes"
	"slices"
	"strings"
)

// https://space.bilibili.com/206214
func smallestNumber(num string, t int64) string {
	tmp := int(t)
	for i := 9; i > 1; i-- {
		for tmp%i == 0 {
			tmp /= i
		}
	}
	if tmp > 1 { // t 包含大于 7 的质因子
		return "-1"
	}

	n := len(num)
	leftT := make([]int, n+1)
	leftT[0] = int(t)
	i0 := n - 1
	for i, c := range num {
		if c == '0' {
			i0 = i
			break
		}
		leftT[i+1] = leftT[i] / gcd(leftT[i], int(c-'0'))
	}
	if leftT[n] == 1 { // num 的数位之积是 t 的倍数
		return num
	}

	// 假设答案和 num 一样长
	s := []byte(num)
	for i := i0; i >= 0; i-- {
		for s[i]++; s[i] <= '9'; s[i]++ {
			tt := leftT[i] / gcd(leftT[i], int(s[i]-'0'))
			k := 9
			for j := n - 1; j > i; j-- {
				for tt%k > 0 {
					k--
				}
				tt /= k
				s[j] = '0' + byte(k)
			}
			if tt == 1 {
				return string(s)
			}
		}
	}

	// 答案一定比 num 长
	ans := []byte{}
	for i := int64(9); i > 1; i-- {
		for t%i == 0 {
			ans = append(ans, '0'+byte(i))
			t /= i
		}
	}
	for len(ans) <= n {
		ans = append(ans, '1')
	}
	slices.Reverse(ans)
	return string(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func smallestNumber2(s string, t int64) string {
	tmp, cnt := int(t), 0
	for _, p := range []int{2, 3, 5, 7} {
		for tmp%p == 0 {
			tmp /= p
			cnt++
		}
	}
	if tmp > 1 { // t 包含其他质因子
		return "-1"
	}

	// 补前导零（至少一个）
	cnt = max(cnt-len(s)+1, 1)
	s = strings.Repeat("0", cnt) + s

	n := len(s)
	ans := bytes.Repeat([]byte{'0'}, n)
	type pair struct{ i, t int }
	vis := map[pair]bool{}

	var dfs func(int, int, bool) bool
	dfs = func(i, t int, isLimit bool) bool {
		if i == n {
			return t == 1
		}
		if !isLimit {
			p := pair{i, t}
			if vis[p] {
				return false
			}
			vis[p] = true
		}

		if isLimit && i < cnt && dfs(i+1, t, true) { // 填 0（跳过）
			return true
		}

		low := 0
		if isLimit {
			low = int(s[i] - '0')
		}
		for d := max(low, 1); d <= 9; d++ {
			if dfs(i+1, t/gcd(t, d), isLimit && d == low) {
				ans[i] = '0' + byte(d)
				return true
			}
		}
		return false
	}
	dfs(0, int(t), true)

	i := bytes.LastIndexByte(ans, '0')
	return string(ans[i+1:])
}
