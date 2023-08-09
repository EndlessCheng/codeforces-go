package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func f(a []int, k, pre int) int {
	if k < 0 || a[0] == a[len(a)-1] {
		return 0
	}
	bit := 1 << k
	if a[0]&bit == a[len(a)-1]&bit {
		return f(a, k-1, pre|a[0]&bit)
	}
	i := sort.SearchInts(a, pre|bit)
	return min(f(a[:i], k-1, pre), f(a[i:], k-1, pre|bit)) | bit
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	Fprint(out, f(a, 29, 0))
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
