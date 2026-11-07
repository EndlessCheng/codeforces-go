package main

import (
	. "fmt"
	"io"
	. "strings"
)

// https://github.com/EndlessCheng
func cf2250B(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		if k == n-1 {
			Fprintln(out, -1)
			continue
		}
		s := Repeat("0", (k+1)/2+1) + Repeat("1", k/2+1)
		if k < n-2 {
			s = Repeat("1", k%2) + s + Repeat("01", (n-2-k-k%2)/2) + Repeat("0", n%2)
		}
		Fprintln(out, s)
	}
}

//func main() { cf2250B(bufio.NewReader(os.Stdin), os.Stdout) }
