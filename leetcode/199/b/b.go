package main

func minFlips(s string) (ans int) {
	ans = 1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] != s[i+1] {
			ans++
		}
	}
	if s[0] == '0' {
		ans--
	}
	return
}
