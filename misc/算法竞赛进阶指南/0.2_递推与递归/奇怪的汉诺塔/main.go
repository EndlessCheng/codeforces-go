package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 12

	d := [mx + 1]int{0, 1}
	for i := 2; i <= mx; i++ {
		d[i] = 2*d[i-1] + 1
	}
	f := [mx + 1]int{}
	for i := range f {
		f[i] = 1e9
	}
	f[1] = 1
	for i := 2; i <= mx; i++ {
		for j := 1; j < i; j++ {
			f[i] = min(f[i], 2*f[j]+d[i-j])
		}
	}
	for _, v := range f[1:] {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
