package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2160C(in io.Reader, out io.Writer) {
	var T, n uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n > 0 {
			n /= n & -n
		}
		m := bits.Len(n)
		if m%2 > 0 && n>>(m/2)&1 > 0 || n != bits.Reverse(n)>>bits.LeadingZeros(n) {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { cf2160C(bufio.NewReader(os.Stdin), os.Stdout) }
