package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2125C(in io.Reader, out io.Writer) {
	ps := []int{2, 3, 5, 7}
	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		ans := 0
		for sub := range 1 << len(ps) {
			mul := 1
			for s := uint(sub); s > 0; s &= s - 1 {
				mul *= ps[bits.TrailingZeros(s)]
			}
			sign := 1 - bits.OnesCount(uint(sub))%2*2
			ans += (r/mul - (l-1)/mul) * sign
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2125C(bufio.NewReader(os.Stdin), os.Stdout) }
