package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF476C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var a, b int64
	Fscan(in, &a, &b)
	Fprint(out, (b*(b-1)/2%mod)*(a*(a+1)/2%mod*b%mod+a)%mod)
}

//func main() { CF476C(os.Stdin, os.Stdout) }
