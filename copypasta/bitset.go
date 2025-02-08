package copypasta

import (
	. "fmt"
	"math"
	"math/bits"
	"strconv"
	"strings"
)

// Bitset
// 有时候也可以用 big.Int 代替
// 部分参考 C++ 的标准库源码 https://gcc.gnu.org/onlinedocs/libstdc++/libstdc++-html-USERS-3.4/bitset-source.html
// NOTE: 若要求方法内不修改 b 而是返回一个修改后的拷贝，可以在方法开头加上 b = slices.Clone(b) 并返回 b
// NOTE: 如果效率不够高，可以试试 0-1 线段树，见 segment_tree01.go
//
// https://codeforces.com/problemset/problem/33/D（也可以用 LCA）
// https://codeforces.com/contest/1826/problem/E
// https://atcoder.jp/contests/abc258/tasks/abc258_g
// https://atcoder.jp/contests/arc087/tasks/arc087_b
const _w = bits.UintSize

func NewBitset(n int) Bitset { return make(Bitset, (n+_w-1)/_w) } // 需要 ceil(n/_w) 个 _w 位整数

type Bitset []uint

func (b Bitset) Has(p int) bool { return b[p/_w]&(1<<(p%_w)) != 0 } // get / test
func (b Bitset) Set(p int)      { b[p/_w] |= 1 << (p % _w) }        // 置 1
func (b Bitset) Reset(p int)    { b[p/_w] &^= 1 << (p % _w) }       // 置 0
func (b Bitset) Flip(p int)     { b[p/_w] ^= 1 << (p % _w) }        // 翻转

func (b Bitset) SetAll1() {
	for i := range b {
		b[i] = math.MaxUint
	}
}

// 遍历所有 1 的位置
// 如果对范围有要求，可在 f 中 return p < n
func (b Bitset) Foreach(f func(p int) (Break bool)) {
	for i, v := range b {
		for ; v > 0; v &= v - 1 {
			j := i*_w | bits.TrailingZeros(v)
			if f(j) {
				return
			}
		}
	}
}

// 返回第一个 0 的下标，若不存在则返回一个不小于 n 的位置
func (b Bitset) Index0() int {
	for i, v := range b {
		if ^v != 0 {
			return i*_w | bits.TrailingZeros(^v)
		}
	}
	return len(b) * _w
}

// 返回第一个 1 的下标，若不存在则返回一个不小于 n 的位置（同 C++ 中的 _Find_first）
func (b Bitset) Index1() int {
	for i, v := range b {
		if v != 0 {
			return i*_w | bits.TrailingZeros(v)
		}
	}
	return len(b) * _w
}

// 返回下标 >= p 的第一个 1 的下标，若不存在则返回一个不小于 n 的位置（类似 C++ 中的 _Find_next，这里是 >=）
func (b Bitset) Next1(p int) int {
	if i := p / _w; i < len(b) {
		v := b[i] & (^uint(0) << (p % _w)) // mask off bits below bound
		if v != 0 {
			return i*_w | bits.TrailingZeros(v)
		}
		for i++; i < len(b); i++ {
			if b[i] != 0 {
				return i*_w | bits.TrailingZeros(b[i])
			}
		}
	}
	return len(b) * _w
}

// 返回下标 >= p 的第一个 0 的下标，若不存在则返回一个不小于 n 的位置
func (b Bitset) Next0(p int) int {
	if i := p / _w; i < len(b) {
		v := b[i]
		if p%_w > 0 {
			v |= ^(^uint(0) << (p % _w))
		}
		if ^v != 0 {
			return i*_w | bits.TrailingZeros(^v)
		}
		for i++; i < len(b); i++ {
			if ^b[i] != 0 {
				return i*_w | bits.TrailingZeros(^b[i])
			}
		}
	}
	return len(b) * _w
}

// 返回最后第一个 1 的下标，若不存在则返回 -1
// 注意 Lsh 后超出 n 的 1
func (b Bitset) LastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*_w | (bits.Len(b[i]) - 1) // 如果再 +1，需要改成 i*_w + bits.Len(b[i])
		}
	}
	return -1
}

// += 1 << i，模拟进位
func (b Bitset) Add(i int) { b.FlipRange(i, b.Next0(i)+1) }

// -= 1 << i，模拟借位
func (b Bitset) Sub(i int) { b.FlipRange(i, b.Next1(i)+1) }

// 判断 [l,r] 范围内的数是否全为 0
// https://codeforces.com/contest/1107/problem/D（标准做法是二维前缀和）
func (b Bitset) All0(l, r int) bool {
	i := l / _w
	if i == r/_w {
		mask := ^uint(0)<<(l%_w) ^ ^uint(0)<<(r%_w)
		return b[i]&mask == 0
	}
	if b[i]>>(l%_w) != 0 {
		return false
	}
	for i++; i < r/_w; i++ {
		if b[i] != 0 {
			return false
		}
	}
	mask := ^uint(0) << (r % _w)
	return b[r/_w]&^mask == 0
}

// 判断 [l,r] 范围内的数是否全为 1
func (b Bitset) All1(l, r int) bool {
	i := l / _w
	if i == r/_w {
		mask := ^uint(0)<<(l%_w) ^ ^uint(0)<<(r%_w)
		return b[i]&mask == mask
	}
	mask := ^uint(0) << (l % _w)
	if b[i]&mask != mask {
		return false
	}
	for i++; i < r/_w; i++ {
		if ^b[i] != 0 {
			return false
		}
	}
	mask = ^uint(0) << (r % _w)
	return ^(b[r/_w] | mask) == 0
}

// 反转 [l,r) 范围内的比特
// https://codeforces.com/contest/1705/problem/E
func (b Bitset) FlipRange(l, r int) {
	maskL := ^uint(0) << (l % _w)
	maskR := ^uint(0) << (r % _w)
	i := l / _w
	if i == r/_w {
		b[i] ^= maskL ^ maskR
		return
	}
	b[i] ^= maskL
	for i++; i < r/_w; i++ {
		b[i] = ^b[i]
	}
	b[i] ^= ^maskR
}

// 将 [l,r) 范围内的比特全部置 1
func (b Bitset) SetRange(l, r int) {
	maskL := ^uint(0) << (l % _w)
	maskR := ^uint(0) << (r % _w)
	i := l / _w
	if i == r/_w {
		b[i] |= maskL ^ maskR
		return
	}
	b[i] |= maskL
	for i++; i < r/_w; i++ {
		b[i] = ^uint(0)
	}
	b[i] |= ^maskR
}

// 将 [l,r) 范围内的比特全部置 0
func (b Bitset) ResetRange(l, r int) {
	maskL := ^uint(0) << (l % _w)
	maskR := ^uint(0) << (r % _w)
	i := l / _w
	if i == r/_w {
		b[i] &= ^maskL | maskR
		return
	}
	b[i] &= ^maskL
	for i++; i < r/_w; i++ {
		b[i] = 0
	}
	b[i] &= maskR
}

// 将 >= start 的比特全部置 0
func (b Bitset) ResetFrom(start int) {
	i := start / _w
	b[i] &= ^(^uint(0) << (start % _w))
	clear(b[i+1:])
}

// 左移 k 位
// 注意左移后，超出 n 的 1
// LC1981 https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/
// LC3181 https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/ 
func (b Bitset) Lsh(k int) {
	if k == 0 {
		return
	}
	shift, offset := k/_w, k%_w
	if shift >= len(b) {
		clear(b)
		return
	}
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(_w-offset)
		}
		b[shift] = b[0] << offset
	}
	clear(b[:shift]) // 低位补 0
}

// 右移 k 位
func (b Bitset) Rsh(k int) {
	if k == 0 {
		return
	}
	shift, offset := k/_w, k%_w
	if shift >= len(b) {
		clear(b)
		return
	}
	lim := len(b) - 1 - shift
	if offset == 0 {
		// Fast path
		copy(b, b[shift:])
	} else {
		for i := 0; i < lim; i++ {
			b[i] = b[i+shift]>>offset | b[i+shift+1]<<(_w-offset)
		}
		// 注意：若前后调用 lsh 和 rsh，需要注意超出 n 的范围的 1 对结果的影响（如果需要，可以把范围开大点）
		b[lim] = b[len(b)-1] >> offset
	}
	clear(b[lim+1:]) // 高位补 0
}

// 下面几个方法均需保证长度相同
func (b Bitset) Or(c Bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b Bitset) And(c Bitset) {
	for i, v := range c {
		b[i] &= v
	}
}

func (b Bitset) Xor(c Bitset) {
	for i, v := range c {
		b[i] ^= v
	}
}

func (b Bitset) Equals(c Bitset) bool {
	for i, v := range b {
		if v != c[i] {
			return false
		}
	}
	return true
}

func (b Bitset) HasSubset(c Bitset) bool {
	for i, v := range b {
		if v|c[i] != v {
			return false
		}
	}
	return true
}

// 借用 bits 库中的一些方法的名字
func (b Bitset) OnesCount() (c int) {
	for _, v := range b {
		c += bits.OnesCount(v)
	}
	return
}
func (b Bitset) TrailingZeros() int { return b.Index1() }
func (b Bitset) Len() int           { return b.LastIndex1() + 1 }

// 返回所有是 1 的位置
func (b Bitset) AllIndex1() (idx1 []int) {
	for i, v := range b {
		for ; v > 0; v &= v - 1 {
			j := i*_w | bits.TrailingZeros(v)
			idx1 = append(idx1, j)
		}
	}
	return
}

func (b Bitset) String() string {
	bin := &strings.Builder{}
	for i := len(b) - 1; i >= 0; i-- {
		v := b[i]
		if bin.Len() == 0 {
			if v == 0 {
				continue
			}
			bin.WriteString(strconv.FormatUint(uint64(v), 2))
		} else {
			bin.WriteString(Sprintf("%0*b", _w, v))
		}
	}
	return bin.String()
}
