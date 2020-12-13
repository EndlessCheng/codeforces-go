package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var r, n, ans int
	Fscan(in, &r, &n)
	a := make([]int, n, n+2)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(a, 0, r+1)
	sort.Ints(a)

	mi := int(2e9)
	for i := 1; i < n+2; i++ {
		if d := a[i] - a[i-1] - 1; d > 0 && d < mi {
			mi = d
		}
	}
	if mi == 2e9 {
		Fprint(out, 0)
		return
	}
	for i := 1; i < n+2; i++ {
		if d := a[i] - a[i-1] - 1; d > 0 {
			ans += (d-1)/mi + 1
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
