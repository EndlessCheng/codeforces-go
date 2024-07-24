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
		l := make([]int, n+1)
		for i := 1; i <= n; i++ {
			r := 0
			for j := n; j >= i; j-- {
				xor := sum[j] ^ sum[i-1]
				ok := i == 1 && j == n || l[j] < 0 || r < 0 || l[j]&xor != 0 || r&xor != 0
				if ok {
					if xor == 0 {
						l[j] = -1
						r = -1
					} else {
						high := 1 << (bits.Len(uint(xor)) - 1)
						l[j] |= high
						r |= high
					}
				}
				if j == i {
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
