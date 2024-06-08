首先思考：如果每个字母至多出现一次，要怎么做？

如果对于每个字符串，字母 $x$ 都出现在字母 $y$ 的左边，那么就在 $x$ 到 $y$ 之间连一条有向边。

问题变成求有向无环图（DAG）的最长路，这可以用记忆化搜索（或者拓扑排序）解决。

回到本题。

比如有两个字符串，字母 $x$ 的在第一个字符串的下标为 $i_1$ 和 $i_2$，在第二个字符串中的下标为 $j_1$ 和 $j_2$，那么字母 $x$ 就有 $4$ 种不同的位置组合：

1. 位置为 $(i_1, j_1)$ 的 $x$。
2. 位置为 $(i_1, j_2)$ 的 $x$。
3. 位置为 $(i_2, j_1)$ 的 $x$。
4. 位置为 $(i_2, j_2)$ 的 $x$。

推广到 $n$ 个字符串，每个字母至多有 $2^n$ 种不同的位置组合。

用二进制数 $\textit{mask}$ 记录字母 $c$ 在各个字符串中的位置，记作 $(\textit{mask},c)$。

- $\textit{mask}$ 二进制从低到高第 $i$ 位为 $0$，表示 $c$ 位于 $s[i]$ 中的左边那个 $c$ 的下标。
- $\textit{mask}$ 二进制从低到高第 $i$ 位为 $1$，表示 $c$ 位于 $s[i]$ 中的右边那个 $c$ 的下标。

如果对于每个字符串，字母 $x$（位置组合为 $\textit{mask}$）都出现在字母 $y$（位置组合为 $\textit{mask}_2$）的左边，那么就在 $(\textit{mask},x)$ 到 $(\textit{mask}_2,y)$ 之间连一条有向边。同样地，问题变成求有向无环图（DAG）的最长路，这可以用记忆化搜索（或者拓扑排序）解决。

考虑用记忆化搜索计算。定义 $\textit{dfs}(\textit{mask},c)$ 表示 LCS 首字母为 $c$，位置组合为 $\textit{mask}$ 时的 LCS 长度。

枚举下一个字母 $\textit{ch}$，设 $\textit{ch}$ 在各个字符串中的位置为 $\textit{mask}_2$，则用 $\textit{dfs}(\textit{mask}_2,\textit{ch}) + 1$ 更新答案的最大值（下面代码把 $+1$ 放在返回前计算）。

注意这里 $\textit{mask}_2$ 要贪心地计算，如果两个 $\textit{ch}$ 都满足要求，优先取左边的。

答案为 $\textit{dfs}(0,0)-1$，这里取字符 `'\0'` 作为 LCS 的首字母，这样就无需在 $\textit{dfs}$ 外面枚举 LCS 的首字母了。（假设 `'\0'` 在各个字符串中的下标是 $-1$。）

为了输出具体方案，可以用一个 $\textit{from}$ 数组，在 $\textit{dfs}$ 中记录转移来源。

```go
package main
import ."fmt"

func main() {
	var T, n int
	var s string
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		pos := [123][10][]int{}
		for i := 0; i < n; i++ {
			pos[0][i] = []int{-1} // 假定在 LCS 前面还有个字符 '\0'，下标为 -1
			Scan(&s)
			for j, b := range s {
				pos[b][i] = append(pos[b][i], j) // 记录字母 b 在字符串 s[i] 中的出现位置 j
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
			// 如果有两个 ch 都满足要求，优先取左边的，对应下面代码中的 p[0] > cur
			for ch := byte('A'); ch <= 'z'; {
				mask2 := 0
				for i, p := range pos[ch][:n] {
					if p == nil {
						goto nxt
					}
					cur := pos[c][i][mask>>i&1] // 当前字母 c 的下标
					// p[0] 或者 p[1] 是下一个字母 ch 的下标
					if p[0] > cur {
						// 0
					} else if len(p) > 1 && p[1] > cur {
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
		Println(dfs(0, 0) - 1)

		lcs := []byte{}
		for p := from[0][0]; p.c > 0; p = from[p.mask][p.c] {
			lcs = append(lcs, p.c)
		}
		Printf("%s\n", lcs)
	}
}
```

时间复杂度：$\mathcal{O}(n2^n|\Sigma|^2)$。其中 $|\Sigma|$ 为字符集合的大小，本题 $|\Sigma|=52$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
