设 $k = 2^p$。从 $\textit{nums}$ 中选出的数必须形如 $x^{2^p}$。

设 $m = \max(\textit{nums})$。当 $x\ge 2$ 时，我们有

$$
x^{2^p} \le m
$$

解得

$$
p \le \log_2 \log_x m
$$

本题 $m\le 10^9$，所以在 $x\ge 2$ 的情况下，$p$ 至多为 $4$，满足条件的序列最长是

$$
[x,x^2,x^4,x^8,x^{16},x^8,x^4,x^2,x]
$$

如何求出序列的长度？

暴力枚举 $\textit{nums}$ 中的数，作为 $x$，然后看 $x^2,x^4,\ldots$ 在 $\textit{nums}$ 中的个数，直到 $x^k$ 的个数不足 $2$ 个为止。如果此时 $\textit{nums}$ 包含 $x^k$，那么把 $x^k$ 放正中间，序列长度加一；否则去掉一个 $x^{k/2}$，序列长度减一。

注意特判 $x=1$ 的情况。设 $\textit{cnt}_1$ 为 $\textit{nums}$ 中的 $1$ 的个数。如果 $\textit{cnt}_1$ 是奇数，那么全 $1$ 序列的长度是 $\textit{cnt}_1$，否则是 $\textit{cnt}_1 - 1$。

[本题视频讲解](https://www.bilibili.com/video/BV1we411J7Y8/?t=3m46s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        cnt = Counter(nums)

        ans = (cnt[1] - 1) | 1  # 保证 ans 是奇数（奇数不变，偶数减一）
        del cnt[1]

        for x in cnt:
            res = 0
            while cnt.get(x, 0) >= 2:  # 用 get 而不是 []，避免把不在 nums 中的数插入 cnt
                res += 2
                x *= x
            ans = max(ans, res + (1 if x in cnt else -1))

        return ans
```

```java [sol-Java]
class Solution {
    public int maximumLength(int[] nums) {
        HashMap<Long, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge((long) x, 1, Integer::sum); // cnt[x]++
        }

        Integer cnt1 = cnt.remove(1L);
        int ans = cnt1 != null ? (cnt1 - 1) | 1 : 0; // 保证 ans 是奇数（奇数不变，偶数减一）

        for (long x : cnt.keySet()) {
            int res = 0;
            while (cnt.getOrDefault(x, 0) >= 2) {
                res += 2;
                x *= x;
            }
            ans = Math.max(ans, res + (cnt.containsKey(x) ? 1 : -1));
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(vector<int>& nums) {
        unordered_map<long long, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }

        int ans = (cnt[1] - 1) | 1; // 保证 ans 是奇数（奇数不变，偶数减一）
        cnt.erase(1);

        for (auto& [num, _] : cnt) {
            int res = 0;
            auto x = num;
            while (cnt.contains(x) && cnt[x] >= 2) {
                res += 2;
                x *= x;
            }
            ans = max(ans, res + (cnt.contains(x) ? 1 : -1));
        }

        return ans;
    }
};
```

```go [sol-Go]
func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	ans := cnt[1] - 1 | 1 // 保证 ans 是奇数（奇数不变，偶数减一）
	delete(cnt, 1)

	for x := range cnt {
		res := 0
		for cnt[x] >= 2 {
			res += 2
			x *= x
		}
		res += cnt[x]
		ans = max(ans, res-1|1) // 保证 ans 是奇数（奇数不变，偶数减一）
	}

	return ans
}
```

```js [sol-JavaScript]
var maximumLength = function (nums) {
    const cnt = new Map();
    for (const x of nums) {
        cnt.set(x, (cnt.get(x) ?? 0) + 1);
    }

    let ans = ((cnt.get(1) ?? 0) - 1) | 1; // 保证 ans 是奇数（奇数不变，偶数减一）
    cnt.delete(1);

    for (let x of cnt.keys()) {
        let res = 0;
        while ((cnt.get(x) ?? 0) >= 2) {
            res += 2;
            x *= x;
        }
        ans = Math.max(ans, res + (cnt.has(x) ? 1 : -1));
    }

    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn maximum_length(nums: Vec<i32>) -> i32 {
        let mut cnt = HashMap::new();
        for x in nums {
            *cnt.entry(x as i64).or_insert(0) += 1;
        }

        let mut ans = (cnt.remove(&1).unwrap_or(0) - 1) | 1; // 保证 ans 是奇数（奇数不变，偶数减一）

        for &x in cnt.keys() {
            let mut res = 0;
            let mut x = x;
            while *cnt.get(&x).unwrap_or(&0) >= 2 {
                res += 2;
                x *= x;
            }
            ans = ans.max(res + if cnt.contains_key(&x) { 1 } else { -1 });
        }

        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log \log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

## 附：O(n) 做法

利用 [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/) 的技巧，可以做到 $\mathcal{O}(n)$ 时间复杂度。原理见 [我的题解](https://leetcode.cn/problems/longest-consecutive-sequence/solutions/3005726/ha-xi-biao-on-zuo-fa-pythonjavacgojsrust-whop/)。

```py
class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        cnt = Counter(nums)

        ans = (cnt[1] - 1) | 1  # 保证 ans 是奇数（奇数不变，偶数减一）
        del cnt[1]

        for x in cnt:
            rt = isqrt(x)
            if rt * rt == x and cnt.get(rt, 0) >= 2:
                continue  # 如果有两个 sqrt(x) 也在 nums 中，那么从 sqrt(x) 开始算的序列更长

            res = 0
            while cnt.get(x, 0) >= 2:
                res += 2
                x *= x
            ans = max(ans, res + (1 if x in cnt else -1))

        return ans
```

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
