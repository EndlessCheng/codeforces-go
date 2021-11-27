package main

// github.com/EndlessCheng/codeforces-go
func countWords(words1, words2 []string) (ans int) {
	cnt1 := map[string]int{}
	cnt2 := map[string]int{}
	for _, s := range words1 { cnt1[s]++ } // 统计单词出现次数
	for _, s := range words2 { cnt2[s]++ } // 统计单词出现次数
	for _, s := range words2 { if cnt1[s] == 1 && cnt2[s] == 1 { ans++ }} // 单词都恰好出现一次
	return
}
