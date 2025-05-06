package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// https://github.com/EndlessCheng
type vec18 struct{ x, y float64 }

func (a vec18) add(b vec18) vec18 { return vec18{a.x + b.x, a.y + b.y} }
func (a vec18) rotateCCW(rad float64) vec18 {
	sin, cos := math.Sincos(rad)
	return vec18{a.x*cos - a.y*sin, a.x*sin + a.y*cos}
}

type seg18 []struct {
	l, r int
	v    vec18 // 在上一个向量的基础上，增加的偏移量
	ang  int   // 下一个向量需要旋转的角度
}

func (t seg18) maintain(o int) {
	a, b := t[o<<1], t[o<<1|1]
	t[o].v = a.v.add(b.v.rotateCCW(float64(a.ang%360) / 180 * math.Pi))
	t[o].ang = a.ang + b.ang
}

func (t seg18) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		if l > 0 { // l=0 是原点，无偏移量
			t[o].v.x = 1
		}
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg18) update(o, i, incX, incAng int) {
	if t[o].l == t[o].r {
		t[o].v.x += float64(incX)
		t[o].ang -= incAng
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, incX, incAng)
	} else {
		t.update(o<<1|1, i, incX, incAng)
	}
	t.maintain(o)
}

func cf618E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, i, inc int
	Fscan(in, &n, &m)
	t := make(seg18, 2<<bits.Len(uint(n)))
	t.build(1, 0, n)
	for range m {
		Fscan(in, &op, &i, &inc)
		if op == 1 {
			t.update(1, i, inc, 0)
		} else {
			t.update(1, i-1, 0, inc)
		}
		p := t[1].v
		Fprintf(out, "%.4f %.4f\n", p.x, p.y)
	}
}

//func main() { cf618E(bufio.NewReader(os.Stdin), os.Stdout) }
