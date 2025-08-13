package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1954B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		p := []int{}
		for i := range a {
			Fscan(in, &a[i])
			if a[i] != a[0] {
				p = append(p, i)
			}
		}
		if len(p) > 0 {
			Fprintln(out, min(p[0], n-1-p[len(p)-1]))
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { cf1954B(bufio.NewReader(os.Stdin), os.Stdout) }
