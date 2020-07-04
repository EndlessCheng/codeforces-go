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

异或和相关
http://oeis.org/A003987 异或矩阵
http://oeis.org/A003815 异或和 i  a(0)=0, a(4n+1)=1, a(4n+2)=4n+3, a(4n+3)=0, a(4n+4)=4n+4
http://oeis.org/A145768 异或和 i*i
http://oeis.org/A126084 异或和 质数
http://oeis.org/A018252 异或和 合数?
http://oeis.org/A072594 异或和 质因数分解 是积性函数 a(p^k)=p*(k&1)
	http://oeis.org/A072595 满足 A072594(n)=0 的数
http://oeis.org/A178910 异或和 因子
	http://oeis.org/A178911 满足 A178910(n)=n 的数 Perfex number

二进制长度
https://oeis.org/A070939 a(0)=1, a(n)=bits.Len(n)
https://oeis.org/A083652 A070939 的前缀和

OnesCount 相当于二进制的 digsum
https://oeis.org/A000120 wt(n) = OnesCount(n)
https://oeis.org/A000788 A000120 的前缀和 a(2^n)=n*2^(n-1)+1
https://oeis.org/A092391 n+OnesCount(n)
	https://oeis.org/A010061 二进制自我数/哥伦比亚数（A092391 的补集）
https://oeis.org/A011371 n-OnesCount(n) Also highest power of 2 dividing n!
							a(n) = floor(n/2) + a(floor(n/2))
https://oeis.org/A245788 n*OnesCount(n)
https://oeis.org/A049445 OnesCount(n)|n
	-  n/OnesCount(n)
https://oeis.org/A199238 n%OnesCount(n)
https://oeis.org/A010062 a(0)=1, a(n+1)=a(n)+OnesCount(a(n))
	https://oeis.org/A096303 从 n 出发不断执行 n+=OnesCount(n)，直到 n 在 A010062 中，所需要的迭代次数
	Number of iterations of n -> n + (number of 1's in binary representation of n) needed for the trajectory of n to join the trajectory of A010062
		https://oeis.org/A229743 Positions of records
		https://oeis.org/A229744 Values of records
	相关题目 https://www.luogu.com.cn/problem/P5891 https://class.luogu.com.cn/classroom/lgr66

https://oeis.org/A023416 Number of 0's in binary expansion of n
							a(n) = a(n/2) + 1 - n&1
https://oeis.org/A059015 A023416 的前缀和

十进制 digsum
https://oeis.org/A007953 digsum(n)
https://oeis.org/A062028 n+digsum(n)    质数 https://oeis.org/A047791    合数 https://oeis.org/A107743
	https://oeis.org/A003052 自我数/哥伦比亚数 Self number / Colombian number
	https://en.wikipedia.org/wiki/Self_number
	1, 3, 5, 7, 9, 20, 31, 42, 53, 64, 75, 86, 97, 108, ...
		https://oeis.org/A006378 自我质数 Self primes
https://oeis.org/A066568 n-digsum(n)
https://oeis.org/A057147 n*digsum(n)
https://oeis.org/A005349 digsum(n)|n   Niven (or Harshad) numbers
	https://oeis.org/A065877 digsum(n)∤n   Non-Niven (or non-Harshad) numbers
	https://oeis.org/A001101 Moran numbers: n such that (n / digsum(n)) is prime
https://oeis.org/A016052 a(1)=3, a(n+1)=a(n)+digsum(a(n))
https://oeis.org/A051885 Smallest number whose digsum = n
							int64(n%9+1) * int64(math.Pow10(n/9)) - 1
							相关题目 https://codeforces.com/contest/1373/problem/E
https://oeis.org/A118137 digsum(n)+digsum(n+1)

*/

// 参考 strings/strings.go 中的 asciiSet
type bitset []uint32 // b := make(bitset, n>>5+1)

func (b bitset) set(p int)           { b[p>>5] |= 1 << (p & 31) }
func (b bitset) reset(p int)         { b[p>>5] &^= 1 << (p & 31) }
func (b bitset) flip(p int)          { b[p>>5] ^= 1 << (p & 31) }
func (b bitset) contains(p int) bool { return 1<<(p&31)&b[p>>5] > 0 }

func bitsCollection() {
	// ^n+1 = (-1-n)+1 = -n
	lowbit := func(n int64) int64 { return n & -n }

	isPow2 := func(x int64) bool { return x > 0 || x&(x-1) == 0 }

	bits31 := func(n int) []byte {
		bits := make([]byte, 31)
		for i := range bits {
			bits[i] = byte(n >> (30 - i) & 1)
		}
		return bits
	}
	_bits31 := func(n int) string { return Sprintf("%031b", n) }
	_bits32 := func(n uint) string { return Sprintf("%032b", n) }

	digitSum := func(v int) (s int) {
		for ; v > 0; v /= 10 {
			s += v % 10
		}
		return
	}

	_ = []interface{}{lowbit, isPow2, bits31, _bits31, _bits32, digitSum}
}

// https://halfrost.com/go_s2_de_bruijn/

// LC137 https://leetcode-cn.com/problems/single-number-ii/
// 除了某个元素只出现一次以外，其余每个元素均出现了三次。返回只出现了一次的元素
// 		定义两个集合 ones 和 twos，初始为空
// 		第一次出现就放在 ones 中
//		第二次出现就在 ones 中删除并放在 twos
//		第三次出现就从 twos 中删除
//		这样最终 ones 中就留下了最后的结果
func singleNumber(a []int) int {
	ones, twos := 0, 0
	for _, v := range a {
		ones = (ones ^ v) &^ twos
		twos = (twos ^ v) &^ ones
	}
	return ones
}
