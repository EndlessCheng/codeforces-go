package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1680C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		s1 := make([]int, n+1)
		for i, b := range s {
			s1[i+1] = s1[i] + int(b&1)
		}
		Fprintln(out, sort.Search(n, func(mx int) bool {
			c, l := 0, 0
			for r, b := range s {
				c += int(b&1 ^ 1)
				for c > mx {
					c -= int(s[l]&1 ^ 1)
					l++
				}
				if s1[n]-(s1[r+1]-s1[l]) <= mx {
					return true
				}
			}
			return false
		}))
	}
}

//func main() { CF1680C(os.Stdin, os.Stdout) }
