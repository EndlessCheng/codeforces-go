同 [2841. 几乎唯一子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/)，把 $\ge m$ 改成 $=k$ 即可。

```py [sol-Python3]
class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
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
            if len(cnt) == k:
                ans = max(ans, s)

            # 3. 离开窗口
            out = nums[left]
            s -= out
            cnt[out] -= 1
            if cnt[out] == 0:
                del cnt[out]  # 保证 len(cnt) 不计入出现 0 次的元素

        return ans
```

```java [sol-Java]
class Solution {
    public long maximumSubarraySum(int[] nums, int k) {
        long ans = 0;
        long s = 0;
        Map<Integer, Integer> cnt = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            // 1. 进入窗口
            s += nums[i];
            cnt.merge(nums[i], 1, Integer::sum); // cnt[nums[i]]++

            int left = i - k + 1;
            if (left < 0) { // 窗口大小不足 k
                continue;
            }

            // 2. 更新答案
            if (cnt.size() == k) {
                ans = Math.max(ans, s);
            }

            // 3. 离开窗口
            int out = nums[left];
            s -= out;
            int c = cnt.merge(out, -1, Integer::sum); // c = --cnt[out]
            if (c == 0) {
                cnt.remove(out); // 保证 cnt.size() 不计入出现 0 次的元素
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumSubarraySum(vector<int>& nums, int k) {
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
            if (cnt.size() == k) {
                ans = max(ans, s);
            }

            // 3. 离开窗口
            int out = nums[left];
            s -= out;
            if (--cnt[out] == 0) {
                cnt.erase(out); // 保证 cnt.size() 不计入出现 0 次的元素
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func maximumSubarraySum(nums []int, k int) (ans int64) {
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
        if len(cnt) == k {
            ans = max(ans, s)
        }

        // 3. 离开窗口
        out := nums[left]
        s -= int64(out)
        cnt[out]--
        if cnt[out] == 0 {
            delete(cnt, out) // 保证 len(cnt) 不计入出现 0 次的元素
        }
    }
    return
}
```

```js [sol-JavaScript]
var maximumSubarraySum = function(nums, k) {
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
        if (cnt.size == k) {
            ans = Math.max(ans, s);
        }

        // 3. 离开窗口
        const out = nums[left];
        s -= out;
        const c = cnt.get(out);
        if (c > 1) {
            cnt.set(out, c - 1);            
        } else {
            cnt.delete(out); // 保证 cnt.size 不包含出现 0 次的元素
        }
    }

    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn maximum_subarray_sum(nums: Vec<i32>, k: i32) -> i64 {
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
            if cnt.len() == k {
                ans = ans.max(s);
            }

            // 3. 离开窗口
            let out = nums[i - k + 1];
            s -= out as i64;
            let c = cnt.entry(out).or_insert(0);
            *c -= 1;
            if *c == 0 {
                cnt.remove(&out); // 保证 cnt.len() 不计入出现 0 次的元素
            }
        }

        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。哈希表的大小不会超过窗口长度 $k$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
