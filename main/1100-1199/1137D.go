package main

import (
	"bufio"
	. "fmt"
	. "os"
)

// github.com/EndlessCheng/codeforces-go
func CF1137D() {
	in := bufio.NewScanner(Stdin)
	for {
		Println("next 0 1")
		in.Scan()
		Println("next 1")
		if in.Scan(); in.Bytes()[0] == '2' {
			break
		}
	}
	for {
		Println("next 0 1 2 3 4 5 6 7 8 9")
		if in.Scan(); in.Bytes()[0] == '1' {
			break
		}
	}
	Println("done")
}

//func main() { CF1137D() }
