package main

import (
	. "fmt"
	"io"
	"math/bits"
)

func cf1863F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		sum := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &sum[i])
			sum[i] ^= sum[i-1]
		}
		leftBits := make([]int, n)
		for l := 0; l < n; l++ {
			rightBits := 0
			for r := n - 1; r >= l; r-- {
				xor := sum[r+1] ^ sum[l]
				ok := l == 0 && r == n-1 || leftBits[r] < 0 || rightBits < 0 || leftBits[r]&xor != 0 || rightBits&xor != 0
				if ok {
					if xor == 0 {
						leftBits[r] = -1
						rightBits = -1
					} else {
						high := 1 << (bits.Len(uint(xor)) - 1)
						leftBits[r] |= high
						rightBits |= high
					}
				}
				if r == l {
					if ok {
						Fprint(out, "1")
					} else {
						Fprint(out, "0")
					}
				}
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1863F(bufio.NewReader(os.Stdin), os.Stdout) }
