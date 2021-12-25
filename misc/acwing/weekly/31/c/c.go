package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
const _w = bits.UintSize
func NewBitset(n int) Bitset { return make(Bitset, n/_w+1) } // (n+_w-1)/_w
type Bitset []uint
func (b Bitset) Set(p int) { b[p/_w] |= 1 << (p % _w) }

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, v, w int
	Fscan(in, &n, &m, &k)
	ps := make([]struct{ x, y int }, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
	}
	cs := make([]struct{ x, y, r int }, m)
	for i := range cs {
		Fscan(in, &cs[i].r, &cs[i].x, &cs[i].y)
	}
	bs := make([]Bitset, n)
	for i, p := range ps {
		bs[i] = NewBitset(m)
		for j, c := range cs {
			if (p.x-c.x)*(p.x-c.x)+(p.y-c.y)*(p.y-c.y) <= c.r*c.r {
				bs[i].Set(j)
			}
		}
	}
	for ; k > 0; k-- {
		Fscan(in, &v, &w)
		s := 0
		for i, x := range bs[v-1] {
			s += bits.OnesCount(x ^ bs[w-1][i])
		}
		Fprintln(out, s)
	}
}

func main() { run(os.Stdin, os.Stdout) }
