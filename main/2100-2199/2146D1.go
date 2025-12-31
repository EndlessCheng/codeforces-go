package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2146D1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &n)
		n++
		a := make([]int, n)
		cur := 0
		for n > 0 {
			hb := 1 << (bits.Len(uint(n)) - 1)
			left := n - hb
			for i := hb - 1; i >= left; i-- {
				a[i] = cur
				cur++
			}
			for i := n - 1; i >= hb; i-- {
				a[i] = cur
				cur++
			}
			n = left
		}

		or := 0
		for i, v := range a {
			or += v | i
		}
		Fprintln(out, or)
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2146D1(bufio.NewReader(os.Stdin), os.Stdout) }
