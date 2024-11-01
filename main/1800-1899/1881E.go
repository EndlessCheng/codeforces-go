package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1881E(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f := make([]int, n+2)
		f[n+1] = 1e9
		for i := n - 1; i >= 0; i-- {
			f[i] = min(f[i+1]+1, f[min(i+a[i]+1, n+1)])
		}
		Fprintln(out, f[0])
	}
}

//func main() { cf1881E(bufio.NewReader(os.Stdin), os.Stdout) }
