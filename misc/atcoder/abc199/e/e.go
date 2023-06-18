package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, M int
	Fscan(in, &n, &M)
	rules := make([]struct{ x, y, z int }, M)
	for i := range rules {
		Fscan(in, &rules[i].x, &rules[i].y, &rules[i].z)
	}

	m := 1 << n
	u := m - 1
	f := make([]int, m)
	f[0] = 1
	for s, dv := range f {
		i := bits.OnesCount(uint(s))
	o:
		for cus, lb := u^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			ns := s | lb
			for _, p := range rules { // 检查每个要求
				if i < p.x && bits.OnesCount(uint(ns&(1<<p.y-1))) > p.z {
					continue o
				}
			}
			f[ns] += dv
		}
	}
	Fprint(out, f[m-1])
}

func main() { run(os.Stdin, os.Stdout) }
