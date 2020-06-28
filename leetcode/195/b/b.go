package main

func canArrange(a []int, k int) (ans bool) {
	cnt := make([]int, k)
	for _, v := range a {
		cnt[(v%k+k)%k]++
	}
	for i, c := range cnt {
		if i == 0 && c&1 > 0 || i > 0 && c != cnt[k-i] {
			return
		}
	}
	return true
}
