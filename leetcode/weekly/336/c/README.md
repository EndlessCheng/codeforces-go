示例 1 说 $[4,3,1,2,4]$ 是个美丽子数组，来看看为什么。

把每个数都写成二进制：

$$
\begin{aligned}
100     \\
011     \\
001     \\
010     \\
100     \\
\end{aligned}
$$

一次操作，可以把某一列中的两个 $1$ 都变成 $0$。

多次操作，可以把某一列中的偶数个 $1$ 都变成 $0$。

所以，如果**每一列都有偶数个** $1$，我们就能把所有数都变成 $0$。反之，如果某一列有奇数个 $1$，就不行。

设一列的元素和为 $s$，那么需要满足

$$
s\bmod 2 = 0
$$

其中 $s\bmod 2$ 可以用**异或**运算代替，因为一堆 $0$ 和 $1$ 的异或和，等同于这些数相加模 $2$ 的结果。

推广到多个列（多个比特位）：

- 如果每一列都有偶数个 $1$，那么所有数的异或和必然等于 $0$。
- 如果某一列有奇数个 $1$，那么所有数的异或和必然不等于 $0$。

所以美丽子数组等价于：

- 子数组的异或和等于 $0$。

由于异或的运算性质类似加法，可以用 [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/) 的做法（前缀和+哈希表）解决。本题相当于 $k=0$。

为什么代码要初始化 $\textit{cnt}[0] = 1$？为什么要先更新 $\textit{ans}$ 再更新 $\textit{cnt}[s]$？请看 560 题 [我的题解](https://leetcode.cn/problems/subarray-sum-equals-k/solutions/2781031/qian-zhui-he-ha-xi-biao-cong-liang-ci-bi-4mwr/)。

```py [sol-Python3]
class Solution:
    def beautifulSubarrays(self, nums: List[int]) -> int:
        ans = s = 0
        cnt = defaultdict(int)
        cnt[0] = 1
        for x in nums:
            s ^= x
            ans += cnt[s]
            cnt[s] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long beautifulSubarrays(int[] nums) {
        long ans = 0;
        int s = 0;
        Map<Integer, Integer> cnt = new HashMap<>(nums.length + 1); // 预分配空间
        cnt.put(0, 1);
        for (int x : nums) {
            s ^= x;
            int c = cnt.getOrDefault(s, 0);
            ans += c;
            cnt.put(s, c + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long beautifulSubarrays(vector<int>& nums) {
        long long ans = 0;
        int s = 0;
        unordered_map<int, int> cnt{{0, 1}};
        for (int x : nums) {
            s ^= x;
            ans += cnt[s]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func beautifulSubarrays(nums []int) (ans int64) {
    s := 0
    cnt := make(map[int]int, len(nums)+1) // 预分配空间
    cnt[0] = 1
    for _, x := range nums {
        s ^= x
        ans += int64(cnt[s])
        cnt[s]++
    }
    return
}
```

```js [sol-JavaScript]
var beautifulSubarrays = function(nums) {
    let ans = 0, s = 0;
    const cnt = new Map();
    cnt.set(0, 1);
    for (const x of nums) {
        s ^= x;
        const c = cnt.get(s) ?? 0;
        ans += c;
        cnt.set(s, c + 1);
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn beautiful_subarrays(nums: Vec<i32>) -> i64 {
        let mut ans = 0;
        let mut s = 0;
        let mut cnt = HashMap::with_capacity(nums.len() + 1); // 预分配空间
        cnt.insert(0, 1);
        for x in nums {
            s ^= x;
            let e = cnt.entry(s).or_insert(0);
            ans += *e as i64;
            *e += 1;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面数据结构题单中的「**§1.2 前缀和与哈希表**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
