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
	bad := [1e6 + 1]bool{}
	vis := [1e6 + 1]bool{}
	var n, v, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &v)
		a[i] = v
		if vis[v] {
			bad[v] = true
		} else {
			vis[v] = true
		}
		if !bad[v] {
			for j := 2 * v; j <= 1e6; j += v {
				bad[j] = true
			}
		}
	}
	for _, v := range a {
		if !bad[v] {
			ans++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
