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
		} else {
			m := n - 2 - k
			Fprintln(out, Repeat("0", k/2+1)+Repeat("1", (k+1)/2+1)+Repeat("01", m/2)+Repeat("0", m%2))
		}
	}
}

//func main() { cf2250B(bufio.NewReader(os.Stdin), os.Stdout) }
