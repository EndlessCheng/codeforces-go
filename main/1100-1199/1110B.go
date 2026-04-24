package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1110B(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	ans := b[n-1] - b[0] + 1
	for i := n - 1; i > 0; i-- {
		b[i] -= b[i-1]
	}
	slices.Sort(b[1:])
	for _, v := range b[n-k+1:] {
		ans -= v - 1
	}
	Fprint(out, ans)
}

//func main() { cf1110B(bufio.NewReader(os.Stdin), os.Stdout) }
