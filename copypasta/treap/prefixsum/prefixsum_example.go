package multiset

// https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-ii/
func findXSum(a []int, windowSize, x int) (ans []int64) {
	t := newTreapPair()
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
