package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1628A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		r := make([]int, n+1)
		for i := range r {
			r[i] = -1
		}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			r[a[i]] = i
		}
		ans := []int{}
		vis := make([]int, n+2)
		i := 0
		for k := 1; i <= r[0]; k++ {
			mex := 0
			for ; i <= r[mex]; i++ {
				vis[a[i]] = k
				for vis[mex] == k {
					mex++
				}
			}
			ans = append(ans, mex)
		}
		Fprintln(out, len(ans)+n-i)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out, strings.Repeat("0 ", n-i))
	}
}

//func main() { CF1628A(os.Stdin, os.Stdout) }
