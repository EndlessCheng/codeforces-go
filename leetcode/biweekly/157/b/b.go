package main

// https://space.bilibili.com/206214
func maxSubstrings1(word string) (ans int) {
	pos := [26]int{}
	for i, b := range word {
		b -= 'a'
		if pos[b] == 0 {
			pos[b] = i + 1
		} else if i-pos[b] > 1 {
			ans++
			clear(pos[:])
		}
	}
	return
}

func maxSubstrings(word string) (ans int) {
	seen := 0
	for i := 3; i < len(word); i++ {
		seen |= 1 << (word[i-3] - 'a')
		if seen>>(word[i]-'a')&1 > 0 { // 再次遇到 word[i]
			ans++
			seen = 0
			i += 3
		}
	}
	return
}

func maxSubstringsDP(word string) int {
	pos := [26][]int{}
	n := len(word)
	f := make([]int, n+1)
	for i, b := range word {
		b -= 'a'
		for len(pos[b]) > 1 && i-pos[b][1] > 2 {
			pos[b] = pos[b][1:]
		}
		f[i+1] = f[i] // 不选 s[i]
		if len(pos[b]) > 0 && i-pos[b][0] > 2 {
			f[i+1] = max(f[i+1], f[pos[b][0]]+1) // 选 s[i]
		}
		pos[b] = append(pos[b], i)
	}
	return f[n]
}
