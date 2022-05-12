package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF878A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, vis, xor int
	var op string
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &op, &v)
		if op == "^" {
			xor ^= v
		} else if op == "&" {
			vis |= ^v
			xor &= v
		} else {
			vis |= v
			xor |= v
		}
	}
	Fprintln(out, 3)
	Fprintln(out, "^", xor)
	Fprintln(out, "&", 1023&^vis|xor)
	Fprintln(out, "|", vis&xor)
}

//func main() { CF878A(os.Stdin, os.Stdout) }
