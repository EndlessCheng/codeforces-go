package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, 2*n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	copy(a[n:], a[:n])

	idQ := make([]int, 2*n)
	l, r := 0, 1
	for i := 1; i < 2*n; i++ {
		for ; l < r && idQ[l] < i-n/2; l++ {
		}
		ans = max(ans, a[i]+a[idQ[l]]+i-idQ[l])
		for ; l < r && a[i]-i >= a[idQ[r-1]]-idQ[r-1]; r-- {
		}
		idQ[r] = i
		r++
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
