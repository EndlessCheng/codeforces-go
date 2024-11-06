个人赛五道题目的 [视频讲解](https://www.bilibili.com/video/BV1zN4y1K762) 已出炉，欢迎点赞三连，在评论区分享你对这场比赛的看法~

---

定义状态 (当前节点，祖先节点开关 2 的切换次数的奇偶性，父节点是否切换了开关 3)，每个状态表示从当前状态出发，最少需要操作多少次开关，可以关闭子树所有节点的灯。

跑一个树形 DP。如果当前受到祖先节点的开关影响后，变成开灯状态，那么可以操作一个或三个开关：

- 操作开关 1；
- 操作开关 2；
- 操作开关 3；
- 操作开关 123；
- 这四种操作取最小值。

如果变成关灯状态，那么可以操作零个或两个开关：

- 不操作任何一个开关；
- 操作开关 12；
- 操作开关 13；
- 操作开关 23；
- 这四种操作取最小值。

```py [sol1-Python3]
class Solution:
    def closeLampInTree(self, root: TreeNode) -> int:
        @cache  # 记忆化搜索
        def dfs(node: TreeNode, switch2: bool, switch3: bool) -> int:
            if node is None:
                return 0
            if (node.val == 1) == (switch2 == switch3):  # 当前节点为开灯
                res1 = dfs(node.left, switch2, False) + dfs(node.right, switch2, False) + 1
                res2 = dfs(node.left, not switch2, False) + dfs(node.right, not switch2, False) + 1
                res3 = dfs(node.left, switch2, True) + dfs(node.right, switch2, True) + 1
                res123 = dfs(node.left, not switch2, True) + dfs(node.right, not switch2, True) + 3
                return min(res1, res2, res3, res123)
            else:  # 当前节点为关灯
                res0 = dfs(node.left, switch2, False) + dfs(node.right, switch2, False)
                res12 = dfs(node.left, not switch2, False) + dfs(node.right, not switch2, False) + 2
                res13 = dfs(node.left, switch2, True) + dfs(node.right, switch2, True) + 2
                res23 = dfs(node.left, not switch2, True) + dfs(node.right, not switch2, True) + 2
                return min(res0, res12, res13, res23)
        return dfs(root, False, False)
```

```go [sol1-Go]
func closeLampInTree(root *TreeNode) int {
    type tuple struct {
        node             *TreeNode
        switch2, switch3 bool
    }
    memo := map[tuple]int{} // 记忆化搜索
    var dfs func(*TreeNode, bool, bool) int
    dfs = func(node *TreeNode, switch2, switch3 bool) int {
        if node == nil {
            return 0
        }
        p := tuple{node, switch2, switch3}
        if res, ok := memo[p]; ok { // 之前计算过
            return res
        }
        if node.Val == 1 == (switch2 == switch3) { // 当前节点为开灯
            res1 := dfs(node.Left, switch2, false) + dfs(node.Right, switch2, false) + 1
            res2 := dfs(node.Left, !switch2, false) + dfs(node.Right, !switch2, false) + 1
            res3 := dfs(node.Left, switch2, true) + dfs(node.Right, switch2, true) + 1
            r123 := dfs(node.Left, !switch2, true) + dfs(node.Right, !switch2, true) + 3
            memo[p] = min(res1, res2, res3, r123)
        } else { // 当前节点为关灯
            res0 := dfs(node.Left, switch2, false) + dfs(node.Right, switch2, false)
            res12 := dfs(node.Left, !switch2, false) + dfs(node.Right, !switch2, false) + 2
            res13 := dfs(node.Left, switch2, true) + dfs(node.Right, switch2, true) + 2
            res23 := dfs(node.Left, !switch2, true) + dfs(node.Right, !switch2, true) + 2
            memo[p] = min(res0, res12, res13, res23)
        }
        return memo[p]
    }
    return dfs(root, false, false)
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为二叉树的节点个数。
- 空间复杂度：$O(n)$。

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
