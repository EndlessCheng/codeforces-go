分类讨论：

- 如果 $k=n$，只有一个子数组，所以每个数都满足要求，返回 $\textit{nums}$ 中的最大值。
- 如果 $k=1$，有 $n$ 个长为 $1$ 的子数组，只有出现次数等于 $1$ 的元素满足要求，返回出现次数等于 $1$ 的最大元素。
- 如果 $1 < k < n$，只有 $\textit{nums}[0]$ 和 $\textit{nums}[n-1]$ 是可能满足要求的数，因为**其他元素至少出现在两个子数组中**。返回这两个数中的出现次数等于 $1$ 的最大元素。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1QP9bY3EL6/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def f(self, nums: List[int], x: int) -> int:
        return -1 if x in nums else x

    def largestInteger(self, nums: List[int], k: int) -> int:
        if k == len(nums):
            return max(nums)
        if k == 1:
            ans = -1
            for x, c in Counter(nums).items():
                if c == 1:
                    ans = max(ans, x)
            return ans
        # nums[0] 不能出现在其他地方，nums[-1] 同理
        return max(self.f(nums[1:], nums[0]), self.f(nums[:-1], nums[-1]))
```

```java [sol-Java]
class Solution {
    public int largestInteger(int[] nums, int k) {
        int n = nums.length;
        if (k == n) {
            return Arrays.stream(nums).max().getAsInt();
        }
        if (k == 1) {
            Map<Integer, Integer> cnt = new HashMap<>();
            for (int x : nums) {
                cnt.merge(x, 1, Integer::sum); // cnt[x]++
            }
            int ans = -1;
            for (var e : cnt.entrySet()) {
                if (e.getValue() == 1) {
                    ans = Math.max(ans, e.getKey());
                }
            }
            return ans;
        }
        // nums[0] 不能出现在其他地方，nums[n-1] 同理
        return Math.max(f(nums, 1, n, nums[0]), f(nums, 0, n - 1, nums[n - 1]));
    }

    private int f(int[] nums, int begin, int end, int x) {
        for (int i = begin; i < end; i++) {
            if (nums[i] == x) {
                return -1;
            }
        }
        return x;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int largestInteger(vector<int>& nums, int k) {
        int n = nums.size();
        if (k == n) {
            return ranges::max(nums);
        }
        if (k == 1) {
            unordered_map<int, int> cnt;
            for (int x : nums) {
                cnt[x]++;
            }
            int ans = -1;
            for (auto& [x, c] : cnt) {
                if (c == 1) {
                    ans = max(ans, x);
                }
            }
            return ans;
        }
        auto f = [&](int begin, int end, int x) {
            if (find(nums.begin() + begin, nums.begin() + end, x) != nums.begin() + end) {
                return -1;
            }
            return x;
        };
        // nums[0] 不能出现在其他地方，nums[n-1] 同理
        return max(f(1, n, nums[0]), f(0, n - 1, nums[n - 1]));
    }
};
```

```go [sol-Go]
func largestInteger(nums []int, k int) int {
    n := len(nums)
    if k == n {
        return slices.Max(nums)
    }
    if k == 1 {
        cnt := map[int]int{}
        for _, x := range nums {
            cnt[x]++
        }
        ans := -1
        for x, c := range cnt {
            if c == 1 {
                ans = max(ans, x)
            }
        }
        return ans
    }
    // nums[0] 不能出现在其他地方，nums[n-1] 同理
    return max(f(nums[1:], nums[0]), f(nums[:n-1], nums[n-1]))
}

func f(a []int, x int) int {
    if slices.Contains(a, x) {
        return -1
    }
    return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
