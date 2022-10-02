package main

// https://space.bilibili.com/206214
func equalFrequency(word string) bool {
next:
	for i := range word {
		cnt := map[rune]int{}
		for _, c := range word[:i] + word[i+1:] {
			cnt[c]++
		}
		same := 0
		for _, c := range cnt {
			if same == 0 {
				same = c
			} else if c != same {
				continue next
			}
		}
		return true
	}
	return false
}
