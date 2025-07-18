package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2065C1(in io.Reader, out io.Writer) {
	var T, n, b int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &b)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		Fscan(in, &b)

		pre := int(-1e9)
		for _, v := range a {
			if b-v >= pre {
				mn := b - v
				if v >= pre {
					mn = min(mn, v)
				}
				v = mn
			} else if v < pre {
				Fprintln(out, "NO")
				continue o
			}
			pre = v
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf2065C1(bufio.NewReader(os.Stdin), os.Stdout) }
