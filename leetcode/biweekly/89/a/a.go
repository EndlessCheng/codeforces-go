package main

// https://space.bilibili.com/206214
func count(t string, period int) (ans int) {
	for i := 0; i < period; i++ {
		if (t[0] == '?' || i/10 == int(t[0]-'0')) &&
			(t[1] == '?' || i%10 == int(t[1]-'0')) {
			ans++
		}
	}
	return
}

func countTime(time string) int {
	return count(time[:2], 24) * count(time[3:], 60)
}
