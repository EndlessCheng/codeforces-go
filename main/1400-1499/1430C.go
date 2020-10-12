package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1430C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			a[i] = i + 1
		}
		Fprintln(out, 2)
		for i := n - 1; i > 0; i-- {
			Fprintln(out, a[i], a[i-1])
			a[i-1] = (a[i] + a[i-1] + 1) / 2
		}
	}
}

//func main() { CF1430C(os.Stdin, os.Stdout) }
