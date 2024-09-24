package main

import (
	. "fmt"
	"math/rand"
	"sort"
)

// https://github.com/EndlessCheng
// 失败概率见 https://www.luogu.com.cn/article/xjo8wc2w
func cf1114E() {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, v, d int
	Scan(&n)
	mx := sort.Search(1e9, func(x int) bool {
		Println(">", x)
		Scan(&v)
		return v == 0
	})
	pre := mx
	for i := 0; i < 30; i++ {
		Println("?", rand.Intn(n)+1)
		Scan(&v)
		d = gcd(d, pre-v)
		pre = v
	}
	if d < 0 {
		d = -d
	}
	Println("!", mx-(n-1)*d, d)
}

//func main() { cf1114E() }
