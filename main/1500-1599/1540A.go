package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1540A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		d := make([]int, n)
		for i := range d {
			Fscan(in, &d[i])
		}
		sort.Ints(d)
		ans := int64(0)
		for i := int64(1); i < n; i++ {
			ans += int64(d[i]-d[i-1]) * (1 - i*(n-i))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1540A(os.Stdin, os.Stdout) }
