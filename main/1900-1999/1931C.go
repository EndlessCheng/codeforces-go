package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1931C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		l := 0
		for i := range a {
			Fscan(in, &a[i])
			if l == 0 && a[i] != a[0] {
				l = i
			}
		}
		r := -1
		for i, v := range a {
			if v != a[n-1] {
				r = i
			}
		}
		if a[0] == a[n-1] {
			Fprintln(out, r-l+1)
		} else {
			Fprintln(out, min(n-l, r+1))
		}
	}
}

//func main() { cf1931C(bufio.NewReader(os.Stdin), os.Stdout) }
