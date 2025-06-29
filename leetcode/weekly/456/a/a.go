package main

// https://space.bilibili.com/206214
func partitionString(s string) (ans []string) {
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
