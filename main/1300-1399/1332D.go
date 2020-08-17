package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func CF1332D() {
	k, I := 0, 1<<17
	Scan(&k)
	Println(2, 3)
	Println(I|k, I, 0)
	Println(k, I|k, k)
}

//func main() { CF1332D() }
