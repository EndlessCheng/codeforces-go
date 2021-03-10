package main

// github.com/EndlessCheng/codeforces-go

// 另一种方法是判断三进制是否包含 2

func f(s, p int) bool {
	return s == 0 || p < 1e7 && (f(s, p*3) || f(s-p, p*3))
}

func checkPowersOfThree(n int) bool {
	return f(n, 1)
}
