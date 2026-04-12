用**埃氏筛**预处理每个数是不是质数。

如何找下一个质数（非质数）？可以暴力枚举，也可以在质数列表中二分查找。

[本题视频讲解](https://www.bilibili.com/video/BV14hDQBDEUu/)，欢迎点赞关注~

```py [sol-Python3]
MX = 100_004  # 1e5 的下一个质数是 1e5 + 3
is_prime = [0, 0] + [1] * (MX - 2)
for i in range(2, isqrt(MX) + 1):
    if is_prime[i]:
        for j in range(i * i, MX, i):
            is_prime[j] = 0

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        for i, x in enumerate(nums):
            # 如果 i 是偶数，那么循环直到 is_prime[x] == 1（x 是质数）
            # 如果 i 是奇数，那么循环直到 is_prime[x] == 0（x 不是质数）
            while is_prime[x] == i % 2:
                ans += 1
                x += 1
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MX = 100_004; // 1e5 的下一个质数是 1e5 + 3
    private static final int[] notPrime = new int[MX];
    private static boolean initialized = false;

    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        notPrime[1] = 1;
        for (int i = 2; i * i < MX; i++) {
            if (notPrime[i] == 0) {
                for (int j = i * i; j < MX; j += i) {
                    notPrime[j] = 1;
                }
            }
        }
    }

    public int minOperations(int[] nums) {
        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            // 如果 i 是偶数，那么循环直到 notPrime[x] == 0（x 是质数）
            // 如果 i 是奇数，那么循环直到 notPrime[x] == 1（x 不是质数）
            while (notPrime[x] != i % 2) {
                ans++;
                x++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
constexpr int MX = 100'004; // 1e5 的下一个质数是 1e5 + 3
int8_t not_prime[MX];

auto init = [] {
    not_prime[1] = 1;
    for (int i = 2; i * i < MX; i++) {
        if (!not_prime[i]) {
            for (int j = i * i; j < MX; j += i) {
                not_prime[j] = 1;
            }
        }
    }
    return 0;
}();

class Solution {
public:
    int minOperations(vector<int>& nums) {
        int ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            // 如果 i 是偶数，那么循环直到 not_prime[x] == 0（x 是质数）
            // 如果 i 是奇数，那么循环直到 not_prime[x] == 1（x 不是质数）
            while (not_prime[x] != i % 2) {
                ans++;
                x++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
const mx = 100_004 // 1e5 的下一个质数是 1e5 + 3
var notPrime = [mx]int{1, 1}

func init() {
	for i := 2; i*i < mx; i++ {
		if notPrime[i] == 0 {
			for j := i * i; j < mx; j += i {
				notPrime[j] = 1
			}
		}
	}
}

func minOperations(nums []int) (ans int) {
	for i, x := range nums {
		// 如果 i 是偶数，那么循环直到 notPrime[x] == 0（x 是质数）
		// 如果 i 是奇数，那么循环直到 notPrime[x] == 1（x 不是质数）
		for notPrime[x] != i%2 {
			ans++
			x++
		}
	}
	return
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(nG)$，其中 $n$ 是 $\textit{nums}$ 的长度，$G\le 72$ 是本题数据范围下的最大质数间隔。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数学题单的「**§1.2 预处理质数（筛质数）**」。

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
