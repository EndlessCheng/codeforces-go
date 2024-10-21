借鉴 [1. 两数之和](https://leetcode.cn/problems/two-sum/) 的思路，遍历 $\textit{hours}$ 的同时，用一个哈希表（或者数组）记录元素的出现次数。

举几个例子：

- 如果 $\textit{hours}[i]=1$，那么需要知道左边有多少个模 $24$ 是 $23$ 的数，这些数加上 $1$ 都是 $24$ 的倍数。
- 如果 $\textit{hours}[i]=2$，那么需要知道左边有多少个模 $24$ 是 $22$ 的数，这些数加上 $2$ 都是 $24$ 的倍数。
- 如果 $\textit{hours}[i]=26$，那么需要知道左边有多少个模 $24$ 是 $22$ 的数，这些数加上 $26$ 都是 $24$ 的倍数。

一般地，对于 $\textit{hours}[i]$，需要知道左边有多少个模 $24$ 是 $24-\textit{hours}[i]\bmod 24$ 的数。

特别地，如果 $\textit{hours}[i]$ 模 $24$ 是 $0$，那么需要知道左边有多少个模 $24$ 也是 $0$ 的数。

这两种情况可以合并为：累加左边 

$$
(24-\textit{hours}[i]\bmod 24)\bmod 24
$$

的出现次数。

代码实现时，用一个长为 $24$ 的数组 $\textit{cnt}$ 维护 $\textit{hours}[i]\bmod 24$ 的出现次数。

请看本题 [视频讲解](https://www.bilibili.com/video/BV1T1421k7Hi/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countCompleteDayPairs(self, hours: List[int]) -> int:
        H = 24
        ans = 0
        cnt = [0] * H
        for t in hours:
            # 先查询 cnt，再更新 cnt，因为题目要求 i < j
            # 如果先更新，再查询，就把 i = j 的情况也考虑进去了
            ans += cnt[(H - t % H) % H]
            cnt[t % H] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countCompleteDayPairs(int[] hours) {
        final int H = 24;
        long ans = 0;
        int[] cnt = new int[H];
        for (int t : hours) {
            // 先查询 cnt，再更新 cnt，因为题目要求 i < j
            // 如果先更新，再查询，就把 i = j 的情况也考虑进去了
            ans += cnt[(H - t % H) % H];
            cnt[t % H]++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countCompleteDayPairs(vector<int> &hours) {
        const int H = 24;
        long long ans = 0;
        int cnt[H]{};
        for (int t : hours) {
            // 先查询 cnt，再更新 cnt，因为题目要求 i < j
            // 如果先更新，再查询，就把 i = j 的情况也考虑进去了
            ans += cnt[(H - t % H) % H];
            cnt[t % H]++;
        }
        return ans;
    }
};
```

```c [sol-C]
long long countCompleteDayPairs(int* hours, int hoursSize) {
    const int H = 24;
    long long ans = 0;
    int cnt[H] = {};
    for (int i = 0; i < hoursSize; i++) {
        int t = hours[i] % H;
        // 先查询 cnt，再更新 cnt，因为题目要求 i < j
        // 如果先更新，再查询，就把 i = j 的情况也考虑进去了
        ans += cnt[(H - t) % H];
        cnt[t]++;
    }
    return ans;
}
```

```go [sol-Go]
func countCompleteDayPairs(hours []int) (ans int64) {
    const H = 24
    cnt := [H]int{}
    for _, t := range hours {
        // 先查询 cnt，再更新 cnt，因为题目要求 i < j
        // 如果先更新，再查询，就把 i = j 的情况也考虑进去了
        ans += int64(cnt[(H-t%H)%H])
        cnt[t%H]++
    }
    return
}
```

```js [sol-JavaScript]
var countCompleteDayPairs = function(hours) {
    const H = 24;
    const cnt = Array(H).fill(0);
    let ans = 0;
    for (const t of hours) {
        // 先查询 cnt，再更新 cnt，因为题目要求 i < j
        // 如果先更新，再查询，就把 i = j 的情况也考虑进去了
        ans += cnt[(H - t % H) % H];
        cnt[t % H]++;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_complete_day_pairs(hours: Vec<i32>) -> i64 {
        const H: usize = 24;
        let mut ans = 0i64;
        let mut cnt = vec![0; H];
        for t in hours {
            let t = t as usize % H;
            // 先查询 cnt，再更新 cnt，因为题目要求 i < j
            // 如果先更新，再查询，就把 i = j 的情况也考虑进去了
            ans += cnt[(H - t) % H] as i64;
            cnt[t] += 1;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+H)$，其中 $n$ 为 $\textit{hours}$ 的长度，$H=24$。
- 空间复杂度：$\mathcal{O}(H)$。

## 一句话总结

对于有两个变量的题目，通常可以枚举其中一个变量，把它视作常量，从而转化成只有一个变量的问题。

更多相似题目，见下面数据结构题单的第零章。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
