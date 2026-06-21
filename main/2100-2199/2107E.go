package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2107E(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		if k-2 >= n*(n-1)*(n-2)/6 {
			Fprintln(out, "NO")
			continue
		}
		Fprintln(out, "YES")
		x := 1
		for i := n - 1; i > 0; i-- {
			Fprintln(out, x, n-i+1)
			if k >= i*(i-1)/2 {
				k -= i * (i - 1) / 2
				x = n - i + 1
			}
		}
	}
}

//func main() { cf2107E(bufio.NewReader(os.Stdin), os.Stdout) }
