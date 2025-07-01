package main

// https://space.bilibili.com/206214
func partitionString1(s string) (ans []string) {
	vis := map[string]bool{}
	left := 0
	for i := range s {
		t := s[left : i+1]
		if !vis[t] {
			vis[t] = true
			ans = append(ans, t)
			left = i + 1
		}
	}
	return
}

func partitionString(s string) (ans []string) {
	type node struct{ son [26]*node }
	root := &node{}
	cur := root
	left := 0
	for i, c := range s {
		c -= 'a'
		if cur.son[c] == nil { // 无路可走？
			cur.son[c] = &node{} // 那就造路！
			ans = append(ans, s[left:i+1])
			left = i + 1
			cur = root // 重置
		} else {
			cur = cur.son[c]
		}
	}
	return
}
