package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type uf06 map[int32]int32

func (u uf06) find(x int32) int32 {
	fx, ok := u[x]
	if !ok {
		u[x] = x
		fx = x
	}
	if fx != x {
		u[x] = u.find(fx)
		return u[x]
	}
	return x
}

func (u uf06) merge(x, y int32) {
	u[u.find(x)] = u.find(y)
}

func (u uf06) same(x, y int32) bool {
	return u.find(x) == u.find(y)
}

func cf506D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, c, q int32
	Fscan(in, &n, &m)
	us := make([]uf06, m+1)
	for i := range us {
		us[i] = uf06{}
	}
	cs := make([][]int32, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &c)
		cs[v] = append(cs[v], c)
		cs[w] = append(cs[w], c)
		us[c].merge(v, w)
	}
	for _, u := range us {
		for v := range u {
			u.find(v)
		}
	}
	for i, a := range cs {
		if a != nil {
			slices.Sort(a)
			cs[i] = slices.Compact(a)
		}
	}

	Fscan(in, &q)
	type pair struct{ v, w int32 }
	memo := make(map[pair]int32, q)
	for ; q > 0; q-- {
		Fscan(in, &v, &w)
		if len(cs[v]) > len(cs[w]) {
			v, w = w, v
		}
		p := pair{v, w}
		if ans, ok := memo[p]; ok {
			Fprintln(out, ans)
			continue
		}
		ans := int32(0)
		for _, c := range cs[v] {
			u := us[c]
			if u[v] == u[w] {
				ans++
			}
		}
		Fprintln(out, ans)
		memo[p] = ans
	}
}

//func main() { cf506D(os.Stdin, os.Stdout) }
