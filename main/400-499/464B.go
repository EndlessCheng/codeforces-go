package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF464B(in io.Reader, out io.Writer) {
	perm3 := [][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	dis := func(p, q [3]int64) int64 { return (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1]) + (p[2]-q[2])*(p[2]-q[2]) }

	ps := make([][3]int64, 8)
	for i := range ps {
		Fscan(in, &ps[i][0], &ps[i][1], &ps[i][2])
	}
	var f func(int, map[int64]int8) bool
	f = func(i int, dmp map[int64]int8) bool {
		if i == 8 {
			return len(dmp) == 3
		}
		p := ps[i]
	o:
		for _, m := range perm3 {
			mp := map[int64]int8{}
			for k, c := range dmp {
				mp[k] = c
			}
			ps[i] = [3]int64{p[m[0]], p[m[1]], p[m[2]]}
			for j := 0; j < i; j++ {
				d := dis(ps[i], ps[j])
				// 剪枝：立方体的性质是距离只有三种且每种个数不超过 12
				if mp[d]++; len(mp) > 3 || mp[d] > 12 {
					continue o
				}
			}
			if f(i+1, mp) {
				return true
			}
		}
		return false
	}
	// 注意从 1 开始，避免无效运算
	if f(1, map[int64]int8{}) {
		Fprintln(out, "YES")
		for _, p := range ps {
			Fprintln(out, p[0], p[1], p[2])
		}
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF464B(os.Stdin, os.Stdout) }
