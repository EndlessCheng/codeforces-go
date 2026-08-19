package main

import (
	. "fmt"
	"io"
	"maps"
)

// https://github.com/EndlessCheng
func cf2254F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f := func() map[int]int {
			a := make([]int, n)
			s := 0
			for i := range a {
				Fscan(in, &a[i])
				s ^= a[i]
			}
			cnt := map[int]int{s: 1}
			for _, v := range a {
				cnt[v^s]++
			}
			return cnt
		}
		if maps.Equal(f(), f()) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf2254F(bufio.NewReader(os.Stdin), os.Stdout) }
