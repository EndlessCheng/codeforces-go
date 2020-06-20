package copypasta

import (
	. "fmt"
	_ "math/bits"
)

/*
标准库 "math/bits" 包含了部分位运算需要的函数，如二进制中 1 的个数、二进制表示的长度等
注意：bits.Len(0) 返回的是 0 而不是 1
注意：bits.Len(x) 相当于 int(log_2(x))-1

https://oeis.org/A070939 Length of binary representation of n
https://oeis.org/A083652 前缀和

运算符优先级 (https://golang.org/ref/spec#Operators)
Precedence    Operator
    5         *  /  %  <<  >>  &  &^
    4         +  -  |  ^
    3         ==  !=  <  <=  >  >=
    2         &&
    1         ||

a|b = a^b + a&b

S∪{i}: S|1<<i
S\{i}:  S&^(1<<i)

注意 11100001 的情况（特判 lowbit = 1）

构造 2^n-1，即 n 个 1：^(-1<<n)

一些子集的枚举算法见 loopCollection
*/

// 参考 strings/strings.go 中的 asciiSet
type bitset []uint32 // b := make(bitset, n>>5+1)

func (b bitset) set(c int)           { b[c>>5] |= 1 << (c & 31) }
func (b bitset) contains(c int) bool { return 1<<(c&31)&b[c>>5] > 0 }

func bitsCollection() {
	// ^n+1 = (-1-n)+1 = -n
	lowbit := func(n int64) int64 { return n & -n }

	bits31 := func(n int) []byte {
		bits := make([]byte, 31)
		for i := range bits {
			bits[i] = byte(n >> (30 - i) & 1)
		}
		return bits
	}
	_bits31 := func(n int) string { return Sprintf("%031b", n) }
	_bits32 := func(n uint) string { return Sprintf("%032b", n) }

	_ = []interface{}{lowbit, bits31, _bits31, _bits32}
}

// https://halfrost.com/go_s2_de_bruijn/
