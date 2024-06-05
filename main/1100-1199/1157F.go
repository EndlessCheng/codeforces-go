package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1157F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 2e5
	cnt := [mx + 2]int{}
	var n, v, ll, rr int
	ans := 1
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
		if cnt[v-1] > 0 {
			ans = 2
			ll, rr = v-1, v
		} else if cnt[v+1] > 0 {
			ans = 2
			ll, rr = v, v+1
		}
	}
	if ans == 1 {
		ll, rr = v, v
	}
	for i := 1; i <= mx; i++ {
		if cnt[i] < 2 {
			continue
		}
		l := i
		s := 0
		for ; i <= mx && cnt[i] > 1; i++ {
			s += cnt[i]
		}
		r := i - 1
		if cnt[l-1] > 0 {
			s++
			l--
		}
		if cnt[i] > 0 {
			s++
			r++
		}
		if s > ans {
			ans = s
			ll, rr = l, r
		}
		ans = max(ans, s)
	}
	Fprintln(out, ans)
	for i := ll; i <= rr; i++ {
		Fprint(out, i, " ")
	}
	for i := rr; i >= ll; i-- {
		for c := cnt[i]; c > 1; c-- {
			Fprint(out, i, " ")
		}
	}
}

//func main() { cf1157F(bufio.NewReader(os.Stdin), os.Stdout) }
