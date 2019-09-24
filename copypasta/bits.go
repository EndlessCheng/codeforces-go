package copypasta

/*
注意 11100001 的情况
*/
func bitsCollection() {
	maxPow2 := func(n int64) int64 {
		if n == 0 {
			return 0
		}
		msb := uint(0)
		for n >>= 1; n != 0; n >>= 1 {
			msb++
		}
		return int64(1) << msb
	}

	_ = []interface{}{maxPow2}
}
