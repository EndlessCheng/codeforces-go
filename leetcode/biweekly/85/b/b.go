package main

// https://space.bilibili.com/206214
func secondsToRemoveOccurrences(s string) (f int) {
	pre0 := 0
	for _, c := range s {
		if c == '0' {
			pre0++
		} else if pre0 > 0 {
			f = max(f+1, pre0)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
