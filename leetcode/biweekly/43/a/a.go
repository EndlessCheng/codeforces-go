package main

// https://space.bilibili.com/206214
func totalMoney(n int) int {
	const D = 7
	w, r := n/D, n%D
	return (w*D*(w+D) + r*(w*2+r+1)) / 2
}
