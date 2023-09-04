package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([][3]int, n)
	for i := range a {
		Fscan(in, &a[i][0], &a[i][1], &a[i][2])
		sort.Ints(a[i][:])
	}
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })

	f := fenwick{}
	for i := 0; i < n; {
		st := i
		for ; i < n && a[i][0] == a[st][0]; i++ {
			if f.pre(a[i][1]-1) < a[i][2] {
				Fprint(out, "Yes")
				return
			}
		}
		for ; st < i; st++ {
			f.upd(a[st][1], a[st][2])
		}
	}
	Fprint(out, "No")
}

func main() { run(os.Stdin, os.Stdout) }

type fenwick map[int]int

func (f fenwick) upd(i, val int) {
	for ; i <= 1e9; i += i & -i {
		if fi, ok := f[i]; !ok || val < fi {
			f[i] = val
		}
	}
}

func (f fenwick) pre(i int) int {
	res := int(1e9)
	for ; i > 0; i &= i - 1 {
		if fi, ok := f[i]; ok && fi < res {
			res = fi
		}
	}
	return res
}
