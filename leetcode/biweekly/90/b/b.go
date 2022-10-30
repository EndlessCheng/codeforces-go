package main

// https://space.bilibili.com/206214
func twoEditWords(queries, dictionary []string) (ans []string) {
	for _, q := range queries {
		for _, s := range dictionary {
			c := 0
			for i := range s {
				if q[i] != s[i] {
					c++
				}
			}
			if c <= 2 {
				ans = append(ans, q)
				break
			}
		}
	}
	return
}
