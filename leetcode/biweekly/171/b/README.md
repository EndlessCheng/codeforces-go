暴力做法是枚举范围内的所有回文数，但如果把值域改成 $[1,10^9]$，暴力做法就会超时。怎么办？

## 方法一：预处理

知道了回文数的左半边，就知道了回文数的右半边，所以可以**枚举**回文数的左半边，**从小到大**预处理所有回文数列表 $\textit{pal}$。

然后在 $\textit{pal}$ 中**二分查找**离 $\textit{nums}[i]$ 最近的回文数。关于二分查找的原理，请看[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
def init() -> List[int]:
    MX = 5000

    # 哨兵，0 也是回文数
    palindromes = [0]
    pow2 = 1

    while True:
        # 生成奇数长度回文数
        for i in range(pow2, pow2 * 2):
            s = bin(i)[2:]
            x = int(s + s[::-1][1:], 2)
            if x > MX:
                return palindromes
            palindromes.append(x)

        # 生成偶数长度回文数
        for i in range(pow2, pow2 * 2):
            s = bin(i)[2:]
            x = int(s + s[::-1], 2)
            if x > MX:
                return palindromes
            palindromes.append(x)

        pow2 *= 2

palindromes = init()
# 哨兵，5049 是大于 5000 的第一个二进制回文数
palindromes.append(5049)


class Solution:
    def minOperations(self, nums: List[int]) -> List[int]:
        for i, x in enumerate(nums):
            j = bisect_left(palindromes, x)
            # 变成 >= x 的二进制回文数，或者变成 < x 的二进制回文数
            nums[i] = min(palindromes[j] - x, x - palindromes[j - 1])
        return nums
```

```java [sol-Java]
class Solution {
    private static final List<Integer> palindromes = new ArrayList<>();
    private static boolean initialized = false;

    // 这样写比 static block 快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        final int MX = 5000;
        final int BASE = 2;

        // 哨兵，0 也是回文数
        palindromes.add(0);

        for (int pw = 1; ; pw *= BASE) {
            // 生成奇数长度回文数
            for (int i = pw; i < pw * BASE; i++) {
                int x = i;
                for (int t = i / BASE; t > 0; t /= BASE) {
                    x = x * BASE + t % BASE;
                }
                if (x > MX) {
                    // 哨兵，5049 是大于 5000 的第一个二进制回文数
                    palindromes.add(5049);
                    return;
                }
                palindromes.add(x);
            }

            // 生成偶数长度回文数
            for (int i = pw; i < pw * BASE; i++) {
                int x = i;
                for (int t = i; t > 0; t /= BASE) {
                    x = x * BASE + t % BASE;
                }
                if (x > MX) {
                    palindromes.add(5049);
                    return;
                }
                palindromes.add(x);
            }
        }
    }

    public int[] minOperations(int[] nums) {
        init();
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            int j = Collections.binarySearch(palindromes, x);
            if (j < 0) j = ~j;
            // 变成 >= x 的二进制回文数，或者变成 < x 的二进制回文数
            nums[i] = Math.min(palindromes.get(j) - x, x - palindromes.get(j - 1));
        }
        return nums;
    }
}
```

```cpp [sol-C++]
vector<int> palindromes;

auto init = []() {
    constexpr int MX = 5000;
    constexpr int BASE = 2;

    // 哨兵，0 也是回文数
    palindromes.push_back(0);

    for (int pw = 1;; pw *= BASE) {
        // 生成奇数长度回文数
        for (int i = pw; i < pw * BASE; i++) {
            long long x = i;
            for (int t = i / BASE; t > 0; t /= BASE) {
                x = x * BASE + t % BASE;
            }
            if (x > MX) {
                // 哨兵，5049 是大于 5000 的第一个二进制回文数
                palindromes.push_back(5049);
                return 0;
            }
            palindromes.push_back(x);
        }

        // 生成偶数长度回文数
        for (int i = pw; i < pw * BASE; i++) {
            long long x = i;
            for (int t = i; t > 0; t /= BASE) {
                x = x * BASE + t % BASE;
            }
            if (x > MX) {
                palindromes.push_back(5049);
                return 0;
            }
            palindromes.push_back(x);
        }
    }
}();

class Solution {
public:
    vector<int> minOperations(vector<int>& nums) {
        for (int& x : nums) {
            int j = ranges::lower_bound(palindromes, x) - palindromes.begin();
            // 变成 >= x 的二进制回文数，或者变成 < x 的二进制回文数
            x = min(palindromes[j] - x, x - palindromes[j - 1]);
        }
        return nums;
    }
};
```

```go [sol-Go]
var palindromes []int

// 预处理二进制回文数
func init() {
	const mx = 5000
	const base = 2

	// 哨兵，0 也是回文数
	palindromes = append(palindromes, 0)

outer:
	for pw := 1; ; pw *= base {
		// 生成奇数长度回文数
		for i := pw; i < pw*base; i++ {
			x := i
			for t := i / base; t > 0; t /= base {
				x = x*base + t%base
			}
			if x > mx {
				break outer
			}
			palindromes = append(palindromes, x)
		}

		// 生成偶数长度回文数
		for i := pw; i < pw*base; i++ {
			x := i
			for t := i; t > 0; t /= base {
				x = x*base + t%base
			}
			if x > mx {
				break outer
			}
			palindromes = append(palindromes, x)
		}
	}

	// 哨兵，5049 是大于 5000 的第一个二进制回文数
	palindromes = append(palindromes, 5049)
}

func minOperations(nums []int) []int {
	for i, x := range nums {
		j := sort.SearchInts(palindromes, x)
		// 变成 >= x 的二进制回文数，或者变成 < x 的二进制回文数
		nums[i] = min(palindromes[j]-x, x-palindromes[j-1])
	}
	return nums
}
```

#### 复杂度分析

> 预处理的时空复杂度为 $\mathcal{O}(\sqrt U)$，$U=5000$，不计入。

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=5000$ 是 $\textit{nums}[i]$ 的最大值。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：位运算

设 $x= \textit{nums}[i]$，$x$ 的左半（奇数长度时包含中心）是 $\textit{left}$。

最终变成的回文数的左半在 $\textit{left}-1, \textit{left}, \textit{left}+1$ 中。枚举这三种情况。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> List[int]:
        for i, x in enumerate(nums):
            res = inf
            n = x.bit_length()
            m = n // 2
            left = x >> m
            for l in range(left - 1, left + 2):
                # 左半反转到右半
                right = self.reverseBits(l >> (n % 2)) >> (32 - m)
                pal = l << m | right
                res = min(res, abs(x - pal))
            nums[i] = res
        return nums

    # 190. 颠倒二进制位
    # https://leetcode.cn/problems/reverse-bits/
    def reverseBits(self, n: int) -> int:
        # 交换 16 位
        n = ((n >> 16) | (n << 16)) & 0xFFFFFFFF
        # 交换每个 8 位块
        n = (((n & 0xFF00FF00) >> 8) | ((n & 0x00FF00FF) << 8)) & 0xFFFFFFFF
        # 交换每个 4 位块
        n = (((n & 0xF0F0F0F0) >> 4) | ((n & 0x0F0F0F0F) << 4)) & 0xFFFFFFFF
        # 交换每个 2 位块
        n = (((n & 0xCCCCCCCC) >> 2) | ((n & 0x33333333) << 2)) & 0xFFFFFFFF
        # 交换相邻位
        n = (((n & 0xAAAAAAAA) >> 1) | ((n & 0x55555555) << 1)) & 0xFFFFFFFF
        return n
```

```java [sol-Java]
class Solution {
    public int[] minOperations(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            int res = Integer.MAX_VALUE;
            int n = 32 - Integer.numberOfLeadingZeros(x);
            int m = n / 2;
            int left = x >> m;
            for (int l = left - 1; l <= left + 1; l++) {
                // 左半反转到右半
                // 如果 n 是奇数，那么去掉回文中心再反转
                int right = Integer.reverse(l >> (n % 2)) >>> (32 - m);
                int pal = l << m | right;
                res = Math.min(res, Math.abs(x - pal));
            }
            nums[i] = res;
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minOperations(vector<int>& nums) {
        for (int& x : nums) {
            if (x == 1) {
                x = 0;
                continue;
            }
            int res = INT_MAX;
            int n = bit_width((uint32_t) x);
            int m = n / 2;
            int left = x >> m;
            for (int l = left - 1; l <= left + 1; l++) {
                // 左半反转到右半
                // 如果 n 是奇数，那么去掉回文中心再反转
                int right = __builtin_bitreverse32(l >> (n % 2)) >> (32 - m);
                int pal = l << m | right;
                res = min(res, abs(x - pal));
            }
            x = res;
        }
        return nums;
    }
};
```

```go [sol-Go]
func minOperations(nums []int) []int {
	for i, x := range nums {
		res := math.MaxInt
		n := bits.Len(uint(x))
		m := n / 2
		left := x >> m
		for l := left - 1; l <= left+1; l++ {
			// 左半反转到右半
			// 如果 n 是奇数，那么去掉回文中心再反转
			right := bits.Reverse(uint(l>>(n%2))) >> (bits.UintSize - m)
			pal := l<<m | int(right)
			res = min(res, abs(x-pal))
		}
		nums[i] = res
	}
	return nums
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [564. 寻找最近的回文数](https://leetcode.cn/problems/find-the-closest-palindrome/)
- [3677. 统计二进制回文数字的数目](https://leetcode.cn/problems/count-binary-palindromic-numbers/)

## 专题训练

见下面数学题单的「**§7.1 回文数**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
