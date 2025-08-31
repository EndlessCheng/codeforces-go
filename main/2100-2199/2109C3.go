package main

import . "fmt"

// https://github.com/EndlessCheng
func cf2109C3() {
	var T, n, r int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		Println("mul 999999999")
		Scan(&r)
		Println("digit")
		Scan(&r)
		if n != 81 {
			Println("add", n-81)
			Scan(&r)
		}
		Println("!")
		Scan(&r)
	}
}

//func main() { cf2109C3() }
