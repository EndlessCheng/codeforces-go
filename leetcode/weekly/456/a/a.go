package main

// https://space.bilibili.com/206214
func partitionString1(s string) (ans []string) {
	vis := map[string]bool{}
	t := ""
	for _, c := range s {
		t += string(c)
		if !vis[t] {
			vis[t] = true
			ans = append(ans, t)
			t = ""
		}
	}
	return
}

func partitionString(s string) (ans []string) {
	type node struct{ son [26]*node }
	root := &node{}
	cur := root
	t := []byte{}
	for _, c := range s {
		t = append(t, byte(c))
		c -= 'a'
		if cur.son[c] == nil { // 无路可走？
			cur.son[c] = &node{} // 那就造路！
			ans = append(ans, string(t))
			t = t[:0]  // 重置
			cur = root // 重置
		} else {
			cur = cur.son[c]
		}
	}
	return
}
