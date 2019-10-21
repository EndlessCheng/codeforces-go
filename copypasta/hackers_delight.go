package copypasta

/*
注意 11100001 的情况（特判 lowbit = 1）

Precedence    Operator
    5         *  /  %  <<  >>  &  &^
    4         +  -  |  ^
    3         ==  !=  <  <=  >  >=
    2         &&
    1         ||

a|b = a^b + a&b
*/

func bitsCollection() {
	// ^n+1 = (-1-n)+1 = -n
	lowbit := func(n int64) int64 { return n & -n }

	maxPow2 := func(n int64) int64 {
		if n == 0 {
			return 0
		}
		for lb := lowbit(n); n != lb; lb = lowbit(n) {
			n -= lb
		}
		return n
	}

	bitLength := func(n int) int {
		c := 1
		if n>>16 > 0 {
			c += 16
			n >>= 16
		}
		if n>>8 > 0 {
			c += 8
			n >>= 8
		}
		if n>>4 > 0 {
			c += 4
			n >>= 4
		}
		if n>>2 > 0 {
			c += 2
			n >>= 2
		}
		if n-1 > 0 {
			c++
		}
		return c
	}

	_ = []interface{}{maxPow2, bitLength}
}
