package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f0, f1 := 0, abs(a[1]-a[0])
	for i := 2; i < n; i++ {
		f0, f1 = f1, min(f0+abs(a[i]-a[i-2]), f1+abs(a[i]-a[i-1]))
	}
	Fprint(out, f1)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
func abs(x int) int { if x < 0 { return -x }; return x }
