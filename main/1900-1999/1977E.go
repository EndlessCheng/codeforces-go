package main

import . "fmt"

// https://github.com/EndlessCheng
func q77(i, j int) bool {
	Println("?", i+1, j+1)
	s := ""
	Scan(&s)
	return s[0] == 'Y'
}

func cf1977E() {
	var T, n int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		c := make([]int, n)
		for i, j := 0, 0; i+1 < n; i = j {
			j = i + 1
			for j < n && q77(j-1, j) {
				c[j] = c[i]
				j++
			}
			if j == n {
				break
			}
			if i == 0 || q77(i-1, j) {
				c[j] = c[i] ^ 1
				continue
			}
			c[j] = c[i]
			c[j-1] ^= 1
			for k := j - 2; k > i && !q77(k, j); k-- {
				c[k] ^= 1
			}
		}

		Print("!")
		for _, v := range c {
			Print(" ", v)
		}
		Println()
	}
}

//func main() { cf1977E() }
