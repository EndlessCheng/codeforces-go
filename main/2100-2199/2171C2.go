package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2171C2(in io.Reader, out io.Writer) {
	name := [2]string{"Ajisai", "Mai"}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			s ^= a[i]
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
			s ^= b[i]
		}

		if s == 0 {
			Fprintln(out, "Tie")
			continue
		}

		hb := 1 << (bits.Len(uint(s)) - 1)
		for i := n - 1; ; i-- {
			if a[i]&hb != b[i]&hb {
				Fprintln(out, name[i%2])
				break
			}
		}
	}
}

//func main() { cf2171C2(bufio.NewReader(os.Stdin), os.Stdout) }
