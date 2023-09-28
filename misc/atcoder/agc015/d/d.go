package main
import(."fmt";"math/bits")

// https://space.bilibili.com/206214
func main() {
	var low, high uint
	Scan(&low, &high)
	if low == high {
		Print(1)
		return
	}
	ans := high - low + 1
	mask := uint(1)<<(bits.Len(high^low)-1) - 1
	high &= mask
	low &= mask
	nh := bits.Len(high)
	if bits.Len(low) <= nh {
		ans += mask - high
	} else {
		ans += mask - low + 1<<nh - high
	}
	Print(ans)
}
