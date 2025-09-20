package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1550B(in io.Reader, out io.Writer) {
	var T, n, a, b int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b, &s)
		if b >= 0 {
			Fprintln(out, n*(a+b))
			continue
		}
		c := 1
		for i := 1; i < n; i++ {
			c += int(s[i-1] ^ s[i])
		}
		Fprintln(out, n*a+(c/2+1)*b)
	}
}

//func main() { cf1550B(bufio.NewReader(os.Stdin), os.Stdout) }
