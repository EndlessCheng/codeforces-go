package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1555D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	perm3 := []string{"abc", "acb", "bac", "bca", "cab", "cba"}

	var n, q, l, r int
	var s []byte
	Fscan(in, &n, &q, &s)
	sum := [6][]int{}
	for i, p := range perm3 {
		sum[i] = make([]int, n+1)
		for j, b := range s {
			sum[i][j+1] = sum[i][j]
			if b != p[j%3] {
				sum[i][j+1]++
			}
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		ans := n
		for _, s := range sum {
			ans = min(ans, s[r]-s[l-1])
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1555D(os.Stdin, os.Stdout) }
