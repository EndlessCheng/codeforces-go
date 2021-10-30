package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)

	for i := n - 1; ; i-- {
		v := a[i]
		if v < 0 {
			Fprint(out, v)
			return
		}
		rt := int(math.Sqrt(float64(v)))
		if rt*rt < v {
			Fprint(out, v)
			return
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
