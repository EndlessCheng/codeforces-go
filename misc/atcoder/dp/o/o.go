package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7

	var n int
	Fscan(in, &n)
	ok := make([][]bool, n)
	for i := range ok {
		ok[i] = make([]bool, n)
		for j := range ok[i] {
			Fscan(in, &ok[i][j])
		}
	}
	m := 1 << n
	dp := make([]int, m)
	dp[0] = 1
	for s, dv := range dp {
		i := bits.OnesCount(uint(s))
		for t, lb := m-1^s, 0; t > 0; t ^= lb {
			lb = t & -t
			if ok[i][bits.TrailingZeros(uint(lb))] {
				dp[s|lb] = (dp[s|lb] + dv) % mod
			}
		}
	}
	Fprint(out, dp[m-1])
}

func main() { run(os.Stdin, os.Stdout) }
