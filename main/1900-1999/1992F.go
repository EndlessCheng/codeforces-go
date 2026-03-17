package main

import (
	. "fmt"
	"io"
	"maps"
)

// https://github.com/EndlessCheng
func cf1992F(in io.Reader, out io.Writer) {
	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		ans := 1
		set := map[int]bool{1: true}
		for range n {
			Fscan(in, &v)
			tmp := maps.Clone(set)
			for w := range tmp {
				if x%(v*w) == 0 {
					set[v*w] = true
				}
			}
			if set[x] {
				ans++
				set = map[int]bool{1: true, v: true}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1992F(bufio.NewReader(os.Stdin), os.Stdout) }
