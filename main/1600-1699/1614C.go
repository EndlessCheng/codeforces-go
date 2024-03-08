package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1614C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007

	var T, n, m, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := 0
		for ; m > 0; m-- {
			Fscan(in, &x, &x, &x)
			ans |= x
		}
		ans %= mod
		for ; n > 1; n-- {
			ans = ans * 2 % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1614C(os.Stdin, os.Stdout) }
