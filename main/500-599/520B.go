package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF520B(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	for m > n {
		if m%2 == 0 {
			m /= 2
		} else {
			m++
		}
		ans++
	}
	Fprint(out, ans+n-m)
}

//func main() { CF520B(os.Stdin, os.Stdout) }
