package main

func countCharacters(words []string, chars string) int {
	cnt := [26]int{}
	for _, c := range chars {
		cnt[c-'a']++
	}
	ans := 0
	for _, w := range words {
		cnt2 := [26]int{}
		for _, c := range w {
			cnt2[c-'a']++
		}
		ok := true
		for i, c2 := range cnt2 {
			if c2 > cnt[i] {
				ok = false
				break
			}
		}
		if ok {
			ans += len(w)
		}
	}
	return ans
}
