package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type fenwick struct {
	tree []int
}

func newFenwickTree(n int) fenwick {
	return fenwick{make([]int, n+1)}
}
func (f fenwick) add(i, v int) {
	for i++; i < len(f.tree); i += i & -i {
		f.tree[i] += v
	}
}
func (f fenwick) sum(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}
func (f fenwick) query(l, r int) (res int) {
	return f.sum(r) - f.sum(l-1)
}

func busiestServers(k int, starts, delta []int) (ans []int) {
	n := len(starts)
	type event struct{ p, d, i int }
	es := make([]event, 0, 2*n)
	for i, s := range starts {
		es = append(es, event{s, 1, i}, event{s + delta[i], -1, i})
	}
	sort.Slice(es, func(i, j int) bool { a, b := es[i], es[j]; return a.p < b.p || a.p == b.p && a.d < b.d })

	cnt := make([]int, k)
	maxCnt := 0
	running := make([]int, n)
	ignore := make([]bool, n)
	t := newFenwickTree(k)
	for _, e := range es {
		if e.d > 0 {
			id := e.i % k
			p := id + sort.Search(k-id, func(i int) bool { return t.query(id, id+i) != i+1 })
			if p == k {
				p = sort.Search(id, func(i int) bool { return t.sum(i) != i+1 })
				if p == id {
					ignore[e.i] = true
					continue
				}
			}
			if cnt[p]++; cnt[p] > maxCnt {
				maxCnt = cnt[p]
			}
			running[e.i] = p
			t.add(p, 1)
		} else {
			if ignore[e.i] {
				continue
			}
			t.add(running[e.i], -1)
		}
	}
	for i, v := range cnt {
		if v == maxCnt {
			ans = append(ans, i)
		}
	}
	return
}
