本题和 5.19 的每日一题 [1535. 找出数组游戏的赢家](https://leetcode.cn/problems/find-the-winner-of-an-array-game/) 几乎一样，请看 [我的题解](https://leetcode.cn/problems/find-the-winner-of-an-array-game/solution/mo-ni-fu-ruo-gan-jin-jie-wen-ti-pythonja-zx17/)，把代码中的 $\textit{mx}$ 改成最大值的下标即可。

[视频讲解](https://www.bilibili.com/video/BV1Tx4y1b7wk/)

```py [sol-Python3]
class Solution:
    def findWinningPlayer(self, skills: List[int], k: int) -> int:
        mx_i = 0
        win = -1  # 对于 skills[0] 来说，需要连续 k+1 个回合都是最大值
        for i, x in enumerate(skills):
            if x > skills[mx_i]:  # 新的最大值
                mx_i = i
                win = 0
            win += 1  # 获胜回合 +1
            if win == k:
                break
        return mx_i
```

```java [sol-Java]
class Solution {
    public int findWinningPlayer(int[] skills, int k) {
        int mxI = 0;
        int win = 0;
        for (int i = 1; i < skills.length && win < k; i++) {
            if (skills[i] > skills[mxI]) { // 新的最大值
                mxI = i;
                win = 0;
            }
            win++; // 获胜回合 +1
        }
        return mxI;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findWinningPlayer(vector<int>& skills, int k) {
        int mx_i = 0, win = 0;
        for (int i = 1; i < skills.size() && win < k; i++) {
            if (skills[i] > skills[mx_i]) { // 新的最大值
                mx_i = i;
                win = 0;
            }
            win++; // 获胜回合 +1
        }
        return mx_i;
    }
};
```

```go [sol-Go]
func findWinningPlayer(skills []int, k int) (mxI int) {
	win := 0
	for i := 1; i < len(skills) && win < k; i++ {
		if skills[i] > skills[mxI] { // 新的最大值
			mxI = i
			win = 0
		}
		win++ // 获胜回合 +1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{skills}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
