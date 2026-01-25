package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2035C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%2 > 0 {
			Fprintln(out, n)
			for i := 4; i < n-1; i++ {
				Fprint(out, i, " ")
			}
			Fprintln(out, 2, 1, 3, n-1, n)
		} else {
			w := bits.Len(uint(n))
			Fprintln(out, 1<<w-1)
			hb := 1 << (w - 1)
			for i := 2; i <= n; i++ {
				if i != 5 && (i < hb-2 || i > hb) {
					Fprint(out, i, " ")
				}
			}
			Fprintln(out, 1, 5, hb-2, hb-1, hb)
		}
	}
}

//func main() { cf2035C(bufio.NewReader(os.Stdin), os.Stdout) }
