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
	var n, k, j int
	ans, c := int64(0), map[int]int{}
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		c[a[i]]++
		for len(c) > k {
			if c[a[j]] > 1 {
				c[a[j]]--
			} else {
				delete(c, a[j])
			}
			j++
		}
		ans += int64(i - j + 1)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
