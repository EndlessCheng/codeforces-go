package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	a := [3]int{}
	var k int
	Fscan(_r, &a[0], &a[1], &a[2], &k)
	ans := 0
	for i, v := range a {
		if k <= v {
			ans += (1 - i) * k
			break
		}
		ans += (1 - i) * v
		k -= v
	}
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
