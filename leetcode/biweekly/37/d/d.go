package main

// github.com/EndlessCheng/codeforces-go
const mod = 1_000_000_007

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

type Fancy struct {
	vals []int
	add  int
	mul  int
}

func Constructor() Fancy {
	return Fancy{mul: 1}
}

func (f *Fancy) Append(val int) {
	// 注意这里有减法，计算结果可能是负数，+mod 可以保证计算结果非负
	f.vals = append(f.vals, (val-f.add+mod)*pow(f.mul, mod-2)%mod)
}

func (f *Fancy) AddAll(inc int) {
	f.add = (f.add + inc) % mod
}

func (f *Fancy) MultAll(m int) {
	f.mul = f.mul * m % mod
	f.add = f.add * m % mod
}

func (f *Fancy) GetIndex(idx int) int {
	if idx >= len(f.vals) {
		return -1
	}
	return (f.vals[idx]*f.mul + f.add) % mod
}
