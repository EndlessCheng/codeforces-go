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
		if n == 1 {
			continue
		}

		Fprintln(out, 1, n)
		x := a[0]
		if (a[0]+a[n-1])%2 == 0 {
			x = a[n-1]
		}
		for i := 1; i < n-1; i++ {
			if (a[i]+x)%2 == 0 {
				Fprintln(out, i+1, n)
			} else {
				Fprintln(out, 1, i+1)
			}
		}
	}
}

//func main() { cf1733C(bufio.NewReader(os.Stdin), os.Stdout) }
