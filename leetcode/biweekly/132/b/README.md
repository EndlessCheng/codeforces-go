本题和 5.19 的每日一题 [1535. 找出数组游戏的赢家](https://leetcode.cn/problems/find-the-winner-of-an-array-game/) 几乎一样，请看 [我的题解](https://leetcode.cn/problems/find-the-winner-of-an-array-game/solution/mo-ni-fu-ruo-gan-jin-jie-wen-ti-pythonja-zx17/)，把代码中的 $\textit{mx}$ 改成最大值的下标即可。

[本题视频讲解](https://www.bilibili.com/video/BV1Tx4y1b7wk/?t=2m40s)

```py [sol-Python3]
class Solution:
    def findWinningPlayer(self, skills: List[int], k: int) -> int:
        max_i = win = 0
        for i in range(1, len(skills)):
            if skills[i] > skills[max_i]:  # 打擂台，发现新的最大值
                max_i = i
                win = 0
            win += 1  # 获胜回合 +1
            if win == k:  # 连续赢下 k 场比赛
                break
        # 如果 k 很大，那么 max_i 就是 skills 最大值的下标，毕竟最大值会一直赢下去
        return max_i
```

```java [sol-Java]
class Solution {
    public int findWinningPlayer(int[] skills, int k) {
        int maxI = 0;
        int win = 0;
        for (int i = 1; i < skills.length && win < k; i++) {
            if (skills[i] > skills[maxI]) { // 打擂台，发现新的最大值
                maxI = i;
                win = 0;
            }
            win++; // 获胜回合 +1
        }
        // 如果 k 很大，那么 maxI 就是 skills 最大值的下标，毕竟最大值会一直赢下去
        return maxI;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findWinningPlayer(vector<int>& skills, int k) {
        int max_i = 0, win = 0;
        for (int i = 1; i < skills.size() && win < k; i++) {
            if (skills[i] > skills[max_i]) { // 打擂台，发现新的最大值
                max_i = i;
                win = 0;
            }
            win++; // 获胜回合 +1
        }
        // 如果 k 很大，那么 max_i 就是 skills 最大值的下标，毕竟最大值会一直赢下去
        return max_i;
    }
};
```

```c [sol-C]
int findWinningPlayer(int* skills, int n, int k) {
    int max_i = 0, win = 0;
    for (int i = 1; i < n && win < k; i++) {
        if (skills[i] > skills[max_i]) { // 打擂台，发现新的最大值
            max_i = i;
            win = 0;
        }
        win++; // 获胜回合 +1
    }
    // 如果 k 很大，那么 max_i 就是 skills 最大值的下标，毕竟最大值会一直赢下去
    return max_i;
}
```

```go [sol-Go]
func findWinningPlayer(skills []int, k int) (maxI int) {
    win := 0
    for i := 1; i < len(skills) && win < k; i++ {
        if skills[i] > skills[maxI] { // 打擂台，发现新的最大值
            maxI = i
            win = 0
        }
        win++ // 获胜回合 +1
    }
    // 如果 k 很大，那么 maxI 就是 skills 最大值的下标，毕竟最大值会一直赢下去
    return
}
```

```js [sol-JavaScript]
var findWinningPlayer = function(skills, k) {
    let maxI = 0, win = 0;
    for (let i = 1; i < skills.length && win < k; i++) {
        if (skills[i] > skills[maxI]) { // 打擂台，发现新的最大值
            maxI = i;
            win = 0;
        }
        win++; // 获胜回合 +1
    }
    // 如果 k 很大，那么 maxI 就是 skills 最大值的下标，毕竟最大值会一直赢下去
    return maxI;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_winning_player(skills: Vec<i32>, k: i32) -> i32 {
        let mut max_i = 0;
        let mut win = 0;
        for i in 1..skills.len() {
            if skills[i] > skills[max_i] { // 打擂台，发现新的最大值
                max_i = i;
                win = 0;
            }
            win += 1; // 获胜回合 +1
            if win == k { // 连续赢下 k 场比赛
                break;
            }
        }
        // 如果 k 很大，那么 max_i 就是 skills 最大值的下标，毕竟最大值会一直赢下去
        max_i as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{skills}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
