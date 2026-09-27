package main

// https://space.bilibili.com/206214
func minQueenMoves(source, target []int) int {
	sr, sc := source[0], source[1]
	tr, tc := target[0], target[1]
	if sr == tr && sc == tc {
		return 0
	}
	if sr == tr || sc == tc || sr+sc == tr+tc || sr-sc == tr-tc {
		return 1
	}
	return 2
}
