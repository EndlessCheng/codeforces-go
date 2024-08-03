package main

import . "fmt"

func cf1088D() {
	ans := [2]int{}
	var cmp, t int
	Println("? 0 0")
	Scan(&cmp)
	for b := 1 << 29; b > 0; b >>= 1 {
		Println("?", ans[0]|b, ans[1]|b)
		Scan(&t)
		if t != 0 && t != cmp { // 两个比特一定不同
			ans[(t+1)/2] |= b
			// 看看有没有相等的机会
			Println("?", ans[0], ans[1])
			Scan(&cmp)
		} else { // 两个比特一定相同，要么 0 0 要么 1 1
			Println("?", ans[0]|b, ans[1])
			Scan(&t)
			if t < 0 { // 是 1 1
				ans[0] |= b
				ans[1] |= b
			}
			// cmp 不变
		}
	}
	Println("!", ans[0], ans[1])
}

//func main() { cf1088D() }
