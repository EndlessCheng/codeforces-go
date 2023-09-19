请看 [视频讲解](https://b23.tv/PDz9NBA) 第三题。

挨个判断每台机器最多可以制造多少份合金。

假设要制造 $\textit{num}$ 份合金，由于 $\textit{num}$ 越小，花费的钱越少，$\textit{num}$ 越多，花费的钱越多，有**单调性**，可以二分。

有关二分的基础知识，请看视频讲解[【基础算法精讲 04】](https://b23.tv/CBJnyNJ)。

对于第 $j$ 类金属：

- 如果 $\textit{composition}[i][j]\cdot \textit{num} \le \textit{stock}[j]$，那么无需购买额外的金属。
- 如果 $\textit{composition}[i][j]\cdot \textit{num} > \textit{stock}[j]$，那么需要购买额外的金属，花费为

$$
(\textit{composition}[i][j]\cdot \textit{num} - \textit{stock}[j])\cdot \textit{cost}[j]
$$

如果总花费超过 $\textit{budget}$，则无法制造 $\textit{num}$ 份合金，否则可以制造。

二分上界：粗略计算一下，假设 $\textit{composition}[i][j]$ 和 $\textit{cost}[j]$ 都是 $1$，此时可以制造最多的合金，个数为 $\min(\textit{stock}) + \textit{budget}$。

```py [sol-Python3]
class Solution:
    def maxNumberOfAlloys(self, n: int, k: int, budget: int, composition: List[List[int]], stock: List[int], cost: List[int]) -> int:
        ans = 0
        mx = min(stock) + budget
        for com in composition:
            def check(num: int) -> int:
                money = 0
                for s, base, c in zip(stock, com, cost):
                    if s < base * num:
                        money += (base * num - s) * c
                        if money > budget:
                            return False
                return True

            left, right = 0, mx + 1
            while left + 1 < right:  # 开区间写法
                mid = (left + right) // 2
                if check(mid):
                    left = mid
                else:
                    right = mid
            ans = max(ans, left)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxNumberOfAlloys(int n, int k, int budget, List<List<Integer>> composition, List<Integer> stock, List<Integer> cost) {
        int ans = 0;
        int mx = Collections.min(stock) + budget;
        for (var com : composition) {
            int left = 0, right = mx + 1;
            while (left + 1 < right) { // 开区间写法
                int mid = (left + right) / 2;
                boolean ok = true;
                long money = 0;
                for (int i = 0; i < n; ++i) {
                    if (stock.get(i) < (long) com.get(i) * mid) {
                        money += ((long) com.get(i) * mid - stock.get(i)) * cost.get(i);
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
            ans = Math.max(ans, left);
        }
        return ans;
    }
}
```

```java [sol-Java 数组优化]
// 全部转成 int[] 数组，效率比 List<Integer> 更高
class Solution {
    public int maxNumberOfAlloys(int n, int k, int budget, List<List<Integer>> composition, List<Integer> Stock, List<Integer> Cost) {
        int ans = 0;
        int mx = Collections.min(Stock) + budget;
        int[] stock = Stock.stream().mapToInt(i -> i).toArray();
        int[] cost = Cost.stream().mapToInt(i -> i).toArray();
        for (var Com : composition) {
            int[] com = Com.stream().mapToInt(i -> i).toArray();
            int left = 0, right = mx + 1;
            while (left + 1 < right) { // 开区间写法
                int mid = (left + right) / 2;
                boolean ok = true;
                long money = 0;
                for (int i = 0; i < n; ++i) {
                    if (stock[i] < (long) com[i] * mid) {
                        money += ((long) com[i] * mid - stock[i]) * cost[i];
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
            ans = Math.max(ans, left);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxNumberOfAlloys(int n, int k, int budget, vector<vector<int>> &composition, vector<int> &stock, vector<int> &cost) {
        int ans = 0;
        int mx = *min_element(stock.begin(), stock.end()) + budget;
        for (auto &com: composition) {
            auto check = [&](long long num) -> bool {
                long long money = 0;
                for (int i = 0; i < n; i++) {
                    if (stock[i] < com[i] * num) {
                        money += (com[i] * num - stock[i]) * cost[i];
                        if (money > budget) {
                            return false;
                        }
                    }
                }
                return true;
            };
            int left = 0, right = mx + 1;
            while (left + 1 < right) { // 开区间写法
                int mid = (left + right) / 2;
                (check(mid) ? left : right) = mid;
            }
            ans = max(ans, left);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxNumberOfAlloys(_, _, budget int, composition [][]int, stock, cost []int) (ans int) {
	mx := stock[0]
	for _, s := range stock {
		mx = min(mx, s)
	}
	mx += budget
	for _, com := range composition {
		res := sort.Search(mx, func(num int) bool {
			num++
			money := 0
			for i, s := range stock {
				if s < com[i]*num {
					money += (com[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
		ans = max(ans, res)
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var maxNumberOfAlloys = function (n, k, budget, composition, stock, cost) {
    let ans = 0;
    const mx = Math.min(...stock) + budget;
    for (const com of composition) {
        function check(num) {
            let money = 0;
            for (let i = 0; i < n; i++) {
                if (stock[i] < com[i] * num) {
                    money += (com[i] * num - stock[i]) * cost[i];
                    if (money > budget) {
                        return false;
                    }
                }
            }
            return true;
        }
        let left = 0, right = mx + 1;
        while (left + 1 < right) { // 开区间写法
            const mid = (left + right) >> 1;
            if (check(mid)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        ans = Math.max(ans, left);
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kn\log U)$，其中 $U=\min(\textit{stock}) + \textit{budget}$。
- 空间复杂度：$\mathcal{O}(1)$。

## 题单·二分答案

- [875. 爱吃香蕉的珂珂](https://leetcode.cn/problems/koko-eating-bananas/)
- [1283. 使结果不超过阈值的最小除数](https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/)
- [2187. 完成旅途的最少时间](https://leetcode.cn/problems/minimum-time-to-complete-trips/)
- [2226. 每个小孩最多能分到多少糖果](https://leetcode.cn/problems/maximum-candies-allocated-to-k-children/)
- [1870. 准时到达的列车最小时速](https://leetcode.cn/problems/minimum-speed-to-arrive-on-time/)
- [1011. 在 D 天内送达包裹的能力](https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/)
- [2064. 分配给商店的最多商品的最小值](https://leetcode.cn/problems/minimized-maximum-of-products-distributed-to-any-store/)
- [1760. 袋子里最少数目的球](https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/)
- [1482. 制作 m 束花所需的最少天数](https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets/)
- [1642. 可以到达的最远建筑](https://leetcode.cn/problems/furthest-building-you-can-reach/)
- [1898. 可移除字符的最大数目](https://leetcode.cn/problems/maximum-number-of-removable-characters/)
- [778. 水位上升的泳池中游泳](https://leetcode.cn/problems/swim-in-rising-water/)
- [2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/)
