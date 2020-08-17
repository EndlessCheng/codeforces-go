package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1332B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	ps := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

	var t, n, v int
	for Fscan(in, &t); t > 0; t-- {
		c := map[int]int{}
		Fscan(in, &n)
		ans := make([]interface{}, n)
		for i := range ans {
			Fscan(in, &v)
			for _, p := range ps {
				if v%p == 0 {
					if c[p] == 0 {
						c[p] = len(c) + 1
					}
					ans[i] = c[p]
					break
				}
			}
		}
		Fprintln(out, len(c))
		Fprintln(out, ans...)
	}
}

//func main() { CF1332B(os.Stdin, os.Stdout) }
