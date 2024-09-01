package main

// https://space.bilibili.com/206214
func checkTwoChessboards(s, t string) bool {
	return (s[0]^s[1])&1 == (t[0]^t[1])&1
}
