package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1282B2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, p, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &p, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		s := make([]int, n+1)
		for i, v := range a {
			s[i+1] = s[i] + v
		}
		ans := 0
		for i := 0; i < k; i++ {
			p, c := p, 0
			for j := i + k - 1; j < n && a[j] <= p; j += k {
				p -= a[j]
				c += k
			}
			c += sort.SearchInts(s[1:i+1], p+1)
			if c > ans {
				ans = c
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1282B2(os.Stdin, os.Stdout) }
