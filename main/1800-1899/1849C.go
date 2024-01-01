package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1849C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, l, r int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s)
		l0 := make([]int, n+1)
		for i, b := range s {
			if b == '0' {
				l0[i+1] = i + 1
			} else {
				l0[i+1] = l0[i]
			}
		}
		
		r1 := make([]int, n+2)
		r1[n+1] = n + 1
		for i := n; i > 0; i-- {
			if s[i-1] == '1' {
				r1[i] = i
			} else {
				r1[i] = r1[i+1]
			}
		}
		
		set := map[[2]int]struct{}{}
		for ; m > 0; m-- {
			Fscan(in, &l, &r)
			l = r1[l]
			r = l0[r]
			if l > r {
				l, r = 0, 0
			}
			set[[2]int{l, r}] = struct{}{}
		}
		Fprintln(out, len(set))
	}
}

//func main() { cf1849C(os.Stdin, os.Stdout) }
