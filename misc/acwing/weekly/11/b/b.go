package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	sq := map[int]bool{}
	for i := 1; i*i <= 2e9; i++ {
		sq[i*i] = true
	}

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := 99
		for i := uint(1); i < 1<<len(s); i++ {
			if s[bits.TrailingZeros(i)] == '0' {
				continue
			}
			t := []byte{}
			for j := i; j > 0; j &= j - 1 {
				t = append(t, s[bits.TrailingZeros(j)])
			}
			v, _ := strconv.Atoi(string(t))
			if sq[v] {
				ans = min(ans, len(s)-len(t))
			}
		}
		if ans == 99 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
