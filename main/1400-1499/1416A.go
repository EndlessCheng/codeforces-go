package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1416A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ps := make([][]int, n+1)
		for i := range ps {
			ps[i] = []int{-1}
		}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			ps[v] = append(ps[v], i)
		}
		for i := range ps {
			ps[i] = append(ps[i], n)
		}
		ans := make([]int, n+2)
		for i := range ans {
			ans[i] = 1e9
		}
		for i := n; i > 0; i-- {
			mx := 0
			for j := 1; j < len(ps[i]); j++ {
				if d := ps[i][j] - ps[i][j-1]; d > mx {
					mx = d
				}
			}
			ans[mx] = i
		}
		for i := 1; i <= n; i++ {
			if ans[i-1] < ans[i] {
				ans[i] = ans[i-1]
			}
			if ans[i] == 1e9 {
				Fprint(out, "-1 ")
			} else {
				Fprint(out, ans[i], " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1416A(os.Stdin, os.Stdout) }
