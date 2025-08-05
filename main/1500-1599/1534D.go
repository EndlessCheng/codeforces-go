package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://github.com/EndlessCheng
func cf1534D() {
	in := bufio.NewReader(os.Stdin)
	var n, odd, tar int
	Fscan(in, &n)
	q := func(rt int) []int {
		Println("?", rt)
		d := make([]int, n)
		for i := range d {
			Fscan(in, &d[i])
		}
		return d
	}
	dis0 := q(1)
	for _, d := range dis0 {
		odd += d % 2
	}
	if odd*2 < n {
		tar = 1
	}

	ans := [][2]int{}
	for i, d := range dis0 {
		if d%2 != tar {
			continue
		}
		dis := dis0
		if i > 0 {
			dis = q(i + 1)
		}
		for j, d := range dis {
			if d == 1 {
				ans = append(ans, [2]int{i + 1, j + 1})
			}
		}
	}
	Println("!")
	for _, e := range ans {
		Println(e[0], e[1])
	}
}

//func main() { cf1534D() }
