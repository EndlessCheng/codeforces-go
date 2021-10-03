package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	ans := []byte{}
	unread := strings.Count(s, string(letter)) // 未遍历到的 letter 个数
	inQ := 0                                   // 在单调队列中的 letter 个数
	q := []byte{}
	for i, n := 0, len(s); len(ans) < k; i++ {
		// 将 s[i] 插入单调队列
		for i < n && len(q) > 0 && s[i] < q[len(q)-1] {
			if q[len(q)-1] == letter { // 特判队尾是 letter 的情况
				if inQ+unread <= repetition { // 弹出队尾会导致无法凑足 repetition 个 letter
					break
				}
				inQ--
			}
			q = q[:len(q)-1]
		}
		if i < n {
			if s[i] == letter {
				unread--
				inQ++
			}
			q = append(q, s[i])
		}

		if i >= n-k {
			// 从队首拿到字典序最小的字符
			if q[0] == letter {
				inQ--
				repetition--
				ans = append(ans, q[0])
			} else if len(ans)+repetition < k {
				ans = append(ans, q[0])
			}
			q = q[1:]
		}
	}
	return string(ans)
}
