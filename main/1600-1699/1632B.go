package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1632B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		hb := 1 << (bits.Len(uint(n-1)) - 1)
		for i := 1; i < hb; i++ {
			Fprint(out, i, " ")
		}
		Fprint(out, "0 ", hb)
		for i := hb + 1; i < n; i++ {
			Fprint(out, " ", i)
		}
		Fprintln(out)
	}
}

//func main() { cf1632B(bufio.NewReader(os.Stdin), os.Stdout) }
