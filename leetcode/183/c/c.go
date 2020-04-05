package main

import (
	"bytes"
	"sort"
)

type pair struct {
	cnt    int
	b      byte
	others []byte
}

func longestDiverseString(a int, b int, c int) string {
	ans := []byte{}
	ps := []pair{{a, 'a', nil}, {b, 'b', nil}, {c, 'c', nil}}
	sort.Slice(ps, func(i, j int) bool { return ps[i].cnt > ps[j].cnt })
	data := []pair{}
	for ps[0].cnt > 0 {
		l := 2
		if ps[0].cnt == 1 {
			l = 1
		}
		ps[0].cnt -= l
		data = append(data, pair{l, ps[0].b, nil})
	}
	for i := range data {
		if ps[1].cnt > 0 {
			data[i].others = append(data[i].others, ps[1].b)
			ps[1].cnt--
		} else if ps[2].cnt > 0 {
			data[i].others = append(data[i].others, ps[2].b)
			ps[2].cnt--
		}
	}
	for i := range data {
		for j := 1; j <= 2; j++ {
			if ps[j].cnt > 0 {
				b := ps[j].b
				data[i].others = append(data[i].others, b)
				ps[j].cnt--
				if ps[j].cnt > 0 && bytes.Count(data[i].others, []byte{b}) == 1 {
					data[i].others = append(data[i].others, b)
					ps[j].cnt--
				}
			}
		}
	}
	for i, d := range data {
		if i > 0 && len(d.others) == 0 && len(data[i-1].others) == 0 {
			break
		}
		for j := 0; j < d.cnt; j++ {
			ans = append(ans, d.b)
		}
		ans = append(ans, d.others...)
	}
	return string(ans)
}
