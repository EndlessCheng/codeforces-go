## 什么是第 x 小？

例如数组 $[1,1,1,2,2]$，其中第 $1$ 小、第 $2$ 小和第 $3$ 小的数都是 $1$，第 $4$ 小和第 $5$ 小的数都是 $2$。

- 第 $x$ 小等价于：求**最小**的 $v$，满足 $\le v$ 的数**至少**有 $x$ 个。
- 第 $x$ 大等价于：求**最大**的 $v$，满足 $\ge v$ 的数**至少**有 $x$ 个。

## 思路

本题数据范围很小（$-50\le \textit{nums}[i]\le 50$），我们可以借鉴**计数排序**的思想，用一个 $\textit{cnt}$ 数组维护窗口内每个数的出现次数。

枚举 $v=-50,-49,\ldots,-1$，统计 $\le v$ 的元素个数。如果个数 $\ge x$，那么 $v$ 就是答案。

```py [sol-Python3]
class Solution:
    def getSubarrayBeauty(self, nums: List[int], k: int, x: int) -> List[int]:
        cnt = [0] * 101
        for num in nums[:k - 1]:  # 先往窗口内添加 k-1 个数
            cnt[num] += 1

        ans = [0] * (len(nums) - k + 1)
        for i, (in_, out) in enumerate(zip(nums[k - 1:], nums)):
            cnt[in_] += 1  # 进入窗口（保证窗口有恰好 k 个数）
            left = x  # 从 x 开始倒着减，减到 <= 0 就找到了答案
            for v in range(-50, 0):  # 暴力枚举负数范围 [-50, -1]
                left -= cnt[v]
                if left <= 0:  # 找到答案
                    ans[i] = v
                    break
            cnt[out] -= 1  # 离开窗口
        return ans
```

```java [sol-Java]
class Solution {
    public int[] getSubarrayBeauty(int[] nums, int k, int x) {
        final int BIAS = 50;
        int[] cnt = new int[BIAS * 2 + 1];
        for (int i = 0; i < k - 1; i++) { // 先往窗口内添加 k-1 个数
            cnt[nums[i] + BIAS]++;
        }

        int n = nums.length;
        int[] ans = new int[n - k + 1];
        for (int i = k - 1; i < n; i++) {
            cnt[nums[i] + BIAS]++; // 进入窗口（保证窗口有恰好 k 个数）
            int left = x; // 从 x 开始倒着减，减到 <= 0 就找到了答案
            for (int v = -50; v < 0; v++) { // 暴力枚举负数范围 [-50, -1]
                left -= cnt[v + BIAS];
                if (left <= 0) { // 找到答案
                    ans[i - k + 1] = v;
                    break;
                }
            }
            cnt[nums[i - k + 1] + BIAS]--; // 离开窗口
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> getSubarrayBeauty(vector<int>& nums, int k, int x) {
        constexpr int BIAS = 50;
        int cnt[BIAS * 2 + 1]{};
        for (int i = 0; i < k - 1; i++) { // 先往窗口内添加 k-1 个数
            cnt[nums[i] + BIAS]++;
        }

        int n = nums.size();
        vector<int> ans(n - k + 1);
        for (int i = k - 1; i < n; i++) {
            cnt[nums[i] + BIAS]++; // 进入窗口（保证窗口有恰好 k 个数）
            int left = x; // 从 x 开始倒着减，减到 <= 0 就找到了答案
            for (int v = -50; v < 0; v++) { // 暴力枚举负数范围 [-50, -1]
                left -= cnt[v + BIAS];
                if (left <= 0) { // 找到答案
                    ans[i - k + 1] = v;
                    break;
                }
            }
            cnt[nums[i - k + 1] + BIAS]--; // 离开窗口
        }
        return ans;
    }
};
```

```go [sol-Go]
func getSubarrayBeauty(nums []int, k, x int) []int {
    const bias = 50
    cnt := [bias*2 + 1]int{}
    for _, num := range nums[:k-1] { // 先往窗口内添加 k-1 个数
        cnt[num+bias]++
    }

    ans := make([]int, len(nums)-k+1)
    for i, num := range nums[k-1:] {
        cnt[num+bias]++ // 进入窗口（保证窗口有恰好 k 个数）
        left := x // 从 x 开始倒着减，减到 <= 0 就找到了答案
        for v, c := range cnt[:bias] { // 暴力枚举负数范围 [-50, -1]
            left -= c
            if left <= 0 { // 找到答案
                ans[i] = v - bias
                break
            }
        }
        cnt[nums[i]+bias]-- // 离开窗口
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=50$。
- 空间复杂度：$\mathcal{O}(U)$。

## 附：有序集合写法

```py [sol-Python3]
class Solution:
    def getSubarrayBeauty(self, nums: List[int], k: int, x: int) -> List[int]:
        sl = SortedList(nums[:k - 1])  # SortedList 来自 sortedcontainers
        ans = []
        for in_, out in zip(nums[k - 1:], nums):
            sl.add(in_)  # 进入窗口（保证窗口有恰好 k 个数）
            ans.append(min(sl[x - 1], 0))
            sl.discard(out)  # 离开窗口（也可以写 remove）
        return ans
```

```cpp [sol-C++]
#include <ext/pb_ds/assoc_container.hpp>

using namespace __gnu_pbds;

// 使用 pair<value, index> 支持重复元素
using ordered_set = tree<pair<int, int>, null_type, less<>, rb_tree_tag, tree_order_statistics_node_update>;

class Solution {
public:
    vector<int> getSubarrayBeauty(vector<int>& nums, int k, int x) {
        ordered_set st;
        for (int i = 0; i < k - 1; i++) {
            st.insert({nums[i], i});
        }

        int n = nums.size();
        vector<int> ans(n - k + 1);
        for (int i = k - 1; i < n; i++) {
            st.insert({nums[i], i});
            auto [v, _] = *st.find_by_order(x - 1); // 第 x 小
            ans[i - k + 1] = min(v, 0);
            st.erase({nums[i - k + 1], i - k + 1});
        }
        return ans;
    }
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
