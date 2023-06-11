package main

// https://space.bilibili.com/206214
func smallestString(s string) (ans string) {
	t := []byte(s)
	for i, c := range t {
		if c > 'a' {
			for ; i < len(t) && t[i] > 'a'; i++ {
				t[i]--
			}
			return string(t)
		}
	}
	t[len(t)-1] = 'z'
	return string(t)
}
