把排列转化成大根笛卡尔树，问题变成：
在大小为 n 的大根笛卡尔树中，深度为 m-1 的节点恰好有 k 个。
比如根节点的值最大，深度为 0，任意包含根节点的区间最大值都等于根节点的值。

推荐用记忆化搜索写，思维难度小，还可以剪枝。

定义 dfs(dep, size, need) 表示：
当前距离 m-1 的剩余深度为 dep。
子树大小为 size。
需要在子树中找到 need 个好数（好节点）。

首先，如果 dep = 0，那么 need 减一。
然后枚举分配给左子树 leftSz 个节点，那么分配给右子树就是 size-1-leftSz 个节点。
内层循环枚举分配给左子树 leftNeed 个好节点，那么分配给右子树 need-leftNeed 个好节点。

从 size-1 个节点中选择 leftSz 个节点的方案数为 C(size-1, leftSz)。注：本题 mod 不一定是质数，逆元不一定存在，可以用 O(100^2) 的递推预处理组合数。
左子树方案数为 dfs(dep-1, leftSz, leftNeed)。
右子树方案数为 dfs(dep-1, size-1-leftSz, need-leftNeed)。
三者相乘，加到返回值中。

枚举范围：
leftSz 从 0 到 size-1。
因为要保证左右子树的 need <= size，所以 leftNeed 从 max(need-(size-1-leftSz), 0) 到 min(leftSz, need)。

递归边界：
1. 如果 dep < 0，无需继续递归：
   1.1. 如果 need > 0，不合法，返回 0；
   1.2. 否则 need = 0，这 size 个节点随意排列，有 size! 个方案。注：预处理阶乘。
2. 如果 size = 0，我们找到了一个合法方案，返回 1。

递归入口：dfs(m-1, n, k)。

重要剪枝：
如果递归左子树的返回值是 0，那么无需递归右子树，因为乘积一定是 0。

时间复杂度是 O(n^2 * m * k^2)，不过由于剪枝、循环长度等因素，常数很小。

https://codeforces.com/contest/1580/submission/321732868
https://www.luogu.com.cn/paste/8t13pvw2
