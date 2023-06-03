package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF1582F1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, v int
	f := [512]int{}
	for i := 1; i < len(f); i++ {
		f[i] = 1e9
	}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		for j, fv := range f[:1<<bits.Len(uint(v))] {
			if fv < v {
				f[j^v] = min(f[j^v], v)
			}
		}
	}
	ans := []int{}
	for i, fv := range f { // 可以选空子序列
		if fv < 1e9 {
			ans = append(ans, i)
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1582F1(os.Stdin, os.Stdout) }
