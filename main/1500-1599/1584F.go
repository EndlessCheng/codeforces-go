package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1584F(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		// 记录字母在各个字符串中的出现位置
		pos := [123][10][]int{}
		for i := 0; i < n; i++ {
			pos[0][i] = []int{-1} // 假定在 LCS 前面还有个字符 \0，下标为 -1
			Fscan(in, &s)
			for j, b := range s {
				pos[b][i] = append(pos[b][i], j)
			}
		}

		memo := make([][123]int, 1<<n)
		for i := range memo {
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		type pair struct{ mask int; c byte }
		from := make([][123]pair, 1<<n) // 记录转移来源
		var dfs func(int, byte) int
		dfs = func(mask int, c byte) (res int) {
			p := &memo[mask][c]
			if *p != -1 {
				return *p
			}
			var frm pair
			// 枚举 LCS 的下一个字母 ch
			// 要求：ch 在所有字符串中的下标 > c 在对应字符串中的下标
			// 如果有两个 ch 都满足要求，优先取左边的，对应下面代码中的 p[0] > pos[c][i][mask>>i&1]
			for ch := byte('A'); ch <= 'z'; {
				mask2 := 0
				for i, p := range pos[ch][:n] {
					if p == nil {
						goto nxt
					}
					if p[0] > pos[c][i][mask>>i&1] {
						// 0
					} else if len(p) > 1 && p[1] > pos[c][i][mask>>i&1] {
						mask2 |= 1 << i
					} else {
						goto nxt
					}
				}
				if r := dfs(mask2, ch); r > res {
					res = r
					frm.mask = mask2 // 记录转移来源
					frm.c = ch
				}
			nxt:
				if ch == 'Z' {
					ch = 'a'
				} else {
					ch++
				}
			}
			from[mask][c] = frm
			res++
			*p = res
			return
		}
		Fprintln(out, dfs(0, 0)-1)

		lcs := []byte{}
		for p := from[0][0]; p.c > 0; p = from[p.mask][p.c] {
			lcs = append(lcs, p.c)
		}
		Fprintf(out, "%s\n", lcs)
	}
}

//func main() { cf1584F(bufio.NewReader(os.Stdin), os.Stdout) }
