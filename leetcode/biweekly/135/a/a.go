package main

// https://space.bilibili.com/206214
func losingPlayer(x, y int) string {
	return [2]string{"Bob", "Alice"}[min(x, y/4)%2]
}
