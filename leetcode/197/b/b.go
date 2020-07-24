package main

func numSub(s string) (ans int) {
	p := -1
	for i, b := range s {
		if b == '0' {
			p = i
		} else {
			ans += i - p
		}
	}
	return ans % (1e9 + 7)
}
