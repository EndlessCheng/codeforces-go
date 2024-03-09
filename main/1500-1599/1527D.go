package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1527D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := make([]int, n+1)
		dfn := make([]int, n)
		size := make([]int, n+1)
		t := 0
		var dfs func(int, int) int
		dfs = func(v, fa int) (s int) {
			t++
			dfn[v] = t
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				sz := dfs(w, v)
				if v == 0 {
					ans[0] += sz * (sz - 1) / 2 // 不含 0
				}
				s += sz
			}
			s++
			size[v] = s
			return
		}
		dfs(0, -1)
		isAncestor := func(f, v int) bool { return dfn[f] < dfn[v] && dfn[v] < dfn[f]+size[f] }

		p, q, ap := 0, 0, 0
		s := 1
		for _, w := range g[0] {
			sz := size[w]
			if w == 1 || isAncestor(w, 1) {
				ap = w // ap 是 0 的儿子和 1 的祖先（p 的祖先）
				sz -= size[1] // 不含 1
			}
			ans[1] += sz * s // 含 0 不含 1
			s += sz
		}

		for i := 1; i < n; {
			if q == 0 && isAncestor(p, i) { // i 在 p 下面
				p = i
				for i++; i < n && isAncestor(i, p); i++ { // i 在 0 到 p 的路径上
				}
				// 含 0~i-1 不含 i
				if i == n || isAncestor(p, i) { // i 还在 p 下面
					ans[i] = (n - size[ap]) * (size[p] - size[i])
				} else {
					if i == ap || isAncestor(ap, i) { // LCA(p,i) 是 ap 的后代
						ans[i] = (n - size[ap]) * size[p]
						break
					}
					// i 在其它子树
					ans[i] = (n - size[ap] - size[i]) * size[p]
				}
				continue
			}
			if isAncestor(p, i) { // i 在 p 下面
				p = i
				for i++; i < n && (isAncestor(i, p) || isAncestor(i, q)); i++ { // i 在 p 到 q 的路径上
				}
			} else { // i 在 q 下面
				q = i
				for i++; i < n && (isAncestor(i, p) || isAncestor(i, q)); i++ { // i 在 p 到 q 的路径上
				}
			}
			// 含 0~i-1 不含 i
			if i < n && isAncestor(p, i) { // i 在 p 下面
				ans[i] = (size[p] - size[i]) * size[q]
			} else if i < n && isAncestor(q, i) { // i 在 q 下面
				ans[i] = (size[q] - size[i]) * size[p]
			} else { // i 在其它子树
				ans[i] = size[p] * size[q]
				break
			}
		}

		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1527D(os.Stdin, os.Stdout) }
