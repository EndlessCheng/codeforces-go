package main

import "strings"

// https://space.bilibili.com/206214
func removeSubstring(s string, k int) string {
	type pair struct {
		b   rune
		cnt int
	}
	st := []pair{}
	for _, b := range s {
		if len(st) > 0 && st[len(st)-1].b == b {
			st[len(st)-1].cnt++ // 累计相同括号
		} else {
			st = append(st, pair{b, 1}) // 新的括号
		}
		// 栈顶的 k 个右括号与（栈顶下面的）k 个左括号抵消
		if b == ')' && len(st) > 1 && st[len(st)-1].cnt == k && st[len(st)-2].cnt >= k {
			st = st[:len(st)-1]
			st[len(st)-1].cnt -= k
			if st[len(st)-1].cnt == 0 {
				st = st[:len(st)-1]
			}
		}
	}

	ans := []byte{}
	for _, p := range st {
		ans = append(ans, strings.Repeat(string(p.b), p.cnt)...)
	}
	return string(ans)
}
