package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1468F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n, x, y, u, v int
	for Fscan(in, &T); T > 0; T-- {
		ans := int64(0)
		type pair struct{ x, y int }
		cnt := map[pair]int{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &x, &y, &u, &v)
			u -= x
			v -= y
			g := gcd(abs(u), abs(v))
			u /= g
			v /= g
			ans += int64(cnt[pair{-u, -v}])
			cnt[pair{u, v}]++
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1468F(os.Stdin, os.Stdout) }
