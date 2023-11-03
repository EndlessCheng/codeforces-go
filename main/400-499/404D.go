package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF404D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var s string
	Fscan(bufio.NewReader(in), &s)
	f0, f1, f2 := 1, 0, 1
	for _, b := range s {
		switch b {
		case '0': f1, f2 = 0, 0
		case '1': f0, f1, f2 = f1, 0, f0
		case '2': f0, f1, f2 = 0, 0, f1
		case '*': f0, f1 = 0, f2
		default:  f0, f1, f2 = (f0+f1)%mod, f2, (f0+f1+f2)%mod
		}
	}
	Fprint(out, (f0+f1)%mod)
}

//func main() { CF404D(os.Stdin, os.Stdout) }
