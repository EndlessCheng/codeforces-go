package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF988D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] += 1e9
	}
	sort.Ints(a)
	var ans []interface{}
o:
	for i := 0; i < 31; i++ {
		d := 1 << i
		g := map[int][]int{}
		for _, x := range a {
			g[x%d] = append(g[x%d], x)
		}
		for _, xs := range g {
			for j := 0; j < len(xs); {
				b := []interface{}{}
				for j0 := j; j < len(xs) && xs[j] == xs[j0]+(j-j0)*d; j++ {
					b = append(b, xs[j]-1e9)
					if len(b) == 3 {
						ans = b
						break o
					}
				}
				if len(b) > len(ans) {
					ans = b
				}
			}
		}
	}
	Fprintln(out, len(ans))
	Fprint(out, ans...)
}

//func main() { CF988D(os.Stdin, os.Stdout) }
