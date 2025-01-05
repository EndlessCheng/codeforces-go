package main

// https://space.bilibili.com/206214
func calculateScore(s string) (ans int64) {
	stk := [26][]int{}
	for i, c := range s {
		c -= 'a'
		if st := stk[25-c]; len(st) > 0 {
			ans += int64(i - st[len(st)-1])
			stk[25-c] = st[:len(st)-1]
		} else {
			stk[c] = append(stk[c], i)
		}
	}
	return
}
