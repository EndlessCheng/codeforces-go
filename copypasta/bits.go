package copypasta

import _ "math/bits"

/*
标准库 "math/bits" 包含了部分位运算需要的函数，如二进制中 1 的个数、二进制表示的长度等
注意：bits.Len(0) 返回的是 0 而不是 1

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
*/
func bitsCollection() {
	// ^n+1 = (-1-n)+1 = -n
	lowbit := func(n int64) int64 { return n & -n }

	// 也可以用 strconv.FormatInt(n, 2) + 填充前导零 bits.LeadingZeros 来做，注意 n=0 的情况
	bits32 := func(n int) []byte {
		bits := make([]byte, 32)
		for i := range bits {
			bits[i] = byte(n >> uint(31-i) & 1)
		}
		return bits
	}

	_ = []interface{}{lowbit, bits32}
}

// https://halfrost.com/go_s2_de_bruijn/
