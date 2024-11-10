[本题视频讲解](https://www.bilibili.com/video/BV1Lm4y1N7mf/)

**挨个判断**每台机器最多可以制造多少份合金。

假设要制造 $\textit{num}$ 份合金，由于 $\textit{num}$ 越小，花费的钱越少，$\textit{num}$ 越多，花费的钱越多，有**单调性**，可以二分。

关于二分的基础知识，以及各种开闭区间的写法，请看视频讲解[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

对于第 $j$ 类金属：

- 如果 $\textit{composition}[i][j]\cdot \textit{num} \le \textit{stock}[j]$，那么无需购买额外的金属。
- 如果 $\textit{composition}[i][j]\cdot \textit{num} > \textit{stock}[j]$，那么需要购买额外的金属，花费为

$$
(\textit{composition}[i][j]\cdot \textit{num} - \textit{stock}[j])\cdot \textit{cost}[j]
$$

遍历每类金属，计算总花费。如果总花费超过 $\textit{budget}$，则无法制造 $\textit{num}$ 份合金，否则可以制造。

最后讨论下二分的上下界：

- **二分上界**：粗略计算一下，假设 $\textit{composition}[i][j]$ 和 $\textit{cost}[j]$ 都是 $1$，此时可以制造最多的合金，个数为 $\min(\textit{stock}) + \textit{budget}$。
- **二分下界**：可以设为 $1$。更巧妙的做法是，设当前答案为 $\textit{ans}$，那么可以初始化二分下界为 $\textit{ans}+1$，因为算出小于等于 $\textit{ans}$ 的值是没有意义的，不会让 $\textit{ans}$ 变大。如果这台机器无法制造出至少 $\textit{ans}+1$ 份合金，那么二分循环结束后的结果为 $\textit{ans}$，不影响答案。

下面的代码采用开区间写法，要把上界加一，下界减一。

```py [sol-Python3]
class Solution:
    def maxNumberOfAlloys(self, n: int, k: int, budget: int, composition: List[List[int]], stock: List[int], cost: List[int]) -> int:
        ans = 0
        mx = min(stock) + budget
        for comp in composition:
            def check(num: int) -> bool:
                money = 0
                for s, base, c in zip(stock, comp, cost):
                    if s < base * num:
                        money += (base * num - s) * c
                        if money > budget:
                            return False
                return True

            left, right = ans, mx + 1
            while left + 1 < right:  # 开区间写法
                mid = (left + right) // 2
                if check(mid):
                    left = mid
                else:
                    right = mid
            ans = left
        return ans
```

```py [sol-Python3 库函数]
class Solution:
    def maxNumberOfAlloys(self, n: int, k: int, budget: int, composition: List[List[int]], stock: List[int], cost: List[int]) -> int:
        ans = 0
        mx = min(stock) + budget
        for comp in composition:
            def f(num: int) -> int:
                money = 0
                for s, base, c in zip(stock, comp, cost):
                    if s < base * num:
                        money += (base * num - s) * c
                        if money > budget:
                            break
                return money
            ans += bisect_right(range(ans + 1, mx + 1), budget, key=f)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxNumberOfAlloys(int n, int k, int budget, List<List<Integer>> composition, List<Integer> Stock, List<Integer> Cost) {
        int ans = 0;
        int mx = Collections.min(Stock) + budget;
        Integer[] stock = Stock.toArray(Integer[]::new); // 转成数组更快
        Integer[] cost = Cost.toArray(Integer[]::new);
        for (List<Integer> Comp : composition) {
            Integer[] comp = Comp.toArray(Integer[]::new);
            int left = ans;
            int right = mx + 1;
            while (left + 1 < right) { // 开区间写法
                int mid = left + (right - left) / 2;
                boolean ok = true;
                long money = 0;
                for (int i = 0; i < n; i++) {
                    if (stock[i] < (long) comp[i] * mid) {
                        money += ((long) comp[i] * mid - stock[i]) * cost[i];
                        if (money > budget) {
                            ok = false;
                            break;
                        }
                    }
                }
                if (ok) {
                    left = mid;
                } else {
                    right = mid;
                }
            }
            ans = left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxNumberOfAlloys(int n, int, int budget, vector<vector<int>>& composition, vector<int>& stock, vector<int>& cost) {
        int ans = 0;
        int mx = ranges::min(stock) + budget;
        for (auto& comp : composition) {
            auto check = [&](long long num) -> bool {
                long long money = 0;
                for (int i = 0; i < n; i++) {
                    if (stock[i] < comp[i] * num) {
                        money += (comp[i] * num - stock[i]) * cost[i];
                        if (money > budget) {
                            return false;
                        }
                    }
                }
                return true;
            };
            int left = ans, right = mx + 1;
            while (left + 1 < right) { // 开区间写法
                int mid = left + (right - left) / 2;
                (check(mid) ? left : right) = mid;
            }
            ans = left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxNumberOfAlloys(_, _, budget int, composition [][]int, stock, cost []int) (ans int) {
	mx := slices.Min(stock) + budget
	for _, comp := range composition {
		ans += sort.Search(mx-ans, func(num int) bool {
			num += ans + 1
			money := 0
			for i, s := range stock {
				if s < comp[i]*num {
					money += (comp[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
	}
	return
}
```

```js [sol-JavaScript]
var maxNumberOfAlloys = function(n, k, budget, composition, stock, cost) {
    let ans = 0;
    const mx = Math.min(...stock) + budget;
    for (const comp of composition) {
        function check(num) {
            let money = 0;
            for (let i = 0; i < n; i++) {
                if (stock[i] < comp[i] * num) {
                    money += (comp[i] * num - stock[i]) * cost[i];
                    if (money > budget) {
                        return false;
                    }
                }
            }
            return true;
        }
        let left = ans, right = mx + 1;
        while (left + 1 < right) { // 开区间写法
            const mid = Math.floor((left + right) / 2);
            if (check(mid)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        ans = left;
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kn\log U)$，其中 $U=\min(\textit{stock}) + \textit{budget}$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. 【本题相关】[二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
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
