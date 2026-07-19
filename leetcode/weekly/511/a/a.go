package main

// https://space.bilibili.com/206214
func canReach(start, target []int) bool {
	return (start[0]+start[1])%2 == (target[0]+target[1])%2
}
