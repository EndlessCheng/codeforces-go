package main

// https://space.bilibili.com/206214
func removeAlmostEqualCharacters(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		if abs(int(s[i-1])-int(s[i])) <= 1 {
			ans++
			i++
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
