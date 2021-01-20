package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1474C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, m, v int
O:
	for Fscan(in, &T); T > 0; T-- {
		cnt := map[int]int{}
		Fscan(in, &m)
		for i := 0; i < 2*m; i++ {
			Fscan(in, &v)
			cnt[v]++
		}
		n := len(cnt)
		a := make([]int, 0, n)
		for k := range cnt {
			a = append(a, k)
		}
		sort.Ints(a)
		X := a[n-1]
		cnt[X]--
	o:
		for _, v := range a {
			if cnt[v] == 0 {
				continue
			}
			c := make(map[int]int, n)
			for k, v := range cnt {
				c[k] = v
			}
			ans := append(make([]interface{}, 0, 2*m), X, v)
			c[v]--
			x := X
			for j := n - 1; j >= 0; j-- {
				v := a[j]
				if c[v] == 0 {
					continue
				}
				c[v]--
				if c[x-v] == 0 {
					continue o
				}
				c[x-v]--
				if c[v] > 0 {
					continue o
				}
				ans = append(ans, v, x-v)
				x = v
			}
			Fprintln(out, "YES")
			Fprintln(out, X+v)
			Fprintln(out, ans...)
			continue O
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1474C(os.Stdin, os.Stdout) }
