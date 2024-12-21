## 贪心策略

首先，每条水平线和垂直线，最终都要全部切完。

- 水平线（横切）开销 $\textit{horizontalCut}[i]$ 对答案的**贡献**，等于 $\textit{horizontalCut}[i]$ 乘以横切次数（经过多少块蛋糕），即在此之前的竖切次数加一。
- 垂直线（竖切）开销 $\textit{verticalCut}[j]$ 对答案的**贡献**，等于 $\textit{verticalCut}[j]$ 乘以竖切次数（经过多少块蛋糕），即在此之前的横切次数加一。

⚠注意：本题虽然不要求「一切到底」，但计算贡献的式子是不变的。无论是否一切到底，最终得到的式子，合并同类项之后，仍然是「横切开销」乘以「横切次数」，以及「竖切开销」乘以「竖切次数」，**计算结果与一切到底是一样的**。所以下面的分析，**只需考虑一切到底的情况**。

以示例 1 为例，其操作序列为

$$
竖切\ 0,\ 横切\ 0,\ 横切\ 1
$$

最小总开销为

$$
\textit{verticalCut}[0]\cdot 1 + \textit{horizontalCut}[0] \cdot 2 + \textit{horizontalCut}[1] \cdot 2
$$

对于一个操作序列，交换其中两个相邻的横切，不影响答案；交换其中两个相邻的竖切，也不影响答案。所以重点考察交换两个相邻的横切和竖切。

> 为方便讨论，交换相邻的，也可以交换不相邻的，最终得出的结论是一样的。

设横切的开销为 $h$，如果先横切，设需要横切 $\textit{cntH}$ 次。

设竖切的开销为 $v$，如果先竖切，设需要竖切 $\textit{cntV}$ 次。

- 先横切，再竖切，那么竖切的次数（这一刀经过的蛋糕块数）要多 $1$，开销为 
   $$
   h\cdot \textit{cntH} + v\cdot (\textit{cntV}+1)
   $$
- 先竖切，再横切，那么横切的次数（这一刀经过的蛋糕块数）要多 $1$，开销为 
   $$
   v\cdot \textit{cntV} + h\cdot (\textit{cntH}+1)
   $$

如果先横再竖开销更小，则有

$$
h\cdot \textit{cntH} + v\cdot (\textit{cntV}+1) < v\cdot \textit{cntV} + h\cdot (\textit{cntH}+1)
$$

化简得

$$
h > v
$$

这意味着，**谁的开销更大，就先切谁**，并且这个先后顺序与 $\textit{cntH}$ 和 $\textit{cntV}$ 无关。换句话说，按照该规则去切蛋糕，得到的操作序列，如果把开销大的操作移动后面，必然会得到更大的总开销。

## 写法一

1. 把 $\textit{horizontalCut}$ 和 $\textit{verticalCut}$ 从大到小排序。
2. 初始化 $\textit{cntH} = 1, \textit{cntV} = 1, i = 0, j = 0$。
3. 双指针遍历 $\textit{horizontalCut}$ 和 $\textit{verticalCut}$。
4. 如果 $\textit{horizontalCut}[i] > \textit{verticalCut}[j]$，那么优先横切，把 $\textit{horizontalCut}[i]\cdot \textit{cntH}$ 加入答案，$i$ 增加 $1$，然后需要竖切的次数增加，把 $\textit{cntV}$ 增加 $1$；否则优先竖切，把 $\textit{verticalCut}[j]\cdot \textit{cntV}$ 加入答案，$j$ 增加 $1$，然后需要横切的次数增加，把 $\textit{cntH}$ 增加 $1$。
5. 循环直到两个数组都遍历完。
6. 返回答案。

注意 $i=m-1$ 和 $j=n-1$ 的情况。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1LZ421u7Ut/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumCost(self, m: int, n: int, horizontalCut: List[int], verticalCut: List[int]) -> int:
        horizontalCut.sort(reverse=True)
        verticalCut.sort(reverse=True)
        ans = i = j = 0
        cnt_h = cnt_v = 1
        while i < m - 1 or j < n - 1:
            if j == n - 1 or i < m - 1 and horizontalCut[i] > verticalCut[j]:
                ans += horizontalCut[i] * cnt_h  # 横切
                i += 1
                cnt_v += 1  # 需要竖切的蛋糕块增加
            else:
                ans += verticalCut[j] * cnt_v  # 竖切
                j += 1
                cnt_h += 1  # 需要横切的蛋糕块增加
        return ans
```

```java [sol-Java]
class Solution {
    public long minimumCost(int m, int n, int[] horizontalCut, int[] verticalCut) {
        Arrays.sort(horizontalCut); // 下面倒序遍历
        Arrays.sort(verticalCut);
        long ans = 0;
        int i = m - 2;
        int j = n - 2;
        int cntH = 1;
        int cntV = 1;
        while (i >= 0 || j >= 0) {
            if (j < 0 || i >= 0 && horizontalCut[i] > verticalCut[j]) {
                ans += horizontalCut[i--] * cntH; // 横切
                cntV++; // 需要竖切的蛋糕块增加
            } else {
                ans += verticalCut[j--] * cntV; // 竖切
                cntH++; // 需要横切的蛋糕块增加
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(int m, int n, vector<int>& horizontalCut, vector<int>& verticalCut) {
        ranges::sort(horizontalCut, greater());
        ranges::sort(verticalCut, greater());
        long long ans = 0;
        int cnt_h = 1, cnt_v = 1;
        int i = 0, j = 0;
        while (i < m - 1 || j < n - 1) {
            if (j == n - 1 || i < m - 1 && horizontalCut[i] > verticalCut[j]) {
                ans += horizontalCut[i++] * cnt_h; // 横切
                cnt_v++; // 需要竖切的蛋糕块增加
            } else {
                ans += verticalCut[j++] * cnt_v; // 竖切
                cnt_h++; // 需要横切的蛋糕块增加
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(m, n int, horizontalCut, verticalCut []int) int64 {
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	cntH, cntV := 1, 1
	i, j := 0, 0
	for i < m-1 || j < n-1 {
		if j == n-1 || i < m-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * cntH // 横切
			i++
			cntV++ // 需要竖切的蛋糕块增加
		} else {
			ans += verticalCut[j] * cntV // 竖切
			j++
			cntH++ // 需要横切的蛋糕块增加
		}
	}
	return int64(ans)
}
```

## 写法二（优化）

$\textit{cntH}$ 和 $\textit{cntV}$ 这两个变量可以省略，因为从上面的过程可以发现，$\textit{cntH}=j+1,\ \textit{cntV}=i+1$。

```py [sol-Python3]
class Solution:
    def minimumCost(self, m: int, n: int, horizontalCut: List[int], verticalCut: List[int]) -> int:
        horizontalCut.sort(reverse=True)
        verticalCut.sort(reverse=True)
        ans = i = j = 0
        while i < m - 1 or j < n - 1:
            if j == n - 1 or i < m - 1 and horizontalCut[i] > verticalCut[j]:
                ans += horizontalCut[i] * (j + 1)  # 横切
                i += 1
            else:
                ans += verticalCut[j] * (i + 1)  # 竖切
                j += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long minimumCost(int m, int n, int[] horizontalCut, int[] verticalCut) {
        Arrays.sort(horizontalCut); // 下面倒序遍历
        Arrays.sort(verticalCut);
        long ans = 0;
        int i = m - 2;
        int j = n - 2;
        while (i >= 0 || j >= 0) {
            if (j < 0 || i >= 0 && horizontalCut[i] > verticalCut[j]) {
                ans += horizontalCut[i--] * (n - 1 - j); // 横切
            } else {
                ans += verticalCut[j--] * (m - 1 - i); // 竖切
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(int m, int n, vector<int>& horizontalCut, vector<int>& verticalCut) {
        ranges::sort(horizontalCut, greater());
        ranges::sort(verticalCut, greater());
        long long ans = 0;
        int i = 0, j = 0;
        while (i < m - 1 || j < n - 1) {
            if (j == n - 1 || i < m - 1 && horizontalCut[i] > verticalCut[j]) {
                ans += horizontalCut[i++] * (j + 1); // 横切
            } else {
                ans += verticalCut[j++] * (i + 1); // 竖切
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(m, n int, horizontalCut, verticalCut []int) int64 {
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	i, j := 0, 0
	for i < m-1 || j < n-1 {
		if j == n-1 || i < m-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * (j + 1) // 横切
			i++
		} else {
			ans += verticalCut[j] * (i + 1) // 竖切
			j++
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m + n\log n)$。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。不计入排序的栈开销。

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
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
