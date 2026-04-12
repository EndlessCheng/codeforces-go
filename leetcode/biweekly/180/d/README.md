本题是 [179. 最大数](https://leetcode.cn/problems/largest-number/) 的二进制版本，做法是一样的。

对于本题，无需先生成字符串再比较。对于两个片段 $S$ 和 $T$，如果 $S$ 中的 $\texttt{1}$ 更多，那么 $S$ 在 $T$ 左边。如果 $S$ 和 $T$ 中的 $\texttt{1}$ 一样多，那么 $\texttt{0}$ 更少的在左边。

**特殊情况**：没有 $\texttt{0}$ 的片段（全为 $\texttt{1}$）排在最前面。

设当前答案为 $\textit{ans}$，拼接方法如下：

- 在 $\textit{ans}$ 右边拼接 $k$ 个 $0$，等价于把 $\textit{ans}$ 左移 $k$ 位，等价于 $\textit{ans}\cdot 2^k$。
- 在 $\textit{ans}$ 右边拼接 $k$ 个 $1$，先左移 $k$ 位（乘以 $2^k$），再加上 $2^k - 1$，即 $\textit{ans}\cdot 2^k + 2^k - 1 = (\textit{ans} + 1) \cdot 2^k - 1$。

代码实现时，注意取模。为什么可以在**中途取模**？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

由于本题 $k$ 比较小，可以预处理所有 $2^k$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
MOD = 1_000_000_007
MX = 10001
# 预处理 2 的幂
pow2 = [1] * MX
for i in range(1, MX):
    pow2[i] = pow2[i - 1] * 2 % MOD

class Solution:
    def maxValue(self, nums1: List[int], nums0: List[int]) -> int:
        # 依次判断：没有 0 的排在最前面，1 多的排前面，0 少的排前面
        idx = sorted(range(len(nums1)), key=lambda i: (nums0[i] != 0, -nums1[i], nums0[i]))

        ans = 0
        for i in idx:
            ans = ((ans + 1) * pow2[nums1[i]] - 1) * pow2[nums0[i]] % MOD
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 10001;
    private static final int[] pow2 = new int[MX];
    private static boolean initialized = false;

    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理 2 的幂
        pow2[0] = 1;
        for (int i = 1; i < MX; i++) {
            pow2[i] = pow2[i - 1] * 2 % MOD;
        }
    }

    public int maxValue(int[] nums1, int[] nums0) {
        Integer[] idx = new Integer[nums1.length];
        for (int i = 0; i < nums1.length; i++) {
            idx[i] = i;
        }

        Arrays.sort(idx, (i, j) -> {
            if (nums0[i] == 0 && nums0[j] != 0) {
                return -1;
            }
            if (nums0[i] != 0 && nums0[j] == 0) {
                return 1;
            }
            if (nums1[i] != nums1[j]) {
                return nums1[j] - nums1[i];
            }
            return nums0[i] - nums0[j];
        });

        long ans = 0;
        for (int i : idx) {
            ans = ((ans + 1) * pow2[nums1[i]] - 1) % MOD * pow2[nums0[i]] % MOD;
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
constexpr int MOD = 1'000'000'007;
constexpr int MX = 10001;
int pow2[MX] = {1};

auto init = [] {
    // 预处理 2 的幂
    for (int i = 1; i < MX; i++) {
        pow2[i] = pow2[i - 1] * 2 % MOD;
    }
    return 0;
}();

class Solution {
public:
    int maxValue(vector<int>& nums1, vector<int>& nums0) {
        vector<int> idx(nums1.size());
        ranges::iota(idx, 0); // idx[i] = i

        // 依次判断：没有 0 的排在最前面，1 多的排前面，0 少的排前面
        ranges::sort(idx, {}, [&](int i) {
            return tuple(nums0[i] != 0, -nums1[i], nums0[i]);
        });

        long long ans = 0;
        for (int i : idx) {
            ans = ((ans + 1) * pow2[nums1[i]] - 1) % MOD * pow2[nums0[i]] % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 10001
var pow2 = [mx]int{1}

func init() {
	// 预处理 2 的幂
	for i := 1; i < mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func maxValue(nums1, nums0 []int) (ans int) {
	idx := make([]int, len(nums1))
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int {
		if nums0[i] == 0 {
			return -1
		}
		if nums0[j] == 0 {
			return 1
		}
		return cmp.Or(nums1[j]-nums1[i], nums0[i]-nums0[j])
	})

	for _, i := range idx {
		ans = ((ans+1)*pow2[nums1[i]] - 1) % mod * pow2[nums0[i]] % mod
	}
	return
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：对于更大的值域范围，可以用**快速幂**，原理见[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

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
