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
	var n, sum, s, j, ans int64
	Fscan(in, &n, &sum)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
		for s-a[j] >= sum {
			s -= a[j]
			j++
		}
		if s >= sum {
			ans += j + 1
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
