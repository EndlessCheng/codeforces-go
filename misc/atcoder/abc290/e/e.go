package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	m := n / 2
	ans := m * (m + 1) * (m*4 + n%2*6 - 1) / 6
	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		pos[v-1] = append(pos[v-1], i)
	}
	for _, ps := range pos {
		l, r := 0, len(ps)-1
		for l < r {
			if ps[l]+ps[r] < n {
				ans -= (ps[l] + 1) * (r - l)
				l++
			} else {
				ans -= (n - ps[r]) * (r - l)
				r--
			}
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
