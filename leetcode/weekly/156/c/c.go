package main

import "bytes"

// https://space.bilibili.com/206214
func removeDuplicates(s string, k int) string {
	type pair struct {
		ch  byte
		cnt int
	}
	st := []pair{{}} // 加个哨兵，无需判断栈是否为空

	for _, ch := range s {
		m := len(st)
		if st[m-1].ch != byte(ch) { // ch 与栈顶字母不同
			st = append(st, pair{byte(ch), 1}) // 创建一个新的 pair，计数器从 1 开始
		} else if st[m-1].cnt == k-1 { // 连续 k 个相同字母
			st = st[:m-1] // 消除
		} else { // 相同但无法消除
			st[m-1].cnt++ // 只需把计数器增加 1
		}
	}

	ans := []byte{}
	for _, p := range st {
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}
