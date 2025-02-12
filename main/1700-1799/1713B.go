package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1713B(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		a = slices.Compact(a)
		for i := 1; i < len(a)-1; i++ {
			if a[i-1] > a[i] && a[i] < a[i+1] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1713B(bufio.NewReader(os.Stdin), os.Stdout) }
