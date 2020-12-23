package main

// github.com/EndlessCheng/codeforces-go
func calculateTime(keyboard, word string) (ans int) {
	pos := ['z' + 1]int{}
	for i, b := range keyboard {
		pos[b] = i
	}
	word = string(keyboard[0]) + word
	for i := 1; i < len(word); i++ {
		ans += abs(pos[word[i]] - pos[word[i-1]])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
