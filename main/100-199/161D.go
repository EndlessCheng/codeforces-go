package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF161D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, w, ans int
	Fscan(in, &n, &k)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	hson := make([]int, n)
	var _f func(v, fa int) int
	_f = func(v, fa int) int {
		maxDep, hs := 0, -1
		for _, w := range g[v] {
			if w != fa {
				if mxD := _f(w, v); mxD > maxDep {
					maxDep, hs = mxD, w
				}
			}
		}
		hson[v] = hs
		return maxDep + 1
	}
	_f(0, -1)

	var f func(v, fa int) []int
	f = func(v, fa int) []int {
		if hson[v] == -1 {
			return []int{1}
		}
		cnt := f(hson[v], v)
		for _, w := range g[v] {
			if w != fa && w != hson[v] {
				subCnt := f(w, v)
				// 注意要先计算再合并，否则会把合并的结果额外算到答案上
				for i, c := range subCnt {
					if j := len(cnt) - k + len(subCnt) - i; 0 <= j && j < len(cnt) {
						ans += cnt[j] * c
					}
				}
				shift := len(cnt) - len(subCnt)
				for i, c := range subCnt {
					cnt[i+shift] += c
				}
			}
		}
		cnt = append(cnt, 1)
		if len(cnt) > k {
			ans += cnt[0]
			cnt = cnt[1:]
		}
		return cnt
	}
	f(0, -1)
	Fprint(out, ans)
}

//func main() { CF161D(os.Stdin, os.Stdout) }
