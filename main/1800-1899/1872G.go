package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1872G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		sum := make([]int, n+1)
		mul := make([]int, n+1)
		mul[0] = 1
		pos := []int{}
		ok := false
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			sum[i+1] = sum[i] + v
			mul[i+1] = mul[i] * v
			if v == 1 {
				continue
			}
			if mul[i+1] >= n*2 {
				ok = true
			}
			pos = append(pos, i)
		}
		if ok {
			Fprintln(out, pos[0]+1, pos[len(pos)-1]+1)
			continue
		}
		var mx, l, r int
		for i, p := range pos {
			for _, q := range pos[i+1:] {
				d := mul[q+1]/mul[p] - sum[q+1] + sum[p]
				if d > mx {
					mx, l, r = d, p, q
				}
			}
		}
		Fprintln(out, l+1, r+1)
	}
}

//func main() { cf1872G(os.Stdin, os.Stdout) }
