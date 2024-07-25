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
		for i := 0; i < n; i++ {
			rightBits := 0
			for j := n - 1; j >= i; j-- {
				s2 := sum[j+1] ^ sum[i]
				ok := i == 0 && j == n-1 || // 递归入口
					rightBits < 0 || rightBits&s2 != 0 ||  // 能从 f(i,R) 递归到 f(i,j)
					leftBits[j] < 0 || leftBits[j]&s2 != 0 // 能从 f(L,j) 递归到 f(i,j)
				if ok {
					if s2 == 0 {
						leftBits[j] = -1
						rightBits = -1
					} else {
						high := 1 << (bits.Len(uint(s2)) - 1)
						leftBits[j] |= high
						rightBits |= high
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
