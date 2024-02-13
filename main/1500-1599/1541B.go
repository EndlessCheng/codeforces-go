package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1541B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, aj int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		idx := make([]int, n*2+1)
		ans := 0
		for j := 1; j <= n; j++ {
			Fscan(in, &aj)
			for ai := 1; ai*aj < j*2; ai++ {
				i := idx[ai]
				if i > 0 && ai*aj == i+j {
					ans++
				}
			}
			idx[aj] = j
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1541B(os.Stdin, os.Stdout) }
