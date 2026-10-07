package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2224B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		has := make([]bool, n+1)
		mx := 0
		for range n {
			Fscan(in, &v)
			has[min(v, n)] = true
			mx = max(mx, v)
		}

		mex := 0
		for has[mex] {
			mex++
		}

		if mex < mx {
			Fprintln(out, mx*n+mex*(mex-1)/2+(n-mex)*mex)
		} else {
			Fprintln(out, mx*n+mx*(mx-1)/2+(n-mex+1)*mex)
		}
	}
}

//func main() { cf2224B(bufio.NewReader(os.Stdin), os.Stdout) }
