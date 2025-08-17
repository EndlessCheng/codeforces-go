package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1312G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, v int
	var w string
	Fscan(in, &n)
	f := make([]int, n+1)
	son := make([][26]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v, &w)
		son[v][w[0]-'a'] = i
	}

	Fscan(in, &m)
	idx := make([]int, m)
	inS := make([]bool, n+1)
	for i := range idx {
		Fscan(in, &idx[i])
		inS[idx[i]] = true
	}

	var dfs func(int, int) int
	dfs = func(v, kth int) (sz int) {
		// 如果 v 是目标字符串，则 v 会作为自动补全的一个选项
		// kth 是到达这个选项列表前的成本，v 是列表的第一个，所以总成本是 kth+1
		if inS[v] {
			sz = 1 // 自动补全列表的大小
			f[v] = min(f[v], kth+1)
		}
		kth = min(kth, f[v]) // min(从祖先跳过来，先跳到 v 再继续跳）
		for _, w := range son[v][:] {
			if w > 0 {
				f[w] = f[v] + 1
				sz += dfs(w, kth+sz)
			}
		}
		return
	}
	dfs(0, 0)

	for _, i := range idx {
		Fprint(out, f[i], " ")
	}
}

//func main() { cf1312G(bufio.NewReader(os.Stdin), os.Stdout) }
