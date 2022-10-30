package main

// https://space.bilibili.com/206214
func oddString(words []string) string {
	m := map[string][]string{}
	d := make([]byte, len(words[0])-1)
	for _, s := range words {
		for i := 0; i < len(s)-1; i++ {
			d[i] = s[i] - s[i+1]
		}
		t := string(d)
		m[t] = append(m[t], s)
	}
	for _, g := range m {
		if len(g) == 1 {
			return g[0]
		}
	}
	return ""
}
