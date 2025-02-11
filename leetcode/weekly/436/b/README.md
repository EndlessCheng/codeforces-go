设 $\textit{groups}$ 中的最大值为 $\textit{mx}$。我们直接预处理 $1,2,3,\ldots,\textit{mx}$ 中的每个数能被哪个 $\textit{elements}[i]$ 整除。如果有多个相同的 $\textit{elements}[i]$，只考虑最左边的那个。

从左到右遍历 $\textit{elements}$，设 $x=\textit{elements}[i]$。枚举 $x$ 的倍数 $y=x,2x,3x,\cdots$，标记 $y$ 可以被下标为 $i$ 的元素整除，记作 $\textit{target}[y]=i$。已标记的数字不再重复标记。由于我们是从左到右遍历的，这可以保证如果有多个数字都是 $y$ 的因子，我们只会记录最左边的下标。

⚠**注意**：如果我们之前遍历过 $x$ 的因子 $d$，那么不用枚举 $x$ 的倍数，因为这些数必然已被 $d$ 标记。例如 $\textit{elements}=[2,4]$，由于 $4$ 的倍数一定都是偶数（$2$ 的倍数），所以 $4$ 的倍数一定都被 $2$ 标记，所以无需枚举 $4$ 的倍数。

> 这也保证了每个数 $x$ 我们只会循环枚举其倍数一次，不会在 $\textit{elements}=[2,2,2,\ldots,2]$ 这种数据下退化成暴力。

最后，回答询问，对于 $\textit{groups}[i]$ 来说，答案为 $\textit{target}[\textit{groups}[i]]$。

**小技巧**：为了方便计算 $-1$ 的情况，可以初始化 $\textit{target}[y]=-1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ekN2ebEHx/?t=14m19s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def assignElements(self, groups: List[int], elements: List[int]) -> List[int]:
        mx = max(groups)
        target = [-1] * (mx + 1)
        for i, x in enumerate(elements):
            if x > mx or target[x] >= 0:  # x 及其倍数一定已被标记，跳过
                continue
            for y in range(x, mx + 1, x):  # 枚举 x 的倍数 y
                if target[y] < 0:  # 没有标记过
                    target[y] = i  # 标记 y 可以被 x 整除（记录 x 的下标）
        return [target[x] for x in groups]  # 回答询问
```

```java [sol-Java]
class Solution {
    public int[] assignElements(int[] groups, int[] elements) {
        int mx = 0;
        for (int x : groups) {
            mx = Math.max(mx, x);
        }
        int[] target = new int[mx + 1];
        Arrays.fill(target, -1);

        for (int i = 0; i < elements.length; i++) {
            int x = elements[i];
            if (x > mx || target[x] >= 0) { // x 及其倍数一定已被标记，跳过
                continue;
            }
            for (int y = x; y <= mx; y += x) { // 枚举 x 的倍数 y
                if (target[y] < 0) { // 没有标记过
                    target[y] = i; // 标记 y 可以被 x 整除（记录 x 的下标）
                }
            }
        }

        // 回答询问
        for (int i = 0; i < groups.length; i++) {
            groups[i] = target[groups[i]]; // 原地修改
        }
        return groups;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> assignElements(vector<int>& groups, vector<int>& elements) {
        int mx = ranges::max(groups);
        vector<int> target(mx + 1, -1);
        for (int i = 0; i < elements.size(); i++) {
            int x = elements[i];
            if (x > mx || target[x] >= 0) { // x 及其倍数一定已被标记，跳过
                continue;
            }
            for (int y = x; y <= mx; y += x) { // 枚举 x 的倍数 y
                if (target[y] < 0) { // 没有标记过
                    target[y] = i; // 标记 y 可以被 x 整除（记录 x 的下标）
                }
            }
        }

        // 回答询问
        for (int& x : groups) {
            x = target[x]; // 原地修改
        }
        return groups;
    }
};
```

```go [sol-Go]
func assignElements(groups []int, elements []int) []int {
	mx := slices.Max(groups)
	target := make([]int, mx+1)
	for i := range target {
		target[i] = -1
	}

	for i, x := range elements {
		if x > mx || target[x] >= 0 { // x 及其倍数一定已被标记，跳过
			continue
		}
		for y := x; y <= mx; y += x { // 枚举 x 的倍数 y
			if target[y] < 0 { // 没有标记过
				target[y] = i // 标记 y 可以被 x 整除（记录 x 的下标）
			}
		}
	}

	// 回答询问
	for i, x := range groups {
		groups[i] = target[x] // 原地修改
	}
	return groups
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(U\log n + m)$，其中 $n$ 是 $\textit{elements}$ 的长度，$m$ 是 $\textit{groups}$ 的长度，$U=\max(\textit{groups})$。代码中的二重循环，根据**调和级数**可得，循环次数为 $\mathcal{O}(U\log n)$。
- 空间复杂度：$\mathcal{O}(U)$。返回值不计入。

更多相似题目，可以在数学题单中搜索「**调和级数**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. 【本题相关】[数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
