package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1333F(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(_r, &n)
	vis := make([]bool, n+1)
	cnts := make([]int, n/2+1)
	for i := 2; i <= n; i++ {
		if !vis[i] {
			for j := i; j <= n; j += i {
				if !vis[j] {
					vis[j] = true
					cnts[j/i]++
				}
			}
		}
	}
	for i, c := range cnts {
		for ; c > 0; c-- {
			Fprint(out, i, " ")
		}
	}
}

//func main() { CF1333F(os.Stdin, os.Stdout) }
