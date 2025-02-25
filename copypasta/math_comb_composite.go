package copypasta

// 第一种写法，分解 mod + 欧拉定理
// https://leetcode.cn/problems/check-if-digits-are-equal-in-string-after-operations-ii/
type combMod struct {
	mod    int     // 如果 mod 是个常量，可以改成 const，加快效率
	fac    []int   // fac[i] = i! 去掉所有 primes 因子的结果
	invF   []int   // invF[i] = fac[i]^-1
	primes []int   // mod 的质因数分解
	preE   [][]int // preE[j][i] = i! 的质因子 primes[j] 的个数
	// 注：如果 preE 消耗的内存太多，可以改用勒让德定理，在 comb 中直接计算 n!、k! 和 (n-k)! 的质因子 primes[j] 的个数
}

func (cm *combMod) pow(x, n int) int {
	mod := cm.mod
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func newCombMod(mx, mod int) *combMod {
	// 质因数分解 & 计算欧拉函数 phi(mod)
	n := mod
	primes := []int{}
	phi := n
	for i := 2; i*i <= n; i++ {
		if n%i > 0 {
			continue
		}
		for n /= i; n%i == 0; n /= i {
		}
		primes = append(primes, i)
		phi = phi / i * (i - 1)
	}
	if n > 1 {
		primes = append(primes, n)
		phi = phi / n * (n - 1)
	}

	preE := make([][]int, len(primes))
	for i := range preE {
		preE[i] = make([]int, mx+1)
	}
	invF := make([]int, mx+1)
	fac := make([]int, mx+1)
	fac[0] = 1
	for i := 1; i <= mx; i++ {
		x := i
		for j, p := range primes {
			e := 0
			for x%p == 0 {
				e++
				x /= p
			}
			preE[j][i] = preE[j][i-1] + e
		}
		fac[i] = fac[i-1] * x % mod
		// 小技巧：把 i 去掉 p 的剩余结果 x 保存在 invF[i-1] 中，这样下面无需重新计算 x
		invF[i-1] = x
	}

	cm := &combMod{mod: mod}
	invF[mx] = cm.pow(fac[mx], phi-1) // 欧拉定理
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i-1] * invF[i] % mod
	}

	cm.fac = fac
	cm.invF = invF
	cm.primes = primes
	cm.preE = preE
	return cm
}

func (cm *combMod) comb(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	mod := cm.mod
	res := cm.fac[n] * cm.invF[k] % mod * cm.invF[n-k] % mod
	for i, pp := range cm.preE {
		res = res * cm.pow(cm.primes[i], pp[n]-pp[k]-pp[n-k]) % mod
	}
	return res
}

// 第二种写法，Lucas 定理 + 中国剩余定理 todo
// comb(n, k) 中 n 和 k 的范围可以更大
