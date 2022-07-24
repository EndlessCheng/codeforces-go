package main

// https://space.bilibili.com/206214/dynamic
func repeatedCharacter(s string) byte {
	mask := 0
	for _, c := range s {
		if mask>>(c&31)&1 > 0 {
			return byte(c)
		}
		mask |= 1 << (c & 31)
	}
	return 0
}
