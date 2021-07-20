package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF791D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, w int
	Fscan(in, &n, &k)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	sz := make([][5]int, n)
	sum := make([][5]int64, n)
	var f func(int, int)
	f = func(v, fa int) {
		sz[v][0] = 1
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
				for i := 0; i < k; i++ {
					sum[v][i] += sum[w][(i+k-1)%k] + int64(sz[w][(i+k-1)%k])
					sz[v][i] += sz[w][(i+k-1)%k]
				}
			}
		}
		return
	}
	f(0, -1)

	ans := int64(0)
	var f2 func(int, int)
	f2 = func(v, fa int) {
		ans += sum[v][0] / int64(k)
		for i := 1; i < k; i++ {
			ans += int64(sz[v][i]) + (sum[v][i]-int64(sz[v][i])*int64(i))/int64(k)
		}
		for _, w := range g[v] {
			if w != fa {
				s := sum[v]
				z := sz[v]
				for i := 0; i < k; i++ {
					s[i] -= sum[w][(i+k-1)%k] + int64(sz[w][(i+k-1)%k])
					z[i] -= sz[w][(i+k-1)%k]
				}
				for i := 0; i < k; i++ {
					sum[w][i] += s[(i+k-1)%k] + int64(z[(i+k-1)%k])
					sz[w][i] += z[(i+k-1)%k]
				}
				f2(w, v)
			}
		}
	}
	f2(0, -1)
	Fprint(out, ans/2)
}

//func main() { CF791D(os.Stdin, os.Stdout) }
