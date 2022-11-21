package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, l, r int
	Fscan(in, &n, &m, &q)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for ; m > 0; m-- {
		Fscan(in, &l, &r)
		f[l-1][r-1]++
	}
	for l := n - 1; l >= 0; l-- {
		for r := l + 1; r < n; r++ {
			f[l][r] += f[l+1][r] + f[l][r-1] - f[l+1][r-1]
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		Fprintln(out, f[l-1][r-1])
	}
}

func main() { run(os.Stdin, os.Stdout) }
