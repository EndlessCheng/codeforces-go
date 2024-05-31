package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1977C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		has := map[int]bool{}
		for i := range a {
			Fscan(in, &a[i])
			has[a[i]] = true
		}
		l := 1
		for _, v := range a {
			l = lcm(l, v)
			if l > 1e9 {
				Fprintln(out, n)
				continue o
			}
		}
		ans := 0
		f := func(d int) {
			l, cnt := 1, 0
			for _, v := range a {
				if d%v == 0 {
					l = lcm(l, v)
					cnt++
				}
			}
			if !has[l] {
				ans = max(ans, cnt)
			}
		}
		for d := 1; d*d <= l; d++ {
			if l%d == 0 {
				f(d)
				if d*d < l {
					f(l / d)
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1977C(bufio.NewReader(os.Stdin), os.Stdout) }
