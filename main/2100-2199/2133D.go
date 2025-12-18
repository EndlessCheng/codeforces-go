package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2133D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		f := make([]int, n+1)
		f[1] = a[0]
		for i := 1; i < n; i++ {
			f[i+1] = min(f[i]+a[i]-1, f[i-1]+a[i-1]+max(a[i]-i, 0))
		}
		Fprintln(out, f[n])
	}
}

//func main() { cf2133D(bufio.NewReader(os.Stdin), os.Stdout) }
