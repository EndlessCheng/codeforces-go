package main

// https://space.bilibili.com/206214
func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	cnt := map[string]int{}
	for i := k; i <= n; i += k {
		cnt[word[i-k:i]]++
	}
	mx := 0
	for _, c := range cnt {
		mx = max(mx, c)
	}
	return n/k - mx
}
