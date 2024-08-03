package main

import . "fmt"

func cf1088D() {
	q := func(c, d int) (x int) {
		Println("?", c, d)
		Scan(&x)
		return
	}
	ans := [2]int{}
	cmp := q(0, 0)
	for b := 1 << 29; b > 0; b >>= 1 {
		t := q(ans[0]|b, ans[1]|b)
		if t != 0 && t != cmp { // 两个比特一定不同
			ans[(t+1)/2] |= b
			cmp = q(ans[0], ans[1])
		} else if q(ans[0]|b, ans[1]) < 0 { // 两个比特一定相同，要么 0 0 要么 1 1，这里 < 0 说明是 1 1
			ans[0] |= b
			ans[1] |= b
			// cmp 不变
		}
	}
	Println("!", ans[0], ans[1])
}

//func main() { cf1088D() }
