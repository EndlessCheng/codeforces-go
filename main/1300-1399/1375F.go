package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1375F() {
	const mx int = 9e9
	var a [7]int
	var x, y int
	Scan(&a[1], &a[2], &a[3])
	Println("First")
	Println(mx)
	Scan(&x)
	a[x] += mx
	Println(a[x]*3 - a[1] - a[2] - a[3])
	Scan(&y)
	Println(a[x] - a[6-x-y])
}

//func main() { cf1375F() }
