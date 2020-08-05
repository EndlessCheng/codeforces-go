package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF633D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	cnt := map[int]int{}
	var n, v, ans int
	Fscan(in, &n)
	a := []int{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
		if v != 0 {
			a = append(a, v)
		}
	}
	if c := cnt[0]; c > 0 {
		ans = c
		a = append(a, 0)
	}
	for i, v := range a {
		for _, w := range a[i+1:] {
			c := 2
			mp := map[int]int{v: 1}
			mp[w]++
			for x, y := v, w; mp[x+y] < cnt[x+y]; x, y = y, x+y {
				mp[x+y]++
				c++
			}
			for x, y := w, v; mp[x-y] < cnt[x-y]; x, y = y, x-y {
				mp[x-y]++
				c++
			}
			if c > ans {
				ans = c
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF633D(os.Stdin, os.Stdout) }
