**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

用一个哈希表 $\textit{cnt}$ 维护子数组内的每个元素的出现次数，以及相同数对的个数 $\textit{pairs}$。

滑动窗口外层循环：从小到大枚举子数组右端点 $\textit{right}$。现在准备把 $x=\textit{nums}[\textit{right}]$ 移入子数组，那么子数组中有 $\textit{cnt}[x]$ 个数和 $x$ 相同，所以 $\textit{pairs}$ 会增加 $\textit{cnt}[x]$。

滑动窗口内层循环：如果发现 $\textit{pairs}\ge k$，说明子数组符合要求，则右移左端点 $\textit{left}$，把 $\textit{pairs}$ 减少 $\textit{cnt}[\textit{nums}[\textit{left}]]$。

内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\ldots,\textit{left}-1$ 的所有子数组都是合法的，这一共有 $\textit{left}$ 个，加入答案。

```py [sol-Python3]
class Solution:
    def countGood(self, nums: List[int], k: int) -> int:
        cnt = defaultdict(int)
        ans = left = pairs = 0
        for x in nums:
            pairs += cnt[x]
            cnt[x] += 1
            while pairs >= k:
                cnt[nums[left]] -= 1
                pairs -= cnt[nums[left]]
                left += 1
            ans += left
        return ans
```

```java [sol-Java]
class Solution {
    public long countGood(int[] nums, int k) {
        long ans = 0;
        Map<Integer, Integer> cnt = new HashMap<>();
        int pairs = 0;
        int left = 0;
        for (int x : nums) {
            pairs += cnt.merge(x, 1, Integer::sum) - 1; // pairs += cnt[x]++
            while (pairs >= k) {
                pairs -= cnt.merge(nums[left], -1, Integer::sum); // pairs -= --cnt[nums[left]]
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countGood(vector<int>& nums, int k) {
        long long ans = 0;
        unordered_map<int, int> cnt;
        int pairs = 0, left = 0;
        for (int x : nums) {
            pairs += cnt[x]++;
            while (pairs >= k) {
                pairs -= --cnt[nums[left]];
                left++;
            }
            ans += left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countGood(nums []int, k int) (ans int64) {
    cnt := map[int]int{}
    pairs, left := 0, 0
    for _, x := range nums {
        pairs += cnt[x]
        cnt[x]++
        for pairs >= k {
            cnt[nums[left]]--
            pairs -= cnt[nums[left]]
            left++
        }
        ans += int64(left)
    }
    return
}
````

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面滑动窗口题单中的「**§2.3.1 越长越合法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
