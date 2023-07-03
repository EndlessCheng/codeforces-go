package main

import "fmt"

// https://space.bilibili.com/206214
const mx = 1e6

var primes []int
var isP = [mx + 1]bool{}

func init() {
	for i := 2; i <= mx; i++ {
		isP[i] = true
	}
	for i := 2; i <= mx; i++ {
		if isP[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				isP[j] = false
			}
		}
	}
}

func findPrimePairs(n int) (ans [][]int) {
	if n%2 > 0 {
		if n > 4 && isP[n-2] {
			return [][]int{{2, n - 2}}
		}
		return
	}
	for _, x := range primes {
		y := n - x
		if y < x {
			break
		}
		if isP[y] {
			ans = append(ans, []int{x, y})
		}
	}
	return
}

func main() {
	fmt.Println(len(primes))
}
