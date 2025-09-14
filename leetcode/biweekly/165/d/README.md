我们要计算两个子序列的所有元素的异或和。由于同一个数异或两次等于 $0$，所以只需关心那些出现**一次**的元素。问题相当于：

- 从 $\textit{nums}$ 中选出**一个**子序列，计算子序列异或和的最大值。

这是 [线性基](https://oi-wiki.org/math/basis/) 的标准应用，具体请看 [视频讲解](https://www.bilibili.com/video/BV1pm8vzAEXx/)，欢迎点赞关注~

**注**：异或运算本质是 $w$ 维线性空间中的模 $2$ 加法，一个二进制数可以视作一个 $w$ 维的向量，本题 $w\le 30$。线性基（线性异或基）计算的是这个线性空间中的由 $S'$ 张成的一组基，$S'$ 中的每个二进制数（视作向量）都可以被这组基表出。

```py [sol-Python3]
class XorBasis:
    # n 为值域最大值 U 的二进制长度，例如 U=1e9 时 n=30
    def __init__(self, n: int):
        self.b = [0] * n

    def insert(self, x: int) -> None:
        b = self.b
        # 从高到低遍历，保证计算 max_xor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
        for i in range(len(b) - 1, -1, -1):
            if x >> i:  # 由于大于 i 的位都被我们异或成了 0，所以 x >> i 的结果只能是 0 或 1
                if b[i] == 0:  # x 和之前的基是线性无关的
                    b[i] = x  # 新增一个基，最高位为 i
                    return
                x ^= b[i]  # 保证每个基的二进制长度互不相同
        # 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基

    def max_xor(self) -> int:
        b = self.b
        res = 0
        # 从高到低贪心：越高的位，越必须是 1
        # 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for i in range(len(b) - 1, -1, -1):
            if res ^ b[i] > res:  # 手写 max 更快
                res ^= b[i]
        return res

class Solution:
    def maxXorSubsequences(self, nums: List[int]) -> int:
        m = max(nums).bit_length()
        b = XorBasis(m)
        for x in nums:
            b.insert(x)
        return b.max_xor()
```

```java [sol-Java]
class XorBasis {
    private final int[] b;

    // n 为值域最大值 U 的二进制长度，例如 U=1e9 时 n=30
    public XorBasis(int n) {
        b = new int[n];
    }

    public void insert(int x) {
        // 从高到低遍历，保证计算 maxXor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
        for (int i = b.length - 1; i >= 0; i--) {
            if ((x >> i) > 0) { // 由于大于 i 的位都被我们异或成了 0，所以 x >> i 的结果只能是 0 或 1
                if (b[i] == 0) { // x 和之前的基是线性无关的
                    b[i] = x; // 新增一个基，最高位为 i
                    return;
                }
                x ^= b[i]; // 保证每个基的二进制长度互不相同
            }
        }
        // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
    }

    public int maxXor() {
        int res = 0;
        // 从高到低贪心：越高的位，越必须是 1
        // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for (int i = b.length - 1; i >= 0; i--) {
            res = Math.max(res, res ^ b[i]);
        }
        return res;
    }
}

class Solution {
    public int maxXorSubsequences(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }
        int m = 32 - Integer.numberOfLeadingZeros(mx);
        XorBasis b = new XorBasis(m);
        for (int x : nums) {
            b.insert(x);
        }
        return b.maxXor();
    }
}
```

```cpp [sol-C++]
class XorBasis {
    vector<int> b;

public:
    // n 为值域最大值 U 的二进制长度，例如 U=1e9 时 n=30
    XorBasis(int n) : b(n) {}

    void insert(int x) {
        // 从高到低遍历，保证计算 max_xor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
        for (int i = b.size() - 1; i >= 0; i--) {
            if (x >> i) { // 由于大于 i 的位都被我们异或成了 0，所以 x >> i 的结果只能是 0 或 1
                if (b[i] == 0) { // x 和之前的基是线性无关的
                    b[i] = x; // 新增一个基，最高位为 i
                    return;
                }
                x ^= b[i]; // 保证每个基的二进制长度互不相同
            }
        }
        // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
    }

    int max_xor() {
        int res = 0;
        // 从高到低贪心：越高的位，越必须是 1
        // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for (int i = b.size() - 1; i >= 0; i--) {
            res = max(res, res ^ b[i]);
        }
        return res;
    }
};

class Solution {
public:
    int maxXorSubsequences(vector<int>& nums) {
        int m = bit_width((uint32_t) ranges::max(nums));
        XorBasis b(m);
        for (int x : nums) {
            b.insert(x);
        }
        return b.max_xor();
    }
};
```

```go [sol-Go]
type xorBasis []int

// n 为值域最大值 U 的二进制长度，例如 U=1e9 时 n=30
func newXorBasis(n int) xorBasis {
	return make(xorBasis, n)
}

func (b xorBasis) insert(x int) {
	// 从高到低遍历，保证计算 maxXor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
	for i := len(b) - 1; i >= 0; i-- {
		if x>>i == 0 { // 由于大于 i 的位都被我们异或成了 0，所以 x>>i 的结果只能是 0 或 1
			continue
		}
		if b[i] == 0 { // x 和之前的基是线性无关的
			b[i] = x // 新增一个基，最高位为 i
			return
		}
		x ^= b[i] // 保证每个基的二进制长度互不相同
	}
	// 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
}

func (b xorBasis) maxXor() (res int) {
	// 从高到低贪心：越高的位，越必须是 1
	// 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
	for i := len(b) - 1; i >= 0; i-- {
		res = max(res, res^b[i])
	}
	return
}

func maxXorSubsequences(nums []int) int {
	u := slices.Max(nums)
	m := bits.Len(uint(u))
	b := newXorBasis(m)
	for _, x := range nums {
		b.insert(x)
	}
	return b.maxXor()
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(\log U)$。

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
