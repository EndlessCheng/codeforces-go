package main

func hasAllCodes(s string, k int) bool {
	has := make([]bool, 1<<k)
	cnt := 0
	mask := 1<<k - 1
	x := 0
	for i, ch := range s {
		// 把 ch&1 加到 x 的末尾：x 整体左移一位，然后或上 ch&1
		// &mask 目的是去掉超出 k 的比特位
		x = x<<1&mask | int(ch&1)
		if i < k-1 || has[x] {
			continue
		}
		has[x] = true
		cnt++
		if cnt == 1<<k {
			return true
		}
	}
	return false
}
