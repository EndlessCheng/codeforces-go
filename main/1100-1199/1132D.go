package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1132D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]struct{ pow, dec int64 }, n)
	for i := range a {
		Fscan(in, &a[i].pow)
	}
	for i := range a {
		Fscan(in, &a[i].dec)
	}

	l, r := int64(0), int64(3e12)
o:
	for l < r {
		mid := (l + r) / 2
		cnt := make([]int, k)
		c := 0
		for _, p := range a {
			pow := p.pow
			for t := pow/p.dec + 1; t < int64(k); t = pow/p.dec + 1 {
				if c++; c == k {
					l = mid + 1
					continue o
				}
				cnt[t]++
				pow += mid
			}
		}
		for i := 1; i < k; i++ {
			cnt[i] += cnt[i-1]
			if cnt[i] > i {
				l = mid + 1
				continue o
			}
		}
		r = mid
	}
	if l == 3e12 {
		l = -1
	}
	Fprint(out, l)
}

//func main() { CF1132D(os.Stdin, os.Stdout) }
