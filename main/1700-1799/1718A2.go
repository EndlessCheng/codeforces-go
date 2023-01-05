package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1718A2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans, xor := n, 0
		has := map[int]bool{0: true}
		for ; n > 0; n-- {
			Fscan(in, &v)
			xor ^= v
			if has[xor] {
				ans--
				has = map[int]bool{xor: true}
			} else {
				has[xor] = true
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1718A2(os.Stdin, os.Stdout) }
