package main

func countTriplets(a []int) (ans int) {
	// 也可以统计 a[i]&a[j]，这样可以跳过 cnt[i]==0 的情况，效率较高
	cnt := [1 << 16]int{}
	for _, v := range a {
		for i := range cnt {
			if i&v == 0 {
				cnt[i]++
			}
		}
	}
	for _, v := range a {
		for _, w := range a {
			ans += cnt[v&w]
		}
	}
	return
}
