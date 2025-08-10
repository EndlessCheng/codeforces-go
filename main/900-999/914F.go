package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
const w14 = bits.UintSize

type bitset14 []uint

func (b bitset14) flip(p int) { b[p/w14] ^= 1 << (p % w14) }

func (b bitset14) rsh(k int) bitset14 {
	if k == 0 {
		return b
	}
	shift, offset := k/w14, k%w14
	n := len(b)
	if shift >= n {
		return make(bitset14, n)
	}
	b = slices.Clone(b)
	lim := n - 1 - shift
	if offset == 0 {
		copy(b, b[shift:])
	} else {
		for i := 0; i < lim; i++ {
			b[i] = b[i+shift]>>offset | b[i+shift+1]<<(w14-offset)
		}
		b[lim] = b[n-1] >> offset
	}
	clear(b[lim+1:])
	return b
}

func (b bitset14) and(c bitset14) {
	for i, v := range c {
		b[i] &= v
	}
}

func (b bitset14) onesCountRange(l, r int) int {
	maskL := ^uint(0) << (l % w14)
	maskR := ^uint(0) << (r % w14)
	i := l / w14
	if i == r/w14 {
		return bits.OnesCount(b[i] & (maskL ^ maskR))
	}
	cnt1 := bits.OnesCount(b[i] & maskL)
	for i++; i < r/w14; i++ {
		cnt1 += bits.OnesCount(b[i])
	}
	return cnt1 + bits.OnesCount(b[i]&^maskR)
}

func cf914F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s, t []byte
	var q, op, l, r int
	Fscan(in, &s, &q)

	n := len(s)
	pos := [26]bitset14{}
	for i := range pos {
		pos[i] = make(bitset14, (n+w14-1)/w14)
	}
	for i, b := range s {
		pos[b-'a'].flip(i)
	}

	match := make(bitset14, (n+w14-1)/w14)
	for range q {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &l, &t)
			l--
			pos[s[l]-'a'].flip(l)
			s[l] = t[0]
			pos[s[l]-'a'].flip(l)
		} else {
			Fscan(in, &l, &r, &t)
			if r-l+1 < len(t) {
				Fprintln(out, 0)
				continue
			}
			for i := range match {
				match[i] = math.MaxUint
			}
			for i, b := range t {
				match.and(pos[b-'a'].rsh(i))
			}
			Fprintln(out, match.onesCountRange(l-1, r-len(t)+1))
		}
	}
}

//func main() { cf914F(bufio.NewReader(os.Stdin), os.Stdout) }
