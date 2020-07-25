package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	buf := make([]byte, 4096)
	_i := len(buf)
	rc := func() byte {
		if _i == len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int32) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int32(b&15)
		}
		return
	}
	const no = "dldsgay!!1"
	abs := func(x int32) int32 {
		if x < 0 {
			return -x
		}
		return x
	}

	n := r()
	m := 2 * n
	a := make([]int32, n+1)
	pos := make([]int32, m+1)
	for i := int32(1); i <= n; i++ {
		v := r()
		a[i] = v
		pos[v] = i
	}
	cp := make([]int32, m+1)
	for i := int32(1); i <= n; i++ {
		v := a[i]
		w := r()
		cp[v] = w
		cp[w] = v
	}

	order := make([]int32, n+1)
	b := make([]int32, n+1)
	for i := int32(1); i <= n; i++ {
		v := r()
		if p := pos[v]; p > 0 {
			if abs(p-i)&1 > 0 {
				Fprint(out, no)
				return
			}
			order[i] = p
		}
		b[i] = v
	}
	for i := int32(1); i <= n; i++ {
		v := b[i]
		w := r()
		if cp[v] != w || cp[w] != v {
			Fprint(out, no)
			return
		}
		if p := pos[w]; p > 0 {
			if abs(p-i)&1 == 0 {
				Fprint(out, no)
				return
			}
			order[i] = p
		}
	}

	tree := make([]int32, n+1)
	add := func(i int32) {
		for ; i <= n; i += i & -i {
			tree[i]++
		}
	}
	sum := func(i int32) (res int32) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	ans := 0
	for i := int32(0); i < n; i++ {
		v := order[i+1]
		ans += int(i - sum(v))
		add(v)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
