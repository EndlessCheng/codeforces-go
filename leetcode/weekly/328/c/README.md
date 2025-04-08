**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

**核心思路**：

1. 如果窗口中有 $c$ 个元素 $x$，再进来一个 $x$，会新增 $c$ 个相等数对。
2. 如果窗口中有 $c$ 个元素 $x$，再去掉一个 $x$，会减少 $c-1$ 个相等数对。

用一个哈希表 $\textit{cnt}$ 维护子数组（窗口）中的每个元素的出现次数，以及相同数对的个数 $\textit{pairs}$。

**外层循环**：从小到大枚举子数组右端点 $\textit{right}$。现在准备把 $x=\textit{nums}[\textit{right}]$ 移入窗口，那么窗口中有 $\textit{cnt}[x]$ 个数和 $x$ 相同，所以 $\textit{pairs}$ 会增加 $\textit{cnt}[x]$。然后把 $\textit{cnt}[x]$ 加一。

**内层循环**：如果发现 $\textit{pairs}\ge k$，说明子数组符合要求，右移左端点 $\textit{left}$，先把 $\textit{cnt}[\textit{nums}[\textit{left}]]$ 减少一，然后把 $\textit{pairs}$ 减少 $\textit{cnt}[\textit{nums}[\textit{left}]]$。

内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\ldots,\textit{left}-1$ 的所有子数组都是合法的，这一共有 $\textit{left}$ 个，加入答案。

```py [sol-Python3]
class Solution:
    def countGood(self, nums: List[int], k: int) -> int:
        cnt = defaultdict(int)  # 比 Counter() 快
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
            int c = cnt.getOrDefault(x, 0);
            pairs += c; // 进
            cnt.put(x, c + 1);
            while (pairs >= k) {
                x = nums[left];
                c = cnt.get(x);
                pairs -= c - 1; // 出
                cnt.put(x, c - 1);
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
```

```java [sol-Java 写法二]
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

```js [sol-JavaScript]
var countGood = function(nums, k) {
    const cnt = new Map();
    let ans = 0, pairs = 0, left = 0;
    for (const x of nums) {
        const c = cnt.get(x) ?? 0;
        pairs += c; // 进
        cnt.set(x, c + 1);
        while (pairs >= k) {
            const x = nums[left];
            const c = cnt.get(x);
            pairs -= c - 1; // 出
            cnt.set(x, c - 1);
            left++;
        }
        ans += left;
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn count_good(nums: Vec<i32>, k: i32) -> i64 {
        let mut ans = 0;
        let mut cnt = HashMap::new();
        let mut pairs = 0;
        let mut left = 0;
        for &x in &nums {
            let e = cnt.entry(x).or_insert(0);
            pairs += *e as i64;
            *e += 1;
            while pairs >= k as i64 {
                let e = cnt.get_mut(&nums[left]).unwrap();
                *e -= 1;
                pairs -= *e as i64;
                left += 1;
            }
            ans += left as i64;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
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
