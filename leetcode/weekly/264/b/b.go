package main

import (
	"bytes"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func nextBeautifulNumber1(n int) int {
next:
	for {
		n++
		cnt := [10]int{}
		for x := n; x > 0; x /= 10 {
			cnt[x%10]++
		}
		for x := n; x > 0; x /= 10 {
			if cnt[x%10] != x%10 {
				continue next
			}
		}
		return n
	}
}

// 从 a 中选一个字典序最小的、元素和等于 target 的子序列
// a 已经从小到大排序
// 无解返回 nil
func zeroOneKnapsack(a []int, target int) []int {
	n := len(a)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, target+1)
	}
	f[n][0] = true

	// 倒着 DP，这样后面可以正着（从小到大）选
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for j := range f[i] {
			if j < v {
				f[i][j] = f[i+1][j]
			} else {
				f[i][j] = f[i+1][j] || f[i+1][j-v]
			}
		}
	}

	if !f[0][target] {
		return nil
	}

	ans := []int{}
	j := target
	for i, v := range a {
		if j >= v && f[i+1][j-v] {
			ans = append(ans, v)
			j -= v
		}
	}
	return ans
}

func nextBeautifulNumber(n int) int {
	// 补一个前导零，方便处理答案十进制比 n 的十进制长的情况
	s := "0" + strconv.Itoa(n)
	m := len(s)

	const mx = 10
	cnt := make([]int, mx)
	for i := 1; i < m; i++ {
		cnt[s[i]-'0']++
	}

	// 从右往左尝试
	for i := m - 1; i >= 0; i-- {
		if i > 0 {
			cnt[s[i]-'0']-- // 撤销
		}

		// 增大 s[i] 为 j
		for j := s[i] - '0' + 1; j < mx; j++ {
			cnt[j]++

			// 后面 [i+1, m-1] 需要补满 0 < cnt[k] < k 的数字 k，剩余数位可以随便填
			free := m - 1 - i // 统计可以随便填的数位个数
			for k, c := range cnt {
				if k < c { // 不合法
					free = -1
					break
				}
				if c > 0 {
					free -= k - c
				}
			}
			if free < 0 { // 不合法，继续枚举
				cnt[j]--
				continue
			}

			// 对于可以随便填的数位，计算字典序最小的填法
			a := []int{}
			for k := 1; k < mx; k++ {
				if cnt[k] == 0 {
					a = append(a, k)
				}
			}
			missing := zeroOneKnapsack(a, free)
			if missing == nil { // 无解，继续枚举
				cnt[j]--
				continue
			}

			for _, v := range missing {
				cnt[v] = -v // 用负数表示可以随便填的数
			}

			t := []byte(s[:i+1])
			t[i] = '0' + byte(j)
			for k, c := range cnt {
				if c > 0 {
					c = k - c
				} else {
					c = -c
				}
				d := []byte{'0' + byte(k)}
				t = append(t, bytes.Repeat(d, c)...)
			}
			ans, _ := strconv.Atoi(string(t))
			return ans
		}
	}
	return -1 // 无解（本题不会发生，但为了可扩展性保留）
}
