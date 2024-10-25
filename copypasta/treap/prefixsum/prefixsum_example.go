package multiset

import "cmp"

// https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-ii/
func findXSum(a []int, windowSize, x int) (ans []int64) {
	type pair struct{ v, c int }
	t := newTreapWith[pair](
		func(a, b pair) int { return cmp.Or(b.c-a.c, b.v-a.v) },
		func(key pair) int { return key.v * key.c },
	)
	cnt := map[int]int{}
	for r, v := range a {
		if cnt[v] > 0 {
			t.put(pair{v, cnt[v]}, -1)
		}
		cnt[v]++
		t.put(pair{v, cnt[v]}, 1)

		l := r + 1 - windowSize
		if l < 0 {
			continue
		}

		v = a[l]
		if x >= t.size() {
			ans = append(ans, int64(t.root.getSum()))
		} else {
			ans = append(ans, int64(t.preSum(x)))
		}

		t.put(pair{v, cnt[v]}, -1)
		cnt[v]--
		if cnt[v] > 0 {
			t.put(pair{v, cnt[v]}, 1)
		}
	}
	return
}
