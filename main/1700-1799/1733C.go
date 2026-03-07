package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1733C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		Fprintln(out, n-1)
		idx := -1
		for i := n - 1; i >= 0; i-- {
			if a[i]%2 == a[0]%2 {
				if idx < 0 {
					idx = i
				} else {
					Fprintln(out, i+1, idx+1)
				}
			}
		}
		for i, v := range a {
			if v%2 != a[0]%2 {
				Fprintln(out, 1, i+1)
			}
		}
	}
}

//func main() { cf1733C(bufio.NewReader(os.Stdin), os.Stdout) }
