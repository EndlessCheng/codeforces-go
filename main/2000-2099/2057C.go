package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2057C(in io.Reader, out io.Writer) {
	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		hb := 1 << (bits.Len(uint(l^r)) - 1)
		v := l&r&^(hb-1) | hb
		if v == r {
			v--
		}
		Fprintln(out, v-1, v, v+1)
	}
}

//func main() { cf2057C(bufio.NewReader(os.Stdin), os.Stdout) }
