package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1721C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		for i, j, max := n-1, n-1, b[n-1]; i >= 0; i-- {
			for ; j >= 0 && b[j] >= a[i]; j-- {
			}
			min := b[j+1]
			b[i] = max - a[i]
			a[i] = min - a[i]
			if j+1 == i && j >= 0 {
				max = b[j]
			}
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
		for _, v := range b {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1721C(os.Stdin, os.Stdout) }
