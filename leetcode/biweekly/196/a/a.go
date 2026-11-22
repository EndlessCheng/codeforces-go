package main

// https://space.bilibili.com/206214
func assignUmbrella(rain []int) []string {
	n := len(rain)
	ans := make([]string, n)
	alice, bob := 0, 0
	for i, x := range rain {
		if x == 0 {
			ans[i] = "None"
		} else if alice <= bob && i < n-1 && rain[i+1] > 0 {
			alice++
			ans[i] = "Alice"
		} else {
			bob++
			ans[i] = "Bob"
		}
	}
	return ans
}
