## 题意

题干中的「整个数组」指的是 $\textit{nums}$，不是子数组。

设 $\textit{nums}$ 中的不同元素个数为 $k$，我们要计算的是 $\textit{nums}$ 中的子数组 $b$ 的个数，满足 $b$ 中不同元素个数等于 $k$。

## 思路

**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

由于子数组越长，包含的元素越多，越能满足题目要求；反之，子数组越短，包含的元素越少，越不能满足题目要求。有这种性质的题目，可以用滑动窗口解决。

枚举子数组的右端点 $\textit{right}$。同时用一个哈希表 $\textit{cnt}$ 维护子数组内每个元素的出现次数。

如果 $\textit{nums}[right]$ 加入哈希表后，发现哈希表的大小等于 $k$，说明子数组满足要求，移动子数组的左端点 $\textit{left}$，把 $\textit{nums}[\textit{left}]$ 的出现次数减一。如果 $\textit{nums}[\textit{left}]$ 的出现次数变成 $0$，则从 $\textit{cnt}$ 中去掉，表示子数组内少了一种元素。

内层循环结束后，$[\textit{left},\textit{right}]$ 这个子数组是不满足题目要求的，但在退出循环之前的最后一轮循环，$[\textit{left}-1,\textit{right}]$ 是满足题目要求的（哈希表的大小等于 $k$）。由于子数组越长，越能满足题目要求，所以除了 $[\textit{left}-1,\textit{right}]$，还有 $[\textit{left}-2,\textit{right}],[\textit{left}-3,\textit{right}],\ldots,[0,\textit{right}]$ 都是满足要求的。也就是说，当右端点**固定**在 $\textit{right}$ 时，左端点在 $0,1,2,\ldots,\textit{left}-1$ 的所有子数组都是满足要求的，这一共有 $\textit{left}$ 个。

## 答疑

**问**：刚开始没有满足要求的情况呢？子数组元素比较少，从未进入内层循环。

**答**：这种情况下 $\textit{left}=0$，说明有 $0$ 个满足要求的子数组，不影响答案。

```py [sol-Python3]
class Solution:
    def countCompleteSubarrays(self, nums: List[int]) -> int:
        k = len(set(nums))
        cnt = defaultdict(int)  # 比 Counter() 快
        ans = left = 0
        for x in nums:
            cnt[x] += 1
            while len(cnt) == k:
                out = nums[left]
                cnt[out] -= 1
                if cnt[out] == 0:
                    del cnt[out]
                left += 1
            ans += left
        return ans
```

```java [sol-Java]
class Solution {
    public int countCompleteSubarrays(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for (int x : nums) {
            set.add(x);
        }
        int k = set.size();

        Map<Integer, Integer> cnt = new HashMap<>(k);
        int ans = 0;
        int left = 0;
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
            while (cnt.size() == k) {
                int out = nums[left];
                if (cnt.merge(out, -1, Integer::sum) == 0) { // --cnt[out] == 0
                    cnt.remove(out);
                }
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
    int countCompleteSubarrays(vector<int>& nums) {
        unordered_set<int> st(nums.begin(), nums.end());
        int k = st.size();
        unordered_map<int, int> cnt;
        int ans = 0, left = 0;
        for (int x : nums) {
            cnt[x]++;
            while (cnt.size() == k) {
                int out = nums[left];
                if (--cnt[out] == 0) {
                    cnt.erase(out);
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countCompleteSubarrays(nums []int) (ans int) {
    set := map[int]struct{}{}
    for _, x := range nums {
        set[x] = struct{}{}
    }
    k := len(set)

    cnt := make(map[int]int, k)
    left := 0
    for _, x := range nums {
        ans += left
        cnt[x]++
        for len(cnt) == k {
            ans++
            out := nums[left]
            cnt[out]--
            if cnt[out] == 0 {
                delete(cnt, out)
            }
            left++
        }
    }
    return
}
```

```js [sol-JavaScript]
var countCompleteSubarrays = function(nums) {
    const k = new Set(nums).size;
    const cnt = new Map();
    let ans = 0, left = 0;
    for (const x of nums) {
        cnt.set(x, (cnt.get(x) ?? 0) + 1);
        while (cnt.size === k) {
            const out = nums[left];
            const c = cnt.get(out);
            if (c === 1) {
                cnt.delete(out);
            } else {
                cnt.set(out, c - 1);
            }
            left++;
        }
        ans += left;
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::{HashSet, HashMap};

impl Solution {
    pub fn count_complete_subarrays(nums: Vec<i32>) -> i32 {
        let k = nums.iter().collect::<HashSet<_>>().len();
        let mut cnt = HashMap::new();
        let mut ans = 0;
        let mut left = 0;
        for &x in &nums {
            *cnt.entry(x).or_insert(0) += 1;
            while cnt.len() == k {
                let out = nums[left];
                let e = cnt.get_mut(&out).unwrap();
                *e -= 1;
                if *e == 0 {
                    cnt.remove(&out);
                }
                left += 1;
            }
            ans += left;
        }
        ans as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(k)$。其中 $k$ 是 $\textit{nums}$ 中的不同元素个数，这不会超过 $n$。

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
