package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2117E(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([][2]int, n)
		for i := range a {
			Fscan(in, &a[i][0])
		}
		for i := range a {
			Fscan(in, &a[i][1])
		}

		pos := make([][2]int, n+1)
		for i := n - 1; i >= 0; i-- {
			if a[i][0] == a[i][1] {
				Fprintln(out, i+1)
				continue o
			}
			for j, v := range a[i] {
				p := pos[v]
				if p[j] > 0 || p[j^1]-i > 1 {
					Fprintln(out, i+1)
					continue o
				}
				if p[j] == 0 {
					pos[v][j] = i
				}
			}
		}
		Fprintln(out, 0)
	}
}

//func main() { cf2117E(bufio.NewReader(os.Stdin), os.Stdout) }
