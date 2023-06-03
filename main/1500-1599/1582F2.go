package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF1582F2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	f := [8192]int{}
	for i := 1; i < len(f); i++ {
		f[i] = len(f)
	}
	vis := [len(f) + 1]bool{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if vis[v] {
			continue
		}
		vis[v] = true
		for xor, fv := range f[:1<<bits.Len16(uint16(v))] {
			if fv < v && v < f[xor^v] {
				for k := v + 1; k <= f[xor^v]; k++ { // 注意 f[xor^v] 在不断变小
					vis[k] = false
				}
				f[xor^v] = v
			}
		}
	}
	ans := []int{}
	for i, fv := range f {
		if fv < len(f) {
			ans = append(ans, i)
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1582F2(os.Stdin, os.Stdout) }
