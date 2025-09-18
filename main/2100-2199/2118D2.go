package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2118D2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, q, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		x := make([]int, n)
		for i := range x {
			Fscan(in, &x[i])
		}
		d := make([]int, n)
		to := make([]int, n*2)
		for i := range to {
			to[i] = -1
		}
		idx := map[int][]int{}
		last := map[int]int{}
		for i := range d {
			Fscan(in, &d[i])
			v := (x[i]%k - d[i] + k) % k
			if j, ok := last[v]; ok {
				to[n+j] = i
			}
			last[v] = i
			idx[v] = append(idx[v], i)
		}

		last = map[int]int{}
		for i := n - 1; i >= 0; i-- {
			v := (x[i] + d[i]) % k
			if j, ok := last[v]; ok {
				to[j] = n + i
			}
			last[v] = i
		}

		vis := make([]int8, n*2)
		var dfs func(int) bool
		dfs = func(v int) bool {
			if vis[v] != 0 {
				return vis[v] > 0
			}
			vis[v] = -1
			if w := to[v]; w < 0 || dfs(w) {
				vis[v] = 1
				return true
			}
			return false
		}

		Fscan(in, &q)
		for range q {
			Fscan(in, &v)
			id := idx[v%k]
			i := sort.Search(len(id), func(i int) bool { return x[id[i]] >= v })
			if i == len(id) || dfs(id[i]) {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}
}

//func main() { cf2118D2(bufio.NewReader(os.Stdin), os.Stdout) }
