用埃氏筛（或者欧拉筛）预处理一个布尔数组，表示哪些数是质数。详见 [本题视频讲解](https://www.bilibili.com/video/BV1GCNRzgEYp/)。

注意 $1$ 不是质数。

用哈希表（或者数组）统计每个元素的出现次数，如果有出现次数为质数的数，返回 $\texttt{true}$。否则返回 $\texttt{false}$。

```py [sol-Python3]
MX = 101
is_prime = [False] * 2 + [True] * (MX - 2)
for i in range(2, isqrt(MX) + 1):
    if is_prime[i]:
        for j in range(i * i, MX, i):
            is_prime[j] = False  # j 是质数 i 的倍数

class Solution:
    def checkPrimeFrequency(self, nums: List[int]) -> bool:
        cnt = Counter(nums)
        return any(is_prime[c] for c in cnt.values())
```

```java [sol-Java]
class Solution {
    private static final int MX = 101;
    private static final boolean[] NOT_PRIME = new boolean[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        NOT_PRIME[0] = NOT_PRIME[1] = true;
        for (int i = 2; i * i < MX; i++) {
            if (NOT_PRIME[i]) {
                continue;
            }
            for (int j = i * i; j < MX; j += i) {
                NOT_PRIME[j] = true; // j 是质数 i 的倍数
            }
        }
    }

    public boolean checkPrimeFrequency(int[] nums) {
        init();

        // 更快的写法见【Java 数组】
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum);
        }

        for (int c : cnt.values()) {
            if (!NOT_PRIME[c]) {
                return true;
            }
        }
        return false;
    }
}
```

```java [sol-Java 数组]
class Solution {
    private static final int MX = 101;
    private static final boolean[] NOT_PRIME = new boolean[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        NOT_PRIME[0] = NOT_PRIME[1] = true;
        for (int i = 2; i * i < MX; i++) {
            if (NOT_PRIME[i]) {
                continue;
            }
            for (int j = i * i; j < MX; j += i) {
                NOT_PRIME[j] = true; // j 是质数 i 的倍数
            }
        }
    }

    public boolean checkPrimeFrequency(int[] nums) {
        init();

        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] cnt = new int[mx + 1];
        for (int x : nums) {
            cnt[x]++;
        }

        for (int c : cnt) {
            if (!NOT_PRIME[c]) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
const int MX = 101;
bool np[MX];

auto init = [] {
    np[1] = true;
    for (int i = 2; i * i < MX; i++) {
        if (!np[i]) {
            for (int j = i * i; j < MX; j += i) {
                np[j] = true; // j 是质数 i 的倍数
            }
        }
    }
    return 0;
}();

class Solution {
public:
    bool checkPrimeFrequency(vector<int>& nums) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }
        for (auto& [_, c] : cnt) {
            if (!np[c]) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
const mx = 101

var np = [mx]bool{1: true}

func init() {
	// 质数=false，非质数=true
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func checkPrimeFrequency(nums []int) bool {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for _, c := range cnt {
		if !np[c] {
			return true
		}
	}
	return false
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

见下面数学题单的「**§1.1 判断质数**」和「**§1.2 预处理质数**」。

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
