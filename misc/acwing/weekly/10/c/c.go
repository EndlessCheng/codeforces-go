package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	cnt := make([]int, 26)
	cntPair := [26][26]int{}
	for _, b := range s {
		b -= 'a'
		for i, c := range cnt {
			cntPair[i][b] += c
		}
		cnt[b]++
	}
	ans := 0
	for _, c := range cnt {
		ans = max(ans, c)
	}
	for _, cp := range cntPair {
		for _, c := range cp {
			ans = max(ans, c)
		}
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
