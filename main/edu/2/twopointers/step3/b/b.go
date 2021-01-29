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
	var n, sum, s, ans int64
	Fscan(in, &n, &sum)
	a := make([]int64, n)
	j := 0
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
		for s > sum {
			s -= a[j]
			j++
		}
		ans += int64(i-j+2) * int64(i-j+1) / 2
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
