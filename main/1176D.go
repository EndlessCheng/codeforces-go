package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1176D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx = 2750132
	id := [mx]int{-1, -1}
	f := [mx]int{1: 1}
	c := 1
	for i := 2; i < mx; i++ {
		if id[i] == 0 {
			id[i] = c
			c++
			for j := 2 * i; j < mx; j += i {
				id[j] = -1
				if f[j] == 0 {
					f[j] = i
				}
			}
		}
	}

	cnt := [mx]int{}
	var n, v int
	Fscan(in, &n)
	for i := 0; i < 2*n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	ans := make([]interface{}, 0, n)
	for i := mx - 1; i > 0; i-- {
		if c := cnt[i]; c > 0 {
			v := id[i]
			if v > 0 {
				cnt[v] -= c
			} else {
				cnt[i/f[i]] -= c
				v = i
			}
			for ; c > 0; c-- {
				ans = append(ans, v)
			}
		}
	}
	Fprint(out, ans...)
}

//func main() { CF1176D(os.Stdin, os.Stdout) }
