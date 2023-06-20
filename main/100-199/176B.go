package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF176B(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var s, t string
	var k, c, f int64
	Fscan(bufio.NewReader(in), &s, &t, &k)
	n := len(s)
	ss := s + s
	for i := 0; i < n; i++ {
		if ss[i:i+n] == t {
			c++
		}
	}
	if s == t {
		f = 1
	}
	g := f ^ 1
	for ; k > 0; k-- {
		f, g = (f*(c-1)+g*c)%mod,
			   (f*(int64(n)-c)+g*(int64(n)-c-1))%mod
	}
	Fprint(out, f)
}

//func main() { CF176B(os.Stdin, os.Stdout) }
