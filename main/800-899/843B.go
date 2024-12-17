package main

import (
	. "fmt"
	"math/rand"
)

// https://github.com/EndlessCheng
func cf843B() {
	var n, st, tar, v, nxt int
	Scan(&n, &st, &tar)
	cur := -1
	for _, i := range rand.Perm(n)[:min(999, n)] {
		Println("?", i+1)
		Scan(&v, &nxt)
		if v <= tar && v > cur {
			cur, st = v, nxt
		}
	}
	for ; st > 0 && cur < tar; cur, st = v, nxt {
		Println("?", st)
		Scan(&v, &nxt)
	}
	if cur < tar {
		cur = -1
	}
	Print("! ", cur)
}

//func main() { cf843B() }
