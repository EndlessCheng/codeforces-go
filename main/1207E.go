package main

import (
	. "fmt"
)

// github.com/EndlessCheng/codeforces-go
func CF1207E() {
	a := [2]int{0, 7}
	for i := range a {
		Print("?")
		for j := 1; j <= 100; j++ {
			Print(" ", j<<a[i])
		}
		Println()
		Scan(&a[i])
	}
	Println("!", a[0]>>7<<7|a[1]&(1<<7-1))
}

//func main() { CF1207E() }
