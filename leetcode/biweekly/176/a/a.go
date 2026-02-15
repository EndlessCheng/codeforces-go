package main

// https://space.bilibili.com/206214
func mapWordWeights(words []string, weights []int) string {
	ans := make([]byte, len(words))
	for i, w := range words {
		sum := 0
		for _, ch := range w {
			sum += weights[ch-'a']
		}
		ans[i] = 'z' - byte(sum%26)
	}
	return string(ans)
}
