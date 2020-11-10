package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF600E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	hson := make([]int, n)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		sz, hsz := 1, 0
		for _, w := range g[v] {
			if w != fa {
				s := f(w, v)
				sz += s
				if s > hsz {
					hsz = s
					hson[v] = w
				}
			}
		}
		return sz
	}
	f(0, -1)

	ans := make([]interface{}, n)
	var f2 func(v, fa int) (map[int]int, map[int]int64, int)
	f2 = func(v, fa int) (cnt map[int]int, val map[int]int64, maxCnt int) {
		if hson[v] == 0 {
			ans[v] = a[v]
			return map[int]int{a[v]: 1}, map[int]int64{1: int64(a[v])}, 1
		}
		cnt, val, maxCnt = f2(hson[v], v)
		merge := func(c, num int) {
			if cnt[c] > 0 {
				val[cnt[c]] -= int64(c)
			}
			cnt[c] += num
			val[cnt[c]] += int64(c)
			if cnt[c] > maxCnt {
				maxCnt = cnt[c]
			}
		}
		for _, w := range g[v] {
			if w != fa && w != hson[v] {
				sc, _, _ := f2(w, v)
				for c, num := range sc {
					merge(c, num)
				}
			}
		}
		merge(a[v], 1)
		ans[v] = val[maxCnt]
		return
	}
	f2(0, -1)
	Fprint(out, ans...)
}

//func main() { CF600E(os.Stdin, os.Stdout) }
