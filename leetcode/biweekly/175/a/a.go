package main

// https://space.bilibili.com/206214
func reverse(t []byte, f func(byte) bool) {
	i, j := 0, len(t)-1
	for i < j {
		for i < j && f(t[i]) {
			i++
		}
		for i < j && f(t[j]) {
			j--
		}
		t[i], t[j] = t[j], t[i]
		i++
		j--
	}
}

func reverseByType(s string) string {
	t := []byte(s)
	reverse(t, func(b byte) bool { return 'a' <= b && b <= 'z' })
	reverse(t, func(b byte) bool { return !('a' <= b && b <= 'z') })
	return string(t)
}
