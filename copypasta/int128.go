package copypasta

import "math/bits"

// Add: (aHi,aLo) + (bHi,bLo) -> (hi, lo)
func Add(aHi, aLo, bHi, bLo uint64) (hi, lo uint64) {
	lo, carry := bits.Add64(aLo, bLo, 0)
	hi, _ = bits.Add64(aHi, bHi, carry)
	return hi, lo
}

// Mul: (aHi,aLo) * (bHi,bLo) -> 256-bit 结果，这里只返回低 128 位（loHi, loLo）
func MulLow128(aHi, aLo, bHi, bLo uint64) (hi, lo uint64) {
	// 使用 64x64 -> 128 分解：
	// a = aHi<<64 | aLo, b = bHi<<64 | bLo
	// product low 128 bits = aLo*bLo + ((aHi*bLo + aLo*bHi) << 64)
	p0Hi, p0Lo := bits.Mul64(aLo, bLo) // aLo*bLo = p0Hi<<64 | p0Lo
	p1Hi, p1Lo := bits.Mul64(aHi, bLo) // aHi*bLo
	p2Hi, p2Lo := bits.Mul64(aLo, bHi) // aLo*bHi

	// accumulate the middle terms into high/low
	midLo, carry1 := bits.Add64(p1Lo, p2Lo, 0)
	midHi := p1Hi + p2Hi + carry1

	// now add p0Hi (carry from low 64x64) to midLo (shifted by 64)
	hi, _ = bits.Add64(p0Hi, midLo, 0) // hi is high 64 of the 128-bit product
	lo = p0Lo                          // low 64
	hi += midHi                        // complete high 64
	return hi, lo
}
