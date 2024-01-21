package main

// https://space.bilibili.com/206214
func minimumPushes(word string) int {
	n := len(word)
	k := n / 8
	return (k*4 + n%8) * (k + 1)
}
