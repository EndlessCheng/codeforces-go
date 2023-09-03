package main

// https://space.bilibili.com/206214
func checkStrings(x, y string) bool {
	cnt := [2][26]int{}
	for i, b := range x {
		cnt[i&1][b-'a']++
		cnt[i&1][y[i]-'a']--
	}
	for _, c := range cnt {
		for _, c := range c {
			if c != 0 {
				return false
			}
		}
	}
	return true
}
