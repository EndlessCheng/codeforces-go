package main

import . "fmt"

// https://github.com/EndlessCheng
func eq(i, j int) (r bool) {
	Println("?", i, j)
	Scan(&r)
	return
}

func cfC() {
	var T, n int
o:
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		for i := 3; i < n*2; i += 2 {
			if eq(i, i+1) {
				Println("!", i)
				continue o
			}
		}
		if eq(1, 3) || eq(1, 4) {
			Println("!", 1)
		} else {
			Println("!", 2)
		}
	}
}

//func main() { cfC() }
