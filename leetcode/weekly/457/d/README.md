正着走不知道执行哪个操作，倒着走就知道了，为什么？且听我细说~

从终点 $(x,y)$ 倒着走到起点 $(\textit{sx}, \textit{sy})$。

不失一般性，假设 $x\ge y$ 且 $x>0$。如果不满足，则交换 $x$ 和 $y$，以及交换 $\textit{sx}$ 和 $\textit{sy}$。

倒退一步，有四种情况：

- 我们执行了操作一：
   - 前一步是 $(x/2, y)$，这意味着 $x$ 是偶数且 $\max(x/2,y)=x/2$，即 $x\ge 2y$。
   - 前一步是 $(x-y, y)$，这意味着 $\max(x-y,y)=y$，即 $x\le 2y$。
- 我们执行了操作二：
   - 前一步是 $(x, y/2)$，这意味着 $y$ 是偶数且 $\max(x,y/2)=y/2$，即 $y/2 \ge x \ge y$，这只在 $y=0$ 时才成立。但 $x>0$，与 $y/2 \ge x$ 矛盾。所以不考虑这种情况。
   - 前一步是 $(x, y-x)$，但是 $y-x\le 0$，所以只有当 $x=y$ 时，才能这么操作。由于 $x>0$，此时 $x=y>0$。

所以：

- 如果 $x=y$：
    - 当 $\textit{sy}>0$ 的时候，只能把 $x$ 变成 $0$。
    - 当 $\textit{sx}>0$ 的时候，只能把 $y$ 变成 $0$。
    - 注意，不需要特判 $\textit{sx}=\textit{sy}=0$ 的情况，$x$ 或者 $y$ 在减半的过程中一定会变成奇数，返回 $-1$。
- 否则 $x>y$，此时只能执行操作一：
    - 如果 $x> 2y$，只能把 $x$ 减半。
    - 如果 $x< 2y$，只能把 $x$ 减少 $y$。
    - 如果 $x=2y$，二者效果一样。

按照上述规则模拟即可，每走一步，把答案加一。**注意这不是贪心，是纯模拟**。换句话说，如果我们没有返回 $-1$，那么起点与终点之间的走法**有且仅有一种**。

边界：

- 如果 $x<\textit{sx}$ 或 $y<\textit{sy}$，走过头了，返回 $-1$。
- 如果 $x=\textit{sx}$ 且 $y=\textit{sy}$，返回答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def minMoves(self, sx: int, sy: int, x: int, y: int) -> int:
        ans = 0
        while x != sx or y != sy:
            if x < sx or y < sy:
                return -1
            if x == y:
                if sy > 0:
                    x = 0
                else:
                    y = 0
                ans += 1
                continue
            # 保证 x > y
            if x < y:
                x, y = y, x
                sx, sy = sy, sx
            if x > y * 2:
                if x % 2 > 0:
                    return -1
                x //= 2
            else:
                x -= y
            ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minMoves(int sx, int sy, int x, int y) {
        int ans = 0;
        for (; x != sx || y != sy; ans++) {
            if (x < sx || y < sy) {
                return -1;
            }
            if (x == y) {
                if (sy > 0) {
                    x = 0;
                } else {
                    y = 0;
                }
                continue;
            }
            // 保证 x > y
            if (x < y) {
                int tmp = x;
                x = y;
                y = tmp;

                tmp = sx;
                sx = sy;
                sy = tmp;
            }
            if (x > y * 2) {
                if (x % 2 > 0) {
                    return -1;
                }
                x /= 2;
            } else {
                x -= y;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minMoves(int sx, int sy, int x, int y) {
        int ans = 0;
        for (; x != sx || y != sy; ans++) {
            if (x < sx || y < sy) {
                return -1;
            }
            if (x == y) {
                if (sy > 0) {
                    x = 0;
                } else {
                    y = 0;
                }
                continue;
            }
            // 保证 x > y
            if (x < y) {
                swap(x, y);
                swap(sx, sy);
            }
            if (x > y * 2) {
                if (x % 2 > 0) {
                    return -1;
                }
                x /= 2;
            } else {
                x -= y;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minMoves(sx, sy, x, y int) (ans int) {
	for ; x != sx || y != sy; ans++ {
		if x < sx || y < sy {
			return -1
		}
		if x == y {
			if sy > 0 {
				x = 0
			} else {
				y = 0
			}
			continue
		}
		// 保证 x > y
		if x < y {
			x, y = y, x
			sx, sy = sy, sx
		}
		if x > y*2 {
			if x%2 > 0 {
				return -1
			}
			x /= 2
		} else {
			x -= y
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log \max( \textit{tx}, \textit{ty}))$。注意我们在 $x\le 2y$ 时把 $x$ 减去 $y$，减去的 $y \ge x/2$，所以 $x$ 至少减半。
- 空间复杂度：$\mathcal{O}(1)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
