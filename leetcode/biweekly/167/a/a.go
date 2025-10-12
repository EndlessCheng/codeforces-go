package main

// https://space.bilibili.com/206214
func scoreBalance(s string) bool {
	total := 0
	for _, b := range s {
		total += int(b & 31)
	}

	left := 0
	for _, b := range s {
		left += int(b & 31)
		if left*2 == total {
			return true
		}
	}
	return false
}
