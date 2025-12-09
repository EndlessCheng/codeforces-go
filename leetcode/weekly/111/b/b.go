package main

// https://space.bilibili.com/206214
func minDeletionSize(strs []string) (ans int) {
	n, m := len(strs), len(strs[0])
	for j := range m {
		for i := range n - 1 { // 遍历 j 列的字母
			if strs[i][j] > strs[i+1][j] {
				ans++
				break
			}
		}
	}
	return
}
