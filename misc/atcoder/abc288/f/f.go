package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var s string
	Fscan(bufio.NewReader(in), &s, &s)
	f, sumF := 0, 1
	for _, c := range s {
		f = (f*10 + sumF*int(c-'0')) % mod
		sumF = (sumF + f) % mod
	}
	Fprint(out, f)
}

func main() { run(os.Stdin, os.Stdout) }
