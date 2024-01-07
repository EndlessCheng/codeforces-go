[本题视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)

## 方法一：BFS

每个操作都可以理解成：从 $x$ 向操作后的数连边。

在这张图上跑 BFS，求出从 $x$ 到 $y$ 的最短路，即为答案。

注意，如果 $x<y$ 那么只能用加一操作，此时可以直接算出操作次数。

代码采用双数组实现 BFS，具体请看[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)

```py [sol-Python3]
class Solution:
    def minimumOperationsToMakeEqual(self, x: int, y: int) -> int:
        if x <= y:
            return y - x
        ans = x - y  # 总操作次数不会超过 x-y
        vis = [False] * (x + ans + 1)  # +1 操作至多执行 x-y 次
        q = []
        step = 0

        def add(v: int) -> None:
            if v < y:
                nonlocal ans
                ans = min(ans, step + 1 + y - v)  # 只能执行 +1 操作
            elif not vis[v]:
                vis[v] = True
                q.append(v)

        add(x)
        while True:
            tmp = q
            q = []
            for v in tmp:
                if v == y:
                    return min(ans, step)
                if v % 11 == 0:
                    add(v // 11)
                if v % 5 == 0:
                    add(v // 5)
                add(v - 1)
                add(v + 1)
            step += 1
```

```java [sol-Java]
class Solution {
    public int minimumOperationsToMakeEqual(int x, int y) {
        if (x <= y) {
            return y - x;
        }
        int ans = x - y; // 总操作次数不会超过 x-y
        boolean[] vis = new boolean[x + ans + 1]; // +1 操作至多执行 x-y 次
        vis[x] = true;
        List<Integer> q = List.of(x);
        int step = 0;
        while (true) {
            List<Integer> tmp = q;
            q = new ArrayList<>();
            for (int v : tmp) {
                if (v == y) {
                    return Math.min(ans, step);
                }
                if (v < y) {
                    ans = Math.min(ans, step + y - v);
                    continue;
                }
                if (v % 11 == 0 && !vis[v / 11]) {
                    vis[v / 11] = true;
                    q.add(v / 11);
                }
                if (v % 5 == 0 && !vis[v / 5]) {
                    vis[v / 5] = true;
                    q.add(v / 5);
                }
                if (!vis[v - 1]) {
                    vis[v - 1] = true;
                    q.add(v - 1);
                }
                if (!vis[v + 1]) {
                    vis[v + 1] = true;
                    q.add(v + 1);
                }
            }
            step++;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperationsToMakeEqual(int x, int y) {
        if (x <= y) {
            return y - x;
        }
        int ans = x - y;  // 总操作次数不会超过 x-y
        vector<int> vis(x + ans + 1); // +1 操作至多执行 x-y 次
        vector<int> q;
        int step = 0;

        auto add = [&](int v) {
            if (v < y) {
                ans = min(ans, step + 1 + y - v); // 只能执行 +1 操作
            } else if (!vis[v]) {
                vis[v] = true;
                q.push_back(v);
            }
        };

        add(x);
        while (true) {
            auto tmp = q;
            q.clear();
            for (int v : tmp) {
                if (v == y) {
                    return min(ans, step);
                }
                if (v % 11 == 0) {
                    add(v / 11);
                }
                if (v % 5 == 0) {
                    add(v / 5);
                }
                add(v - 1);
                add(v + 1);
            }
            step++;
        }
    }
};
```

```go [sol-Go]
func minimumOperationsToMakeEqual(x, y int) int {
	if x <= y {
		return y - x
	}
	ans := x - y // 总操作次数不会超过 x-y
	vis := make([]bool, x+ans+1) // +1 操作至多执行 x-y 次
	q := []int{}
	step := 0
	add := func(v int) {
		if v < y {
			ans = min(ans, step+1+y-v) // 只能执行 +1 操作
		} else if !vis[v] {
			vis[v] = true
			q = append(q, v)
		}
	}
	add(x)
	for {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == y {
				return min(ans, step)
			}
			if v%11 == 0 {
				add(v / 11)
			}
			if v%5 == 0 {
				add(v / 5)
			}
			add(v - 1)
			add(v + 1)
		}
		step++
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(x)$。每个元素至多访问一次。
- 空间复杂度：$\mathcal{O}(x)$。

## 方法二：记忆化搜索

从 $x$ 到 $y$ 有哪些方式？

- 如果 $x= y$，无需操作。
- 如果 $x< y$，只能用加一操作。
- 如果 $x> y$，可以只用减一操作到达 $y$，或者：
- 通过减一操作到达 $x'=x - x\bmod 11$，此时 $x'$ 是 $11$ 的倍数，除以 $11$ 后，问题变成从 $x'/11$ 到 $y$ 的最少操作次数。注意，继续减少到 $x'-11$ 再除以 $11$，不如先把 $x'$ 除以 $11$ 再减一，后者可以到达同一个数，且操作次数更小。所以通过减一到达到 $x'$ 后，就无需再减一了。
- 通过加一操作到达 $x'=x + 11 - x\bmod 11$，此时 $x'$ 是 $11$ 的倍数，除以 $11$ 后，问题变成从 $x'/11$ 到 $y$ 的最少操作次数。
- 通过减一操作到达 $x'=x - x\bmod 5$，此时 $x'$ 是 $5$ 的倍数，除以 $5$ 后，问题变成从 $x'/5$ 到 $y$ 的最少操作次数。
- 通过加一操作到达 $x'=x + 5 - x\bmod 5$，此时 $x'$ 是 $5$ 的倍数，除以 $5$ 后，问题变成从 $x'/5$ 到 $y$ 的最少操作次数。

上述方式取最小值。

```py [sol-Python3]
class Solution:
    @cache
    def minimumOperationsToMakeEqual(self, x: int, y: int) -> int:
        if x <= y:
            return y - x
        return min(x - y,
                   self.minimumOperationsToMakeEqual(x // 11, y) + x % 11 + 1,
                   self.minimumOperationsToMakeEqual(x // 11 + 1, y) + 11 - x % 11 + 1,
                   self.minimumOperationsToMakeEqual(x // 5, y) + x % 5 + 1,
                   self.minimumOperationsToMakeEqual(x // 5 + 1, y) + 5 - x % 5 + 1)
```

```java [sol-Java]
class Solution {
    private final Map<Integer, Integer> memo = new HashMap<>();

    public int minimumOperationsToMakeEqual(int x, int y) {
        if (x <= y) {
            return y - x;
        }
        if (memo.containsKey(x)) {
            return memo.get(x);
        }
        int ans = x - y;
        ans = Math.min(ans, minimumOperationsToMakeEqual(x / 11, y) + x % 11 + 1);
        ans = Math.min(ans, minimumOperationsToMakeEqual(x / 11 + 1, y) + 11 - x % 11 + 1);
        ans = Math.min(ans, minimumOperationsToMakeEqual(x / 5, y) + x % 5 + 1);
        ans = Math.min(ans, minimumOperationsToMakeEqual(x / 5 + 1, y) + 5 - x % 5 + 1);
        memo.put(x, ans);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    unordered_map<int, int> memo;

public:
    int minimumOperationsToMakeEqual(int x, int y) {
        if (x <= y) {
            return y - x;
        }
        auto it = memo.find(x);
        if (it != memo.end()) {
            return it->second;
        }
        return memo[x] = min({x - y,
            minimumOperationsToMakeEqual(x / 11, y) + x % 11 + 1,
            minimumOperationsToMakeEqual(x / 11 + 1, y) + 11 - x % 11 + 1,
            minimumOperationsToMakeEqual(x / 5, y) + x % 5 + 1,
            minimumOperationsToMakeEqual(x / 5 + 1, y) + 5 - x % 5 + 1});
    }
};
```

```go [sol-Go]
func minimumOperationsToMakeEqual(x, y int) int {
	memo := map[int]int{}
	var dfs func(int) int
	dfs = func(x int) int {
		if x <= y {
			return y - x
		}
		if v, ok := memo[x]; ok {
			return v
		}
		res := min(x-y,
			dfs(x/11)+x%11+1,
			dfs(x/11+1)+11-x%11+1,
			dfs(x/5)+x%5+1,
			dfs(x/5+1)+5-x%5+1)
		memo[x] = res
		return res
	}
	return dfs(x)
}
```

#### 复杂度分析

由于除法对 $x$ 的影响远大于加减，可以认为，每次递归都把 $x$ 的规模变成 $x/5$ 和 $x/11$，当 $x\le y$ 时结束递归。故这两种操作都至多执行 $\mathcal{O}(\log (x/y))$ 次。从 $x$ 到 $y$ 的过程，$x$ 的规模会变成 $\dfrac{x}{5^p 11^q}$，其中 $p$ 和 $q$ 各有 $\mathcal{O}(\log (x/y))$ 个，所以状态个数为 $\mathcal{O}(\log^2 (x/y))$。

- 时间复杂度：$\mathcal{O}(\log^2 (x/y))$。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\log^2 (x/y))$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\log^2 (x/y))$。
- 空间复杂度：$\mathcal{O}(\log^2 (x/y))$。保存状态所需的空间等于状态个数。

#### 相似题目

- 方法一：[2059. 转化数字的最小运算数](https://leetcode.cn/problems/minimum-operations-to-convert-number/) 1850
- 方法二：[1553. 吃掉 N 个橘子的最少天数](https://leetcode.cn/problems/minimum-number-of-days-to-eat-n-oranges/) 2048 

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
