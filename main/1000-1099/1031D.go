package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1031D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	mat := make([]string, n)
	for i := range mat {
		Fscan(in, &mat[i])
	}

	ans := make([]byte, 2*n-1)
	checks := map[[2]int]int{{0, 0}: k}
	for i := range ans {
		minC := byte('z')
		for p, k := range checks {
			if k > 0 {
				minC = 'a'
				break
			}
			if c := mat[p[0]][p[1]]; c < minC {
				minC = c
			}
		}
		newChecks := map[[2]int]int{}
		for p, k := range checks {
			x, y := p[0], p[1]
			c := mat[x][y]
			if k <= 0 && c > minC {
				continue
			}
			if c > minC {
				k--
			}
			if x+1 < n {
				if kk, ok := newChecks[[2]int{x + 1, y}]; !ok || k > kk {
					newChecks[[2]int{x + 1, y}] = k
				}
			}
			if y+1 < n {
				if kk, ok := newChecks[[2]int{x, y + 1}]; !ok || k > kk {
					newChecks[[2]int{x, y + 1}] = k
				}
			}
		}
		checks = newChecks
		ans[i] = minC
	}
	Fprint(out, string(ans))
}

//func main() {
//	CF1031D(os.Stdin, os.Stdout)
//}
