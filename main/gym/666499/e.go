package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cfE(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, q, l, r int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &q, &s)
		for range q {
			Fscan(in, &l, &r)
			t := s[l-1 : r]
			m := len(t)

			z := make([]int, m)
			z[0] = m
			boxL, boxR := 0, 0

			type pair struct{ l, r int }
			st := []pair{{0, m - 1}}
			f := make([]int, m+1)
			f[1] = 1
			ans := 1

			for i := 1; i < m; i++ {
				if i <= boxR {
					z[i] = min(z[i-boxL], boxR-i+1)
				}
				for i+z[i] < m && t[z[i]] == t[i+z[i]] {
					boxL, boxR = i, i+z[i]
					z[i]++
				}
				for st[len(st)-1].r < i {
					st = st[:len(st)-1]
				}
				if z[i] > 0 {
					st = append(st, pair{i, i + z[i] - 1})
				}
				f[i+1] = f[st[len(st)-1].l] + 1
				ans += f[i+1]
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { cfE(bufio.NewReader(os.Stdin), os.Stdout) }
