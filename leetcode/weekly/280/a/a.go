package main

// 在辗转相除法的基础上稍作修改，累加两数相除时的商，即为答案。

// github.com/EndlessCheng/codeforces-go
func countOperations(num1, num2 int) (ans int) {
	for num1 > 0 {
		ans += num2 / num1
		num1, num2 = num2%num1, num1
	}
	return
}
