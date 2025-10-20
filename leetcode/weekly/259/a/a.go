package main

// github.com/EndlessCheng/codeforces-go
func finalValueAfterOperations1(operations []string) (ans int) {
	for _, s := range operations {
		if s[1] == '+' {
			ans++
		} else {
			ans--
		}
	}
	return
}

func finalValueAfterOperations(operations []string) (ans int) {
	for _, s := range operations {
		ans += int(s[1]&2) - 1
	}
	return
}
