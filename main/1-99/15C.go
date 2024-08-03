package main

import (
	. "fmt"
	"io"
)

func cf15C(in io.Reader, out io.Writer) {
	xor := func(n int) int {
		switch n % 4 {
		case 0: return n
		case 1: return 1
		case 2: return n + 1
		default: return 0
		}
	}
	var n, l, m, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &l, &m)
		ans ^= xor(l+m-1) ^ xor(l-1)
	}
	if ans > 0 {
		Fprintln(out, "tolik")
	} else {
		Fprintln(out, "bolik")
	}
}

//func main() { cf15C(bufio.NewReader(os.Stdin), os.Stdout) }
