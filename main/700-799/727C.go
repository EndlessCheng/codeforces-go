package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func Sol727C() {
	var n int
	Scan(&n)
	a := make([]int, n)
	Println("? 2 3")
	Scan(&a[0])
	for i := 1; i < n; i++ {
		Printf("? 1 %d\n", i+1)
		Scan(&a[i])
	}
	a[0] = (a[1] + a[2] - a[0]) / 2
	Print("! ", a[0])
	for _, v := range a[1:] {
		Print(" ", v-a[0])
	}
}

//func main() {
//	Sol727C(os.Stdin, os.Stdout)
//}
