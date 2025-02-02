package main

import "math"

// https://space.bilibili.com/206214
func minCostGoodCaption(s string) string {
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
				res2 = f[i+3][minJ[i+3]] + abs(int(s[i]-'a')-j) + abs(int(s[i+1]-'a')-j) + abs(int(s[i+1]-'a')-j)
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

func abs(x int) int { if x < 0 { return -x }; return x }


func minCostGoodCaption2(s string) string {
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
