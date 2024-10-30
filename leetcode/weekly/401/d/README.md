## 思路

对于 $\textit{rewardValues}$ 中的数，如果先选大的，就没法再选小的，所以**按照从小到大的顺序选**是最优的。

把 $\textit{rewardValues}$ 从小到大排序。

排序后，问题类似 0-1 背包，原理请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

定义 $f[i][j]$ 表示能否从 $\textit{rewardValues}$ 的前 $i$ 个数中得到总奖励 $j$。

设 $\textit{rewardValues}$ 的第 $i$ 个数为 $v$，考虑 $v$ 选或不选：

- 不选 $v$，问题变成能否从前 $i-1$ 个数中得到总奖励 $j$，即 $f[i][j] = f[i-1][j]$。
- 选 $v$，问题变成能否从前 $i-1$ 个数中得到总奖励 $j-v$，即 $f[i][j] = f[i-1][j-v]$，前提是 $j$ 满足 $j\ge v$ 且 $j-v < v$，即 $v\le j<2v$。

选或不选满足其一即可，所以有

$$
f[i][j] = f[i-1][j] \vee f[i-1][j-v]
$$

其中 $\vee$ 即编程语言中的 `||`。

初始值 $f[0][0] = \texttt{true}$。

答案为最大的满足 $f[n][j]=\texttt{true}$ 的 $j$。

$j$ 最大枚举到 $2m-1$（见下面答疑），其中 $m$ 是数组的最大值。

这样可以解决 [3180. 执行操作可获得的最大总奖励 I](https://leetcode.cn/problems/maximum-total-reward-using-operations-i/)，但本题数据范围更大，需要去掉第一个维度，并用 **bitset** 优化（也可以用 **bigint**）。

#### 答疑

**问**：为什么 $2m-1$ 是答案的上界？

**答**：如果最后一步选的数是 $x$，而 $x<m$，那么把 $x$ 替换成 $m$ 也符合要求，矛盾，所以最后一步选的一定是 $m$。在选 $m$ 之前，元素和至多是 $m-1$，选了 $m$ 之后，元素和至多是 $2m-1$。我们无法得到比 $2m-1$ 更大的元素和。

## 优化一

上面的转移方程，先去掉第一个维度，得到 $f[j] = f[j] \vee f[j-v]$，其中 $v\le j<2v$。

进一步地，把一维数组压缩成一个二进制数 $f$，其中二进制从低到高第 $j$ 位为 $1$ 表示 $f[j]=\texttt{true}$，为 $0$ 表示 $f[j]=\texttt{false}$。转换后 $\vee$ 就是或运算（OR）了。

比如 $v=3$，我们会把：

- $f[0]$ OR 到 $f[3]$ 中；
- $f[1]$ OR 到 $f[4]$ 中；
- $f[2]$ OR 到 $f[5]$ 中。

这相当于取 $f$ 的低 $v$ 位，再左移 $v$ 位，然后 OR 到 $f$ 中，即编程语言中的 `f |= (f & ((1 << v) - 1)) << v`，具体来说：

- `(1 << v) - 1` 会得到一个低 $v$ 位全为 $1$ 的二进制数。
- `f & ((1 << v) - 1)` 得到 $f$ 的低 $v$ 位。
- `(f & ((1 << v) - 1)) << v` 把 `f & ((1 << v) - 1)` 左移 $v$ 位。
- 最后与 $f$ 计算按位或，更新到 $f$ 中，即 `f |= (f & ((1 << v) - 1)) << v`。

初始值 $f=1$。

答案为 $f$ 的最高位，即 $f$ 的二进制长度减一。

小优化：代码实现时，可以先把 $\textit{rewardValues}$ 中的重复数字去掉再 DP，毕竟无法选两个一样的数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1h7421R78s/?t=14m08s)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maxTotalReward(self, rewardValues: List[int]) -> int:
        f = 1
        for v in sorted(set(rewardValues)):
            f |= (f & ((1 << v) - 1)) << v
        return f.bit_length() - 1
```

```java [sol-Java]
// 超时，请阅读下面的优化二和优化三
import java.math.BigInteger;

class Solution {
    public int maxTotalReward(int[] rewardValues) {
        BigInteger f = BigInteger.ONE;
        for (int v : Arrays.stream(rewardValues).distinct().sorted().toArray()) {
            BigInteger mask = BigInteger.ONE.shiftLeft(v).subtract(BigInteger.ONE);
            f = f.or(f.and(mask).shiftLeft(v));
        }
        return f.bitLength() - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxTotalReward(vector<int>& rewardValues) {
        ranges::sort(rewardValues);
        rewardValues.erase(unique(rewardValues.begin(), rewardValues.end()), rewardValues.end());

        bitset<100000> f{1};
        for (int v : rewardValues) {
            int shift = f.size() - v;
            // 左移 shift 再右移 shift，把所有 >= v 的比特位置 0
            // f |= f << shift >> shift << v;
            f |= f << shift >> (shift - v); // 简化上式
        }
        for (int i = rewardValues.back() * 2 - 1; ; i--) {
            if (f.test(i)) {
                return i;
            }
        }
    }
};
```

```go [sol-Go]
func maxTotalReward(rewardValues []int) int {
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重

	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)
	for _, v := range rewardValues {
		mask := p.Sub(p.Lsh(one, uint(v)), one)
		f.Or(f, p.Lsh(p.And(f, mask), uint(v)))
	}
	return f.BitLen() - 1
}
```

```go [sol-Go bitset]
const w = bits.UintSize

type bitset []uint

// b <<= k
func (b bitset) lsh(k int) bitset {
	shift, offset := k/w, k%w
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(w-offset)
		}
		b[shift] = b[0] << offset
	}
	clear(b[:shift])
	return b
}

// 把 >= start 的清零
func (b bitset) resetRange(start int) bitset {
	i := start / w
	b[i] &= ^(^uint(0) << (start % w))
	clear(b[i+1:])
	return b
}

// b |= c
func (b bitset) unionFrom(c bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b bitset) lastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*w | (bits.Len(b[i]) - 1)
		}
	}
	return -1
}

func maxTotalReward(rewardValues []int) int {
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重

	m := rewardValues[len(rewardValues)-1]
	f := make(bitset, m*2/w+1)
	f[0] = 1
	for _, v := range rewardValues {
		f.unionFrom(slices.Clone(f).lsh(v).resetRange(v * 2))
	}
	return f.lastIndex1()
}
```

## 优化二

设 $m=\max(\textit{rewardValues})$，如果数组中包含 $m-1$，则答案为 $2m-1$，无需计算 DP。

```py [sol-Python3]
class Solution:
    def maxTotalReward(self, rewardValues: List[int]) -> int:
        m = max(rewardValues)
        if m - 1 in rewardValues:
            return m * 2 - 1

        f = 1
        for v in sorted(set(rewardValues)):
            f |= (f & ((1 << v) - 1)) << v
        return f.bit_length() - 1
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    public int maxTotalReward(int[] rewardValues) {
        int m = 0;
        for (int v : rewardValues) {
            m = Math.max(m, v);
        }
        for (int v : rewardValues) {
            if (v == m - 1) {
                return m * 2 - 1;
            }
        }

        BigInteger f = BigInteger.ONE;
        for (int v : Arrays.stream(rewardValues).distinct().sorted().toArray()) {
            BigInteger mask = BigInteger.ONE.shiftLeft(v).subtract(BigInteger.ONE);
            f = f.or(f.and(mask).shiftLeft(v));
        }
        return f.bitLength() - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxTotalReward(vector<int>& rewardValues) {
        int m = ranges::max(rewardValues);
        if (ranges::find(rewardValues, m - 1) != rewardValues.end()) {
            return m * 2 - 1;
        }

        ranges::sort(rewardValues);
        rewardValues.erase(unique(rewardValues.begin(), rewardValues.end()), rewardValues.end());
        bitset<100000> f{1};
        for (int v : rewardValues) {
            int shift = f.size() - v;
            // 左移 shift 再右移 shift，把所有 >= v 的比特位置 0
            // f |= f << shift >> shift << v;
            f |= f << shift >> (shift - v); // 简化上式
        }
        for (int i = m * 2 - 1;; i--) {
            if (f.test(i)) {
                return i;
            }
        }
    }
};
```

```go [sol-Go]
func maxTotalReward(rewardValues []int) int {
	m := slices.Max(rewardValues)
	if slices.Contains(rewardValues, m-1) {
		return m*2 - 1
	}

	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重

	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)
	for _, v := range rewardValues {
		mask := p.Sub(p.Lsh(one, uint(v)), one)
		f.Or(f, p.Lsh(p.And(f, mask), uint(v)))
	}
	return f.BitLen() - 1
}
```

```go [sol-Go bitset]
const w = bits.UintSize

type bitset []uint

// b <<= k
func (b bitset) lsh(k int) bitset {
	shift, offset := k/w, k%w
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(w-offset)
		}
		b[shift] = b[0] << offset
	}
	clear(b[:shift])
	return b
}

// 把 >= start 的清零
func (b bitset) resetRange(start int) bitset {
	i := start / w
	b[i] &= ^(^uint(0) << (start % w))
	clear(b[i+1:])
	return b
}

// b |= c
func (b bitset) unionFrom(c bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b bitset) lastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*w | (bits.Len(b[i]) - 1)
		}
	}
	return -1
}

func maxTotalReward(rewardValues []int) int {
	m := slices.Max(rewardValues)
	if slices.Contains(rewardValues, m-1) {
		return m*2 - 1
	}

	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重
	f := make(bitset, m*2/w+1)
	f[0] = 1
	for _, v := range rewardValues {
		f.unionFrom(slices.Clone(f).lsh(v).resetRange(v * 2))
	}
	return f.lastIndex1()
}
```

## 优化三

如果有两个不同元素之和等于 $m-1$，也可以直接返回 $2m-1$。

由于在随机数据下，几乎 100% 的概率有两数之和等于 $m-1$，而力扣又喜欢出随机数据，所以在大多数情况下，本题就是一道 [1. 两数之和](https://leetcode.cn/problems/two-sum/)。

```py [sol-Python3]
class Solution:
    def maxTotalReward(self, rewardValues: List[int]) -> int:
        m = max(rewardValues)
        s = set()
        for v in rewardValues:
            if v in s:
                continue
            if v == m - 1 or m - 1 - v in s:
                return m * 2 - 1
            s.add(v)

        f = 1
        for v in sorted(s):
            f |= (f & ((1 << v) - 1)) << v
        return f.bit_length() - 1
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    public int maxTotalReward(int[] rewardValues) {
        int m = 0;
        for (int v : rewardValues) {
            m = Math.max(m, v);
        }
        Set<Integer> set = new HashSet<>();
        for (int v : rewardValues) {
            if (v == m - 1) {
                return m * 2 - 1;
            }
            if (set.contains(v)) {
                continue;
            }
            if (set.contains(m - 1 - v)) {
                return m * 2 - 1;
            }
            set.add(v);
        }

        Arrays.sort(rewardValues);
        int pre = 0;
        BigInteger f = BigInteger.ONE;
        for (int v : rewardValues) {
            if (v == pre) {
                continue;
            }
            BigInteger mask = BigInteger.ONE.shiftLeft(v).subtract(BigInteger.ONE);
            f = f.or(f.and(mask).shiftLeft(v));
            pre = v;
        }
        return f.bitLength() - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxTotalReward(vector<int>& rewardValues) {
        int m = ranges::max(rewardValues);
        unordered_set<int> s;
        for (int v : rewardValues) {
            if (s.contains(v)) {
                continue;
            }
            if (v == m - 1 || s.contains(m - 1 - v)) {
                return m * 2 - 1;
            }
            s.insert(v);
        }

        ranges::sort(rewardValues);
        rewardValues.erase(unique(rewardValues.begin(), rewardValues.end()), rewardValues.end());

        bitset<100000> f{1};
        for (int v : rewardValues) {
            int shift = f.size() - v;
            // 左移 shift 再右移 shift，把所有 >= v 的比特位置 0
            // f |= f << shift >> shift << v;
            f |= f << shift >> (shift - v); // 简化上式
        }
        for (int i = m * 2 - 1;; i--) {
            if (f.test(i)) {
                return i;
            }
        }
    }
};
```

```go [sol-Go]
func maxTotalReward(rewardValues []int) int {
	m := slices.Max(rewardValues)
	has := map[int]bool{}
	for _, v := range rewardValues {
		if v == m-1 {
			return m*2 - 1
		}
		if has[v] {
			continue
		}
		if has[m-1-v] {
			return m*2 - 1
		}
		has[v] = true
	}

	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重

	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)
	for _, v := range rewardValues {
		mask := p.Sub(p.Lsh(one, uint(v)), one)
		f.Or(f, p.Lsh(p.And(f, mask), uint(v)))
	}
	return f.BitLen() - 1
}
```

```go [sol-Go bitset]
const w = bits.UintSize

type bitset []uint

// b <<= k
func (b bitset) lsh(k int) bitset {
	shift, offset := k/w, k%w
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(w-offset)
		}
		b[shift] = b[0] << offset
	}
	clear(b[:shift])
	return b
}

// 把 >= start 的清零
func (b bitset) resetRange(start int) bitset {
	i := start / w
	b[i] &= ^(^uint(0) << (start % w))
	clear(b[i+1:])
	return b
}

// b |= c
func (b bitset) unionFrom(c bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b bitset) lastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*w | (bits.Len(b[i]) - 1)
		}
	}
	return -1
}

func maxTotalReward(rewardValues []int) int {
	m := slices.Max(rewardValues)
	has := map[int]bool{}
	for _, v := range rewardValues {
		if v == m-1 {
			return m*2 - 1
		}
		if has[v] {
			continue
		}
		if has[m-1-v] {
			return m*2 - 1
		}
		has[v] = true
	}

	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重
	f := make(bitset, m*2/w+1)
	f[0] = 1
	for _, v := range rewardValues {
		f.unionFrom(slices.Clone(f).lsh(v).resetRange(v * 2))
	}
	return f.lastIndex1()
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm/w)$，其中 $n$ 是 $\textit{rewardValues}$ 的长度，$m=\max(\textit{rewardValues})$，$w=64$ 或 $32$。
- 空间复杂度：$\mathcal{O}(n + m/w)$。

## 相似题目

- [1981. 最小化目标值与所选元素的差](https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/) 也可以用 bitset 解决。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
