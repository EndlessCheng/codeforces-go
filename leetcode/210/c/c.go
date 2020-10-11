package main

// github.com/EndlessCheng/codeforces-go
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j && s[i] == s[j] {
		i++
		j--
	}
	return i >= j
}

func check(a, b string) bool {
	n := len(a)
	i, j := n/2, n/2
	if n&1 == 0 {
		i--
	}
	for i >= 0 && a[i] == a[j] {
		i--
		j++
	}
	i++
	return isPalindrome(a[:i]+b[j:]) || isPalindrome(b[:i]+a[j:])
}

func checkPalindromeFormation(a, b string) (ans bool) {
	return check(a, b) || check(b, a)
}
