package main

import (
	. "fmt"
)

// https://github.com/EndlessCheng
func cf1286C1() {
	var n int
	var s string
	Scan(&n)
	Println("?", 1, n)
	cnt := make([][26]int, n+1)
	for i := 1; i <= (1+n)*n/2; i++ {
		Scan(&s)
		for _, c := range s {
			cnt[len(s)][c-'a']++
		}
	}

	if n > 1 {
		Println("?", 2, n)
		for i := 1; i <= (n-1)*n/2; i++ {
			Scan(&s)
			for _, c := range s {
				cnt[len(s)][c-'a']--
			}
		}
	}

	ans := []byte{}
	for i := 1; i <= n; i++ {
		var j byte
		for j = range 26 {
			if cnt[i][j] != 0 {
				ans = append(ans, 'a'+j)
				break
			}
		}
		for k := i + 1; k <= n; k++ {
			cnt[k][j]--
		}
	}
	Println("!", string(ans))
}

//func main() { cf1286C1() }
