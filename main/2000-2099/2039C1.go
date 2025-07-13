package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2039C1(in io.Reader, out io.Writer) {
	var T, x, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &m)
		ans := 0
		l := bits.Len(uint(x))
		for y := 1; y <= min(1<<l-1, m); y++ {
			if y != x && (x%(x^y) == 0 || y%(x^y) == 0) {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2039C1(bufio.NewReader(os.Stdin), os.Stdout) }
