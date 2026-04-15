package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1665D() {
	var T, g int
	for Scan(&T); T > 0; T-- {
		a := 1
		for i := range 30 {
			Println("?", a, a+1<<30)
			Scan(&g)
			if g == 1<<i {
				a += 1 << i
			}
		}
		Println("!", 1<<30-a)
	}
}

//func main() { cf1665D() }
