package copypasta

import (
	. "fmt"
	_ "math/bits"
)

/*
标准库 "math/bits" 包含了位运算常用的函数，如二进制中 1 的个数、二进制表示的长度等
注意：bits.Len(0) 返回的是 0 而不是 1
     bits.Len(x) 相当于 int(Log2(x)+eps)+1

运算符优先级 https://golang.org/ref/spec#Operators
Precedence    Operator
    5         *  /  %  <<  >>  &  &^
    4         +  -  |  ^
    3         ==  !=  <  <=  >  >=
    2         &&
    1         ||

一些子集的枚举算法见 loopCollection
S∪{i}: S|1<<i
S\{i}:  S&^(1<<i)
构造 2^n-1，即 n 个 1 的另一种方法: ^(-1<<n)
a|b = a^b + a&b

https://oeis.org/A070939 a(0)=1, a(i)=bits.Len(i)
https://oeis.org/A083652 A070939 的前缀和

https://oeis.org/A000120 OnesCount(i)
https://oeis.org/A000788 A000120 的前缀和 a(2^n)=n*2^(n-1)+1

	https://oeis.org/A023416 Number of 0's in binary expansion of n
	a(n) = a(n/2) + 1 - n&1
	https://oeis.org/A059015 A023416 的前缀和

https://oeis.org/A010061 二进制自我数/哥伦比亚数 numbers not of form m + sum of binary digits of m
https://oeis.org/A010062 a(0)=1, a(n+1)=a(n)+OnesCount(a(n))

	https://oeis.org/A096303 从 n 出发不断执行 n+=OnesCount(n)，直到 n 在 A010062 中，所需要的迭代次数
	Number of iterations of n -> n + (number of 1's in binary representation of n) needed for the trajectory of n to join the trajectory of A010062
		https://oeis.org/A229743 Positions of records
		https://oeis.org/A229744 Values of records

	相关题目 https://www.luogu.com.cn/problem/P5891 https://class.luogu.com.cn/classroom/lgr66

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
