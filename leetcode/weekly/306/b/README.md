用一个长为 $n$ 的数组 $\textit{score}$ 记录每个节点的得分。

遍历 $\textit{edges}$，根据题意，把 $i$ 加到 $\textit{score}[\textit{edges}[i]]$ 中。

返回 $\textit{score}[i]$ 最大且（在积分相同时）$i$ 最小的 $i$。

```py [sol-Python3]
class Solution:
    def edgeScore(self, edges: List[int]) -> int:
        ans = 0
        score = [0] * len(edges)
        for i, to in enumerate(edges):
            score[to] += i
            if score[to] > score[ans] or score[to] == score[ans] and to < ans:
                ans = to
        return ans
```

```py [sol-Python3 写法二]
# 虽然简洁，但不是一次遍历
class Solution:
    def edgeScore(self, edges: List[int]) -> int:
        score = [0] * len(edges)
        for i, to in enumerate(edges):
            score[to] += i
        return score.index(max(score))
```

```java [sol-Java]
class Solution {
    public int edgeScore(int[] edges) {
        int ans = 0;
        long[] score = new long[edges.length];
        for (int i = 0; i < edges.length; i++) {
            int to = edges[i];
            score[to] += i;
            if (score[to] > score[ans] || score[to] == score[ans] && to < ans) {
                ans = to;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int edgeScore(vector<int>& edges) {
        int n = edges.size(), ans = 0;
        vector<long long> score(n);
        for (int i = 0; i < n; i++) {
            int to = edges[i];
            score[to] += i;
            if (score[to] > score[ans] || score[to] == score[ans] && to < ans) {
                ans = to;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int edgeScore(int* edges, int edgesSize) {
    int ans = 0;
    long long* score = calloc(edgesSize, sizeof(long long));
    for (int i = 0; i < edgesSize; i++) {
        int to = edges[i];
        score[to] += i;
        if (score[to] > score[ans] || score[to] == score[ans] && to < ans) {
            ans = to;
        }
    }
    free(score);
    return ans;
}
```

```go [sol-Go]
func edgeScore(edges []int) (ans int) {
    score := make([]int, len(edges))
    for i, to := range edges {
        score[to] += i
        if score[to] > score[ans] || score[to] == score[ans] && to < ans {
            ans = to
        }
    }
    return
}
```

```js [sol-JavaScript]
var edgeScore = function(edges) {
    const score = Array(edges.length).fill(0);
    let ans = 0;
    for (let i = 0; i < edges.length; i++) {
        const to = edges[i];
        score[to] += i;
        if (score[to] > score[ans] || score[to] === score[ans] && to < ans) {
            ans = to;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn edge_score(edges: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut score = vec![0i64; edges.len()];
        for (i, &to) in edges.iter().enumerate() {
            let to = to as usize;
            score[to] += i as i64;
            if score[to] > score[ans] || score[to] == score[ans] && to < ans {
                ans = to;
            }
        }
        ans as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
