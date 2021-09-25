package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, k, v int
	Fscan(in, &n, &k)
	vis := map[int]bool{}
	ans := []int{}
	for i := 1; i <= n; i++ {
		if Fscan(in, &v); !vis[v] {
			vis[v] = true
			ans = append(ans, i)
		}
	}
	if len(ans) < k {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for _, v := range ans[:k] {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
