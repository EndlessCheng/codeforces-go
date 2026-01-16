package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1208F(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	const w = 21
	cnt := [1 << w]int{}
	var upd func(int, int)
	upd = func(s, t int) {
		if cnt[s|t] == 2 {
			return
		}
		if s == 0 {
			cnt[t]++
			return
		}
		lb := s & -s
		upd(s^lb, t|lb)
		upd(s^lb, t)
	}

	upd(a[n-1], 0)
	upd(a[n-2], 0)
	for i := n - 3; i >= 0; i-- {
		and := 0
		for j := w - 1; j >= 0; j-- {
			b := 1 << j
			if a[i]&b == 0 && cnt[and|b] == 2 {
				and |= b
			}
		}
		ans = max(ans, a[i]|and)
		upd(a[i], 0)
	}
	Fprint(out, ans)
}

//func main() { cf1208F(bufio.NewReader(os.Stdin), os.Stdout) }
