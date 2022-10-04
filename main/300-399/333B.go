package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF333B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, r, c, ans int
	Fscan(in, &n, &m)
	ban := make([]bool, n*2)
	ban[0] = true
	ban[n-1] = true
	ban[n] = true
	ban[n*2-1] = true
	for ; m > 0; m-- {
		Fscan(in, &r, &c)
		ban[r-1] = true
		ban[n+c-1] = true
	}
	if n&1 > 0 && !ban[n/2] && !ban[n+n/2] {
		ban[n/2] = true
	}
	for _, b := range ban {
		if !b {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF333B(os.Stdin, os.Stdout) }
