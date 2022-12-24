package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	const mod2 = (mod + 1) / 2

	var m, q, r1, r2, c1, c2 int
	Fscan(in, &m, &m, &q)
	f := func(x, y int) int {
		s := (y + y%2 + (x-1)&^1*m) % mod * ((y + 1) / 2) % mod * ((x + 1) / 2)
		s2 := (y&^1 + 2 + x&^1*m) % mod * (y / 2) % mod * (x / 2)
		return s + s2
	}
	for ; q > 0; q-- {
		Fscan(in, &r1, &r2, &c1, &c2)
		res := f(r2, c2) - f(r2, c1-1) - f(r1-1, c2) + f(r1-1, c1-1)
		Fprintln(out, (res%mod+mod)*mod2%mod)
	}
}

func main() { run(os.Stdin, os.Stdout) }
