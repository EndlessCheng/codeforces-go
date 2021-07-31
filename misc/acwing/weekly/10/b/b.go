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
	var n, tot, sum, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i]
	}
	for _, v := range a[:n-1] {
		sum += v
		if sum*2 == tot {
			ans++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
