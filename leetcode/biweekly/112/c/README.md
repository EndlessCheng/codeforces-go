[【套路】教你解决定长滑窗！适用于所有定长滑窗题目！](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)

### 定长滑窗套路

1. **入**：元素 $x=\textit{nums}[i]$ 进入窗口，把 $x$ 加到元素和 $s$ 中，把 $x$ 加到哈希表中（统计 $x$ 的出现次数）。如果 $i<k-1$ 则重复第一步。
2. **更新**：如果哈希表的大小 $\ge m$，用 $s$ 更新答案的最大值。
3. **出**：元素 $x=\textit{nums}[i-k+1]$ 离开窗口，把 $s$ 减少 $x$，把哈希表中 $x$ 的出现次数减一。⚠**注意**：如果 $x$ 的出现次数变成 $0$，要从哈希表中删除 $x$，否则哈希表的大小不正确。

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int], m: int, k: int) -> int:
        ans = s = 0
        cnt = defaultdict(int)
        for i, x in enumerate(nums):
            # 1. 进入窗口
            s += x
            cnt[x] += 1

            left = i - k + 1
            if left < 0:  # 窗口大小不足 k
                continue

            # 2. 更新答案
            if len(cnt) >= m:
                ans = max(ans, s)

            # 3. 离开窗口
            out = nums[left]
            s -= out
            cnt[out] -= 1
            if cnt[out] == 0:
                del cnt[out]

        return ans
```

```java [sol-Java]
class Solution {
    public long maxSum(List<Integer> nums, int m, int k) {
        Integer[] a = nums.toArray(Integer[]::new); // 转成数组效率高

        long ans = 0;
        long s = 0;
        Map<Integer, Integer> cnt = new HashMap<>();

        for (int i = 0; i < a.length; i++) {
            // 1. 进入窗口
            s += a[i];
            cnt.merge(a[i], 1, Integer::sum); // cnt[a[i]]++

            int left = i - k + 1;
            if (left < 0) { // 窗口大小不足 k
                continue;
            }

            // 2. 更新答案
            if (cnt.size() >= m) {
                ans = Math.max(ans, s);
            }

            // 3. 离开窗口
            int out = a[left];
            s -= out;
            int c = cnt.get(out);
            if (c > 1) {
                cnt.put(out, c - 1);
            } else {
                cnt.remove(out);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSum(vector<int>& nums, int m, int k) {
        long long ans = 0, s = 0;
        unordered_map<int, int> cnt;
        for (int i = 0; i < nums.size(); i++) {
            // 1. 进入窗口
            s += nums[i];
            cnt[nums[i]]++;

            int left = i - k + 1;
            if (left < 0) { // 窗口大小不足 k
                continue;
            }

            // 2. 更新答案
            if (cnt.size() >= m) {
                ans = max(ans, s);
            }

            // 3. 离开窗口
            int out = nums[left];
            s -= out;
            if (--cnt[out] == 0) {
                cnt.erase(out);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxSum(nums []int, m, k int) (ans int64) {
    s := int64(0)
    cnt := map[int]int{}
    for i, x := range nums {
        // 1. 进入窗口
        s += int64(x)
        cnt[x]++

        left := i - k + 1
        if left < 0 { // 窗口大小不足 k
            continue
        }

        // 2. 更新答案
        if len(cnt) >= m {
            ans = max(ans, s)
        }

        // 3. 离开窗口
        out := nums[left]
        s -= int64(out)
        cnt[out]--
        if cnt[out] == 0 {
            delete(cnt, out)
        }
    }
    return
}
```

```js [sol-JavaScript]
var maxSum = function(nums, m, k) {
    const cnt = new Map();
    let ans = 0, s = 0;

    for (let i = 0; i < nums.length; i++) {
        // 1. 进入窗口
        s += nums[i];
        cnt.set(nums[i], (cnt.get(nums[i]) ?? 0) + 1);

        let left = i - k + 1;
        if (left < 0) { // 窗口大小不足 k
            continue;
        }

        // 2. 更新答案
        if (cnt.size >= m) {
            ans = Math.max(ans, s);
        }

        // 3. 离开窗口
        const out = nums[left];
        s -= out;
        const c = cnt.get(out);
        if (c > 1) {
            cnt.set(out, c - 1);            
        } else {
            cnt.delete(out);        
        }
    }

    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn max_sum(nums: Vec<i32>, m: i32, k: i32) -> i64 {
        let m = m as usize;
        let k = k as usize;
        let mut ans = 0;
        let mut s = 0;
        let mut cnt = HashMap::new();

        for (i, &x) in nums.iter().enumerate() {
            // 1. 进入窗口
            s += x as i64;
            *cnt.entry(x).or_insert(0) += 1;

            if i < k - 1 { // 窗口大小不足 k
                continue;
            }

            // 2. 更新答案
            if cnt.len() >= m {
                ans = ans.max(s);
            }

            // 3. 离开窗口
            let out = nums[i - k + 1];
            s -= out as i64;
            let c = cnt.entry(out).or_insert(0);
            *c -= 1;
            if *c == 0 {
                cnt.remove(&out);
            }
        }

        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。哈希表的大小不会超过窗口长度 $k$。

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
