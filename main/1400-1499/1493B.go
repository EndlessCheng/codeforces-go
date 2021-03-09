package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1493B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	ds := [][2]int{{}, {1, 1}, {2, 5}, {5, 2}, {8, 8}}

	var T, H, M, h, m, ansH, ansM int
	for Fscanln(in, &T); T > 0; T-- {
		Fscanln(in, &H, &M)
		Fscanf(in, "%d:%d\n", &h, &m)
		t0, minD := h*M+m, int(1e9)
		for _, h0 := range ds {
			for _, h1 := range ds {
				h := h0[0]*10 + h1[0]
				if h >= H || h1[1]*10+h0[1] >= M {
					continue
				}
				for _, m0 := range ds {
					for _, m1 := range ds {
						m := m0[0]*10 + m1[0]
						if m >= M || m1[1]*10+m0[1] >= H {
							continue
						}
						t := h*M + m
						if t < t0 {
							t += H * M
						}
						if t-t0 < minD {
							minD, ansH, ansM = t-t0, h, m
						}
					}
				}
			}
		}
		Fprintf(out, "%02d:%02d\n", ansH, ansM)
	}
}

//func main() { CF1493B(os.Stdin, os.Stdout) }
