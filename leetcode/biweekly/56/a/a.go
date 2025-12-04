package main

import (
	"fmt"
	"math"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func countTriples1(n int) (ans int) {
	for a := 1; a < n; a++ {
		for b := 1; b < a && a*a+b*b <= n*n; b++ {
			c2 := a*a + b*b
			rt := int(math.Sqrt(float64(c2)))
			if rt*rt == c2 {
				ans++
			}
		}
	}
	return ans * 2 // (a,b,c) 和 (b,a,c) 各算一次
}

func countTriples2(n int) (ans int) {
	for u := 3; u*u < n*2; u += 2 {
		for v := 1; v < u && u*u+v*v <= n*2; v += 2 {
			if gcd(u, v) == 1 {
				ans += n * 2 / (u*u + v*v)
			}
		}
	}
	return ans * 2 // (a,b,c) 和 (b,a,c) 各算一次
}

const mx int = 1e5

var mu = [mx]int{1: 1}
var divisors [mx][]int

func init() {
	// 预处理莫比乌斯函数
	// 当 n > 1 时，sum_{d|n} mu[d] = 0
	// 所以 mu[n] = -sum_{d|n ∧ d<n} mu[d]
	for i := 1; i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			mu[j] -= mu[i] // i 是 j 的真因子
		}
	}

	// 预处理不含平方因子的因子列表，用于 countCoprime
	for i := 1; i < mx; i++ {
		if mu[i] == 0 {
			continue
		}
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i) // i 是 j 的因子，且 mu[i] != 0
		}
	}
}

// 返回 [1,n] 中与 x 互质的整数个数
// 用容斥（莫反）可得 sum_{d|x} mu[d] * floor(n/d)
func countCoprime(n, x int) (res int) {
	for _, d := range divisors[x] {
		res += mu[d] * (n / d)
	}
	return
}

// 返回 [1,n] 中与奇数 x 互质的奇数个数
// 与 x 互质的整数个数 - 与 x 互质的偶数个数
func countCoprimeOdd(n, x int) (res int) {
	return countCoprime(n, x) - countCoprime(n/2, x)
}

func countTriples(n int) (ans int) {
	t0 := time.Now()
	
	for u := 3; u*u < n*2; u += 2 {
		for l, r := 1, 0; l < u && u*u+l*l <= n*2; l = r + 1 {
			num := n * 2 / (u*u + l*l)
			// 对于 [l,r] 中的整数 v，floor(2n / (u^2 + v^2)) 都等于 num
			r = min(int(math.Sqrt(float64(n*2/num-u*u))), u-1) // 推导过程见题解
			// 只有与 u 互质的奇数 v 才能得到本原勾股数组
			numCoprimeOddV := countCoprimeOdd(r, u) - countCoprimeOdd(l-1, u)
			ans += num * numCoprimeOddV
		}
	}

	fmt.Println(time.Since(t0))

	return ans * 2 // (a,b,c) 和 (b,a,c) 各算一次
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	fmt.Println(countTriples(1e9))
}
