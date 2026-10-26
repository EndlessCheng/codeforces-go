package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2137D(in io.Reader, out io.Writer) {
	var T, n, c int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([][]int, n+1)
		for i := range n {
			Fscan(in, &c)
			pos[c] = append(pos[c], i)
		}

		ans := make([]any, n)
		cur := 0
		for i := 1; i <= n; i++ {
			if len(pos[i])%i > 0 {
				Fprintln(out, -1)
				continue o
			}
			for j, p := range pos[i] {
				if j%i == 0 {
					cur++
				}
				ans[p] = cur
			}
		}
		Fprintln(out, ans...)
	}
}

//func main() { cf2137D(bufio.NewReader(os.Stdin), os.Stdout) }
