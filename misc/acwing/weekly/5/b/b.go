package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		used := map[int]bool{}
		for _, v := range a {
			for i := 0; v > 0; i++ {
				d := v % k
				v /= k
				if d == 0 {
					continue
				}
				if d > 1 || used[i] {
					Fprintln(out, "NO")
					continue o
				}
				used[i] = true
			}
		}
		Fprintln(out, "YES")
	}
}

func main() { run(os.Stdin, os.Stdout) }
