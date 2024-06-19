package main

import (
	. "fmt"
	"io"
	"math/bits"
)

func cf1950F(in io.Reader, out io.Writer) {
	var T, a, b, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		if a+1 != c {
			Fprintln(out, -1)
			continue
		}
		h := bits.Len(uint(a))
		b -= 1<<h - c
		if b > 0 {
			h += (b-1)/c + 1
		}
		Fprintln(out, h)
	}
}

//func main() { cf1950F(bufio.NewReader(os.Stdin), os.Stdout) }
