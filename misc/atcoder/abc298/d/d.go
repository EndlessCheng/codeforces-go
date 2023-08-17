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
	const inv10 = 299473306

	var q, op, x int
	cur, pow10 := 1, 1
	s := []int{1}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &x)
			cur = (cur*10 + x) % mod
			pow10 = pow10 * 10 % mod
			s = append(s, x)
		} else if op == 2 {
			cur = (cur - s[0]*pow10) % mod
			pow10 = pow10 * inv10 % mod
			s = s[1:]
		} else {
			Fprintln(out, (cur+mod)%mod)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
