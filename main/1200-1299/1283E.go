package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1283E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, low, up int
	Fscan(in, &n)
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	prevLow, prevUp := -1, -1
	for i, c := range cnt {
		if c > 0 && i > prevLow+1 {
			prevLow = i + 1
			low++
		}
		if c > 0 && i > prevUp+1 {
			prevUp = i - 1
			up++
			c--
		}
		if c > 0 && i == prevUp+1 {
			prevUp = i
			up++
			c--
		}
		if c > 0 && i == prevUp {
			prevUp = i + 1
			up++
		}
	}
	Fprint(out, low, up)
}

//func main() { CF1283E(os.Stdin, os.Stdout) }
