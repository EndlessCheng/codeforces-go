package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	ans := int(1e18)
	for left, mid, right := 1, 2, 3; mid < n-1; mid++ {
		midS := s[mid]
		for abs(s[left]*2-midS) > abs(s[left+1]*2-midS) {
			left++
		}
		for abs(s[right]*2-midS-s[n]) > abs(s[right+1]*2-midS-s[n]) {
			right++
		}
		a, b, c, d := s[left], midS-s[left], s[right]-midS, s[n]-s[right]
		ans = min(ans, max(max(max(a, b), c), d)-min(min(min(a, b), c), d))
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
