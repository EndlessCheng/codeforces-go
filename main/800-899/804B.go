package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF804B(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	s := ""
	Fscan(bufio.NewReader(in), &s)
	var ans, b int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'b' {
			b++
		} else {
			ans += b
			b = b * 2 % mod
		}
	}
	Fprint(out, ans%mod)
}

//func main() { CF804B(os.Stdin, os.Stdout) }
