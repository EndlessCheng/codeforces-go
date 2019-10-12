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

	_ = []interface{}{maxPow2}
}
