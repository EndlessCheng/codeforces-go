package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func cf2032E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		m := n / 2
		p := make([]int, n*10+1)
		for i := range n {
			Fscan(in, &v)
			p[i+3] -= v
			p[i+3+m*2] += v * 2
			p[i+3+m*4] -= v
		}

		for i := 2; i <= n*10; i++ {
			p[i] += p[i-2]
		}
		for i := 2; i <= n*10; i++ {
			p[i] += p[i-2]
		}

		sum := make([]int, n)
		for i, v := range p {
			sum[i%n] += v
		}
		minS := slices.Min(sum)
		for _, s := range sum {
			Fprint(out, s-minS, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2032E(bufio.NewReader(os.Stdin), os.Stdout) }
