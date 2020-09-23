package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1158C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pp := make([][]int, n+2)
		pos := make([]int, n)
		for i := range pos {
			Fscan(in, &pos[i])
			if pos[i] == -1 {
				pos[i] = i + 2
			}
			pp[pos[i]] = append(pp[pos[i]], i)
		}
		ans := make([]int, n)
		v := 1
		for _, ps := range pp {
			for i := len(ps) - 1; i >= 0; i-- {
				ans[ps[i]] = v
				v++
			}
		}
		type pair struct{ v, i int }
		s := []pair{{1e9, n}}
		for i := n - 1; i >= 0; i-- {
			v := ans[i]
			for {
				if top := s[len(s)-1]; top.v > v {
					if top.i+1 != pos[i] {
						Fprintln(out, -1)
						continue o
					}
					break
				}
				s = s[:len(s)-1]
			}
			s = append(s, pair{v, i})
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1158C(os.Stdin, os.Stdout) }
