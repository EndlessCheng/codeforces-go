package main

import (
	. "fmt"
	"io"
	"math/bits"
)

func cf1934D1(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		if n^m < n {
			Fprintln(out, 1, n, m)
			continue
		}
		hb := 1 << (bits.Len(uint(n)) - 1)
		x := 1<<bits.Len(uint(n^hb)) - 1
		if x < m {
			Fprintln(out, -1)
		} else {
			Fprintln(out, 2, n, x, m)
		}
	}
}

//func main() { cf1934D1(bufio.NewReader(os.Stdin), os.Stdout) }
