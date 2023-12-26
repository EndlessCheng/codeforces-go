package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1917B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, n, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans, m := 0, 0
		for _, b := range s {
			m |= 1 << (b - 'a')
			ans += bits.OnesCount(uint(m))
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1917B(os.Stdin, os.Stdout) }
