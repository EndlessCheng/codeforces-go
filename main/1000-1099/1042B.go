package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1042B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, c int
	var s []byte
	dp := [8]int{}
	for i := 1; i < 8; i++ {
		dp[i] = -1
	}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &c, &s)
		v := 0
		for _, b := range s {
			v |= 1 << (b - 'A')
		}
		for i, dpi := range dp {
			if dpi != -1 && i|v != i && (dp[i|v] == -1 || dpi+c < dp[i|v]) {
				dp[i|v] = dpi + c
			}
		}
	}
	Fprint(out, dp[7])
}

//func main() {
//	CF1042B(os.Stdin, os.Stdout)
//}
