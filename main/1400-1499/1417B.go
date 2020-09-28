package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1417B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, s, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := make([]interface{}, n)
		ps := []int{}
		for i := range ans {
			if Fscan(in, &v); v <= s/2 {
				ans[i] = 0
				if v == s/2 {
					ps = append(ps, i)
				}
			} else {
				ans[i] = 1
			}
		}
		if s&1 == 0 {
			for _, p := range ps[:len(ps)/2] {
				ans[p] = 1
			}
		}
		Fprintln(out, ans...)
	}
}

//func main() { CF1417B(os.Stdin, os.Stdout) }
