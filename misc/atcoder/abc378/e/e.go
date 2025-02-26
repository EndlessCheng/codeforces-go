package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func mergeCount(a []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	left := append(a[:0:0], a[:n/2]...)
	right := append(a[:0:0], a[n/2:]...)
	cnt := mergeCount(left) + mergeCount(right)
	l, r := 0, 0
	for i := range a {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			a[i] = left[l]
			l++
		} else {
			cnt += n/2 - l
			a[i] = right[r]
			r++
		}
	}
	return cnt
}

func run(in io.Reader, out io.Writer) {
	var n, m, ans, ss int
	Fscan(in, &n, &m)
	s := make([]int, n+1)
	for r := 1; r <= n; r++ {
		Fscan(in, &s[r])
		s[r] = (s[r] + s[r-1]) % m
		ans += r*s[r] - ss
		ss += s[r]
	}
	Fprint(out, ans+mergeCount(s)*m)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
