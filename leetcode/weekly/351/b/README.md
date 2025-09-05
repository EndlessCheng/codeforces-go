## 转化

假设我们操作了 $k$ 次，此时 $\textit{num}_1$ 变成 $\textit{num}_1 - \textit{num}_2\cdot k$ 再减去 $k$ 个 $2^i$。

能否把 $\textit{num}_1$ 变成 $0$，等价于：

- 能否把 $\textit{num}_1 - \textit{num}_2\cdot k$ 拆分成**恰好** $k$ 个 $2$ 的幂之和？

在示例 1 中，$k=3$ 时 $\textit{num}_1 - \textit{num}_2\cdot k = 9$，我们可以把 $9$ 拆分成 $4+4+1$，这三个数都是 $2$ 的幂。

## 上下界分析

设 $x=\textit{num}_1 - \textit{num}_2\cdot k$。

为了判断能否把 $x$ 拆分成恰好 $k$ 个 $2$ 的幂之和，我们可以先做上下界分析：

- 上界：求出 $x$ 最多可以拆分出 $\textit{high}$ 个 $2$ 个幂。
- 下界：求出 $x$ 最少可以拆分出 $\textit{low}$ 个 $2$ 个幂。

由于一个 $2^i$ 可以分解成两个 $2^{i-1}$，而 $2^{i-1}$ 又可以继续分解为 $2^{i-2}$，所以分解出的 $2$ 的幂的个数可以是 $[\textit{low},\textit{high}]$ 中的任意整数。$k$ 只要在这个范围中，那么分解方案就是存在的。

- 上界：由于 $2$ 的幂最小是 $1$，所以 $x$ 最多可以拆分出 $x$ 个 $2$ 个幂（$x$ 个 $1$）。
- 下界：$x$ 的二进制中的 $1$ 的个数。比如 $x$ 的二进制为 $10110$，至少要拆分成 $3$ 个 $2$ 的幂，即 $10000+100+10$。

## 枚举 k

暴力的想法是，从小到大枚举 $k=1,2,3,\ldots$ 计算 $x=\textit{num}_1 - \textit{num}_2\cdot k$，判断 $k$ 是否满足上下界（在区间中）。这样做是否会超时？$k$ 最大枚举到多少呢？

对于上界，即 $k\le x = \textit{num}_1 - \textit{num}_2\cdot k$，变形得 $k\cdot (\textit{num}_2+1)\le \textit{num}_1$。

- 如果 $\textit{num}_2 + 1\le 0$，由于题目保证 $\textit{num}_1\ge 1$，上式恒成立。
- 如果 $\textit{num}_2 + 1> 0$，那么 $k\le \dfrac{\textit{num}_1}{\textit{num}_2+1}$。

对于下界，定义 $\text{popcount}(x)$ 为 $x$ 的二进制中的 $1$ 的个数，我们要满足 $k\ge \text{popcount}(x)$。粗略估计一下，当 $k=60$ 时，在本题数据范围下，当 $\textit{num}_1=10^9$，$\textit{num}_2 = -10^9$ 时 $x$ 最大，为 $61\times 10^9$，二进制长度只有 $36$。由于 $\text{popcount}(x)$ 不会超过 $x$ 的二进制长度，所以此时 $k$ 一定满足下界。所以本题的枚举次数其实很小，暴力枚举不会超时。

综上所述，在枚举 $k=1,2,3,\ldots$ 的过程中：

- 如果 $k > \textit{num}_1 - \textit{num}_2\cdot k$，不满足上界。由于此时 $k > \dfrac{\textit{num}_1}{\textit{num}_2+1}$，对于更大的 $k$，不等式仍然成立，更大 $k$ 也不满足上界。所以可以退出循环，返回 $-1$。
- 否则，如果 $k$ 满足下界，返回 $k$。
- 否则，继续枚举 $k$。

**补充说明**：

- 如果 $\textit{num}_2 + 1\le 0$，由于上界一定满足，下界当 $k$ 大到一定程度就能满足，所以这种情况是一定有解的。
- 如果 $\textit{num}_2 + 1> 0$，可能在 $k$ 枚举到满足下界之前，就已经无法满足上界了，所以可能无解。

[视频讲解](https://www.bilibili.com/video/BV1du41187ZN/)（第二题）

```py [sol-Python3]
class Solution:
    def makeTheIntegerZero(self, num1: int, num2: int) -> int:
        for k in count(1):  # 枚举 k=1,2,3,...
            x = num1 - num2 * k
            if k > x:
                return -1
            if k >= x.bit_count():
                return k
```

```java [sol-Java]
class Solution {
    public int makeTheIntegerZero(int num1, int num2) {
        for (long k = 1; k <= num1 - num2 * k; k++) {
            if (k >= Long.bitCount(num1 - num2 * k)) {
                return (int) k;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int makeTheIntegerZero(int num1, int num2) {
        for (long long k = 1; k <= num1 - num2 * k; k++) {
            if (k >= popcount((uint64_t) num1 - num2 * k)) {
                return k;
            }
        }
        return -1;
    }
};
```

```c [sol-C]
int makeTheIntegerZero(int num1, int num2) {
    for (long long k = 1; k <= num1 - num2 * k; k++) {
        if (k >= __builtin_popcountll(num1 - num2 * k)) {
            return k;
        }
    }
    return -1;
}
```

```go [sol-Go]
func makeTheIntegerZero(num1, num2 int) int {
	for k := 1; k <= num1-num2*k; k++ {
		if k >= bits.OnesCount(uint(num1-num2*k)) {
			return k
		}
	}
	return -1
}
```

```js [sol-JavaScript]
var makeTheIntegerZero = function(num1, num2) {
    for (let k = 1; k <= num1 - num2 * k; k++) {
        if (k >= bitCount64(num1 - num2 * k)) {
            return k;
        }
    }
    return -1;
};

function bitCount64(i) {
    return bitCount32(Math.floor(i / 0x100000000)) + bitCount32(i >>> 0);
}

function bitCount32(i) {
    i = i - ((i >>> 1) & 0x55555555);
    i = (i & 0x33333333) + ((i >>> 2) & 0x33333333);
    i = (i + (i >>> 4)) & 0x0f0f0f0f;
    i = i + (i >>> 8);
    i = i + (i >>> 16);
    return i & 0x3f;
}
```

```rust [sol-Rust]
impl Solution {
    pub fn make_the_integer_zero(num1: i32, num2: i32) -> i32 {
        for k in 1.. {
            let x = num1 as i64 - num2 as i64 * k;
            if k > x {
                return -1;
            }
            if k as u32 >= x.count_ones() {
                return k as _;
            }
        }
        unreachable!()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(f^{-1}(\textit{num}_1+|\textit{num}_2|))$，其中 $f^{-1}(x)$ 是 $f(x)=\dfrac{2^x}{x}$ 的反函数，略大于 $\log_2 x$。在本题的数据范围下，$k\le 36$。
- 空间复杂度：$\mathcal{O}(1)$。

**注**：关于这个反函数的研究，见**朗伯 W 函数**（Lambert W function）。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
