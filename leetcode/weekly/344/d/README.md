视频讲解请看[【周赛 344】](https://www.bilibili.com/video/BV1YL41187Rx/)第四题，欢迎点赞投币！

## 提示 1

考虑根到两个互为兄弟节点（父节点相同）的叶子的两条路径。

由于这两条路径除了叶子节点不一样，其余节点都一样，所以为了让这两条路径的路径和相等，必须修改叶子节点的值。

设叶子节点的值分别为 $x$ 和 $y$，假设 $x\le y$，是否需要同时增加 $x$ 和 $y$ 呢？

这是不需要的，把 $x$ 增加 $y-x$ 就行，因为我们可以增加它们的祖先节点的值，使得它们俩的路径和与其它的路径和相等，这样可以节省操作次数。

## 提示 2

对于不是叶子的兄弟节点，又要如何比较和计算呢？

和上面的分析一样，从根到当前节点的路径，除了这两个兄弟节点不一样，其余节点都一样。所以把路径和从叶子往上传，这样就可以按照提示 1 那样比较了。

示例 1 如下图，把节点 $2$ 的路径和视作 $x+5+3=x+8$，节点 $3$ 的路径和视作 $x+2+3=x+5$（其中 $x$ 是在节点 $2,3$ 之上的路径和），这样可以知道需要把节点 $3$ 的值增加 $(x+8)-(x+5)=8-5=3$。

![lc2673.png](https://pic.leetcode.cn/1709024171-NFqAWc-lc2673.png)

代码实现时，可以直接在 $\textit{cost}$ 上累加路径和。由于 $\textit{cost}$ 数组的下标是从 $0$ 开始的，所以节点编号转成下标需要减一。

```py [sol-Python3]
class Solution:
    def minIncrements(self, n: int, cost: List[int]) -> int:
        ans = 0
        for i in range(n // 2, 0, -1):  # 从最后一个非叶节点开始算
            ans += abs(cost[i * 2 - 1] - cost[i * 2])  # 两个子节点变成一样的
            cost[i - 1] += max(cost[i * 2 - 1], cost[i * 2])  # 累加路径和
        return ans
```

```java [sol-Java]
class Solution {
    public int minIncrements(int n, int[] cost) {
        int ans = 0;
        for (int i = n / 2; i > 0; i--) { // 从最后一个非叶节点开始算
            ans += Math.abs(cost[i * 2 - 1] - cost[i * 2]); // 两个子节点变成一样的
            cost[i - 1] += Math.max(cost[i * 2 - 1], cost[i * 2]); // 累加路径和
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minIncrements(int n, vector<int> &cost) {
        int ans = 0;
        for (int i = n / 2; i > 0; i--) { // 从最后一个非叶节点开始算
            ans += abs(cost[i * 2 - 1] - cost[i * 2]); // 两个子节点变成一样的
            cost[i - 1] += max(cost[i * 2 - 1], cost[i * 2]); // 累加路径和
        }
        return ans;
    }
};
```

```go [sol-Go]
func minIncrements(n int, cost []int) (ans int) {
	for i := n / 2; i > 0; i-- { // 从最后一个非叶节点开始算
		left, right := cost[i*2-1], cost[i*2]
		if left > right { // 保证 left <= right
			left, right = right, left
		}
		ans += right - left // 两个子节点变成一样的
		cost[i-1] += right // 累加路径和
	}
	return
}
```

```js [sol-JavaScript]
var minIncrements = function(n, cost) {
    let ans = 0;
    for (let i = Math.floor(n / 2); i > 0; i--) { // 从最后一个非叶节点开始算
        ans += Math.abs(cost[i * 2 - 1] - cost[i * 2]); // 两个子节点变成一样的
        cost[i - 1] += Math.max(cost[i * 2 - 1], cost[i * 2]); // 累加路径和
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_increments(n: i32, mut cost: Vec<i32>) -> i32 {
        let mut ans = 0;
        for i in (1..=n as usize / 2).rev() { // 从最后一个非叶节点开始算
            ans += (cost[i * 2 - 1] - cost[i * 2]).abs(); // 两个子节点变成一样的
            cost[i - 1] += cost[i * 2 - 1].max(cost[i * 2]); // 累加路径和
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{cost}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 思考题

如果还可以对节点值**减一**要怎么做？

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
