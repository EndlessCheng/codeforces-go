package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2131E(in io.Reader, out io.Writer) {
	var T, n int
o:
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

		if a[n-1] != b[n-1] {
			Fprintln(out, "NO")
			continue
		}
		for i := n - 2; i >= 0; i-- {
			if a[i] != b[i] && a[i]^a[i+1] != b[i] && a[i]^b[i+1] != b[i] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf2131E(bufio.NewReader(os.Stdin), os.Stdout) }
