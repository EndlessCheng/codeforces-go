package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, maxD, q, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &maxD, &q)

	mx := bits.Len(uint(n))
	pa := make([][]int, n)
	l := 0
	for i, x := range a {
		for x-a[l] > maxD {
			l++
		}
		pa[i] = make([]int, mx)
		pa[i][0] = l
	}
	for i := 0; i < mx-1; i++ {
		for x := range pa {
			p := pa[x][i]
			pa[x][i+1] = pa[p][i]
		}
	}

	for ; q > 0; q-- {
		Fscan(in, &v, &w)
		v--
		w--
		if v > w {
			v, w = w, v
		}
		res := 0
		for k := mx - 1; k >= 0; k-- {
			p := pa[w][k]
			if p > v {
				res |= 1 << k
				w = p
			}
		}
		Fprintln(out, res+1)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
