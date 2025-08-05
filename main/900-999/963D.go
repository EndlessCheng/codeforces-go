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
const w63 = bits.UintSize

type bitset63 []uint

func (b bitset63) set(p int) { b[p/w63] |= 1 << (p % w63) }

func (b bitset63) rsh(k int) bitset63 {
	if k == 0 {
		return b
	}
	shift, offset := k/w63, k%w63
	n := len(b)
	if shift >= n {
		return make(bitset63, n)
	}
	b = slices.Clone(b)
	lim := n - 1 - shift
	if offset == 0 {
		copy(b, b[shift:])
	} else {
		for i := 0; i < lim; i++ {
			b[i] = b[i+shift]>>offset | b[i+shift+1]<<(w63-offset)
		}
		b[lim] = b[n-1] >> offset
	}
	clear(b[lim+1:])
	return b
}

func (b bitset63) and(c bitset63) {
	for i, v := range c {
		b[i] &= v
	}
}

func cf963D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s, t string
	var q, k int
	Fscan(in, &s, &q)
	n := len(s)
	pos := [26]bitset63{}
	for i := range pos {
		pos[i] = make(bitset63, (n+w63-1)/w63)
	}
	for i, b := range s {
		pos[b-'a'].set(i)
	}

	match := make(bitset63, (n+w63-1)/w63)
	idx := []int{}
	for range q {
		Fscan(in, &k, &t)
		m := len(t)
		if m+k-1 > n {
			Fprintln(out, -1)
			continue
		}

		for i := range match {
			match[i] = math.MaxUint
		}
		for i, b := range t {
			match.and(pos[b-'a'].rsh(i))
		}
		idx = idx[:0]
		ans := int(1e9)
		for i, v := range match {
			for ; v > 0; v &= v - 1 {
				j := i*w63 | bits.TrailingZeros(v)
				idx = append(idx, j)
				if len(idx) >= k {
					ans = min(ans, j-idx[len(idx)-k])
				}
			}
		}

		if len(idx) < k {
			Fprintln(out, -1)
		} else {
			Fprintln(out, ans+m)
		}
	}
}

//func main() { cf963D(bufio.NewReader(os.Stdin), os.Stdout) }
