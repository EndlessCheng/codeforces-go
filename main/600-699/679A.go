package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func CF679A() {
	s, c := "", 0
	for _, v := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 4, 9, 25, 49} {
		Println(v)
		Scan(&s)
		c += len(s) & 1 // "yes"
	}
	if c < 2 {
		Println("prime")
	} else {
		Println("composite")
	}
}

//func main() { CF679A() }
