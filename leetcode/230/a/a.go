package main

// github.com/EndlessCheng/codeforces-go
func countMatches(items [][]string, ruleKey string, ruleValue string) (ans int) {
	id := map[string]int{"type": 0, "color": 1, "name": 2}[ruleKey]
	for _, item := range items {
		if item[id] == ruleValue {
			ans++
		}
	}
	return
}
