package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func getNoZeroIntegers1(n int) []int {
	for a := 1; ; a++ {
		if !strings.ContainsRune(strconv.Itoa(a), '0') &&
			!strings.ContainsRune(strconv.Itoa(n-a), '0') {
			return []int{a, n - a}
		}
	}
}

func getNoZeroIntegers2(n int) []int {
	for {
		a := rand.Intn(n-1) + 1
		if !strings.ContainsRune(strconv.Itoa(a), '0') &&
			!strings.ContainsRune(strconv.Itoa(n-a), '0') {
			return []int{a, n - a}
		}
	}
}

func getNoZeroIntegers(n int) []int {
	a := 0
	base := 1
	for x := n; x > 1; x /= 10 {
		d := x % 10
		if d <= 1 {
			d += 10
			x -= 10 // 借位
		}
		// a 这一位填 d/2，比如百位数就是 d/2 * 100
		a += d / 2 * base
		base *= 10
	}
	return []int{a, n - a}
}

func main() {
	test := func(n int) (cnt int) {
		for {
			cnt++
			a := 1 + rand.Intn(n-1)
			if !strings.ContainsRune(strconv.Itoa(a), '0') &&
				!strings.ContainsRune(strconv.Itoa(n-a), '0') {
				return
			}
		}
	}
	mx := 0.0
	for n := 2; n <= 1e4; n++ {
		s, u := 0, int(1e4)
		for range u {
			s += test(n)
		}
		r := float64(s) / float64(u)
		if r > mx {
			mx = r
			fmt.Println(n, r)
		}
	}

	for n := 2; n <= 1e6; n++ {
		res := getNoZeroIntegers(n)
		if strings.ContainsRune(strconv.Itoa(res[0]), '0') ||
			strings.ContainsRune(strconv.Itoa(res[1]), '0') {
			panic(n)
		}
	}
}
