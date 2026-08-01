package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2241D(in io.Reader, out io.Writer) {
	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		free := 0
		for _, v := range a {
			Fscan(in, &w)
			free += w - v
			if free < 0 {
				free = -1e18
			}
		}
		if free < 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { cf2241D(bufio.NewReader(os.Stdin), os.Stdout) }
