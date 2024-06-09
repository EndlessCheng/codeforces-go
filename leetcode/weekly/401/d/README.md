对于 $\textit{rewardValues}$ 中的数，如果先选大的，就没法再选小的，所以**按照从小到大的顺序选**是最好的。

把 $\textit{rewardValues}$ 从小到大排序。

排序后，问题变成一个标准的 0-1 背包问题，请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

对于本题，定义 $f[i][j]$ 表示能否从前 $i$ 个数中得到总奖励 $j$。

考虑 $v=\textit{rewardValues}[i]$ 选或不选：

- 不选 $v$：$f[i][j] = f[i-1][j]$。
- 选 $v$：$f[i][j] = f[i-1][j-v]$，前提是 $j\ge v$ 且 $j-v < v$，即 $v\le j<2v$。

满足其一即可，得

$$
f[i][j] = f[i-1][j] \vee f[i-1][j-v]
$$

其中 $\vee$ 即编程语言中的 `||`。

初始值 $f[0][0] = \texttt{true}$。

答案为最大的满足 $f[n][j]=\texttt{true}$ 的 $j$。

这可以解决周赛第三题，但第四题范围很大，需要去掉第一个维度，并用 **bitset** 优化。

用一个二进制数 $f$ 保存状态，二进制从低到高第 $j$ 位为 $1$ 表示 $f[j]=\texttt{true}$，为 $0$ 表示 $f[j]=\texttt{false}$。

对于上面的转移方程 $f[i][j] = f[i-1][j] \vee f[i-1][j-v]$，其中 $0\le j-v < v$，相当于取 $f$ 的低 $v$ 位，再左移 $v$ 位，然后与 $f$ 取按位或。

初始值 $f=1$。

答案为 $f$ 的最高位，即 $f$ 的二进制长度减一。

代码实现时，可以把 $\textit{rewardValues}$ 中的重复数字去掉，毕竟无法选两个一样的数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1h7421R78s/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maxTotalReward(self, rewardValues: List[int]) -> int:
        f = 1
        for v in sorted(set(rewardValues)):
            f |= (f & ((1 << v) - 1)) << v
        return f.bit_length() - 1
```

```java [sol-Java]
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm/w)$，其中 $n$ 是 $\textit{rewardValues}$ 的长度，$m=\max(\textit{rewardValues})$，$w=64$ 或 $32$。
- 空间复杂度：$\mathcal{O}(m/w)$。

#### 相似题目

- [1981. 最小化目标值与所选元素的差](https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/) 也可以用 bitset 解决。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
