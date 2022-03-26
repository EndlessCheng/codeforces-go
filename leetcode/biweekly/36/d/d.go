package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
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
	idle := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		idle.Put(i, nil)
	}
	server := make([]int, n)
	for _, e := range es {
		if e.d > 0 {
			if idle.Size() == 0 {
				server[e.i] = -1
				continue
			}
			o, _ := idle.Ceiling(e.i % k)
			if o == nil {
				o = idle.Left()
			}
			s := o.Key.(int)
			if cnt[s]++; cnt[s] > maxCnt {
				maxCnt = cnt[s]
				ans = []int{s}
			} else if cnt[s] == maxCnt {
				ans = append(ans, s)
			}
			idle.Remove(s)
			server[e.i] = s
		} else if s := server[e.i]; s >= 0 {
			idle.Put(s, nil)
		}
	}
	return
}
