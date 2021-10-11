package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1277D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct {
		s string
		i int
	}

	var T, n int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		a := [4][]pair{}
		has := map[string]bool{}
		Fscan(in, &n)
		for i := 1; i <= n; i++ {
			Fscan(in, &s)
			tp := s[0]&1<<1 | s[len(s)-1]&1
			a[tp] = append(a[tp], pair{s, i})
			if tp == 1 || tp == 2 {
				has[s] = true
			}
		}

		if len(has) == 0 {
			if a[0] != nil && a[3] != nil {
				Fprintln(out, -1)
			} else {
				Fprintln(out, 0)
			}
			continue
		}

		diff := len(a[1]) - len(a[2])
		if diff < 0 {
			diff, a[1], a[2] = -diff, a[2], a[1]
		}
		if diff < 2 {
			Fprintln(out, 0)
			continue
		}

		ans := []int{}
		for _, p := range a[1] {
			s := []byte(p.s)
			for i, n := 0, len(s); i < n/2; i++ {
				s[i], s[n-1-i] = s[n-1-i], s[i]
			}
			if !has[string(s)] {
				ans = append(ans, p.i)
				if diff -= 2; diff < 2 {
					Fprintln(out, len(ans))
					for _, v := range ans {
						Fprint(out, v, " ")
					}
					Fprintln(out)
					continue o
				}
			}
		}

		Fprintln(out, -1)
	}
}

//func main() { CF1277D(os.Stdin, os.Stdout) }
