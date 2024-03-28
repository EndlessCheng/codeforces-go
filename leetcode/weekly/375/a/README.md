**前置知识**：[【算法小课堂】差分数组](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)

1. 初始化 $\textit{dec}=0$，表示需要减一的次数。
2. 设 $x=\textit{batteryPercentages}[i]$，那么该电池的实际百分比为 $x - \textit{dec}$，如果 $x - \textit{dec} > 0$，即 $x > \textit{dec}$，那么后面的数都要减一，根据差分数组的思想，把 $\textit{dec}$ 加一即可。
3. 答案就是 $\textit{dec}$。因为每次遇到 $x > \textit{dec}$ 都把 $\textit{dec}$ 加一，这正是题目要求统计的。

附：[视频讲解](https://www.bilibili.com/video/BV1Lj411s7ga/)

```py [sol-Python3]
class Solution:
    def countTestedDevices(self, batteryPercentages: List[int]) -> int:
        dec = 0
        for x in batteryPercentages:
            if x > dec:
                dec += 1
        return dec
```

```java [sol-Java]
class Solution {
    public int countTestedDevices(int[] batteryPercentages) {
        int dec = 0;
        for (int x : batteryPercentages) {
            if (x > dec) {
                dec++;
            }
        }
        return dec;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countTestedDevices(vector<int> &batteryPercentages) {
        int dec = 0;
        for (int x : batteryPercentages) {
            dec += x > dec;
        }
        return dec;
    }
};
```

```go [sol-Go]
func countTestedDevices(batteryPercentages []int) int {
	dec := 0
	for _, x := range batteryPercentages {
		if x > dec {
			dec++
		}
	}
	return dec
}
```

```js [sol-JavaScript]
var countTestedDevices = function(batteryPercentages) {
    let dec = 0;
    for (const x of batteryPercentages) {
        if (x > dec) {
            dec++;
        }
    }
    return dec;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_tested_devices(battery_percentages: Vec<i32>) -> i32 {
        let mut dec = 0;
        for &x in &battery_percentages {
            if x > dec {
                dec += 1;
            }
        }
        dec
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{batteryPercentages}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 题单：差分数组

- [1893. 检查是否区域内所有整数都被覆盖](https://leetcode.cn/problems/check-if-all-the-integers-in-a-range-are-covered/) 1307
- [1094. 拼车](https://leetcode.cn/problems/car-pooling/) 1441
- [1109. 航班预订统计](https://leetcode.cn/problems/corporate-flight-bookings/) 1570
- [2406. 将区间分为最少组数](https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups/) 1713
- [2381. 字母移位 II](https://leetcode.cn/problems/shifting-letters-ii/) 1793
- [995. K 连续位的最小翻转次数](https://leetcode.cn/problems/minimum-number-of-k-consecutive-bit-flips/) 1835
- [1943. 描述绘画结果](https://leetcode.cn/problems/describe-the-painting/) 1969
- [2251. 花期内花的数目](https://leetcode.cn/problems/number-of-flowers-in-full-bloom/) 2022
- [2772. 使数组中的所有元素都等于零](https://leetcode.cn/problems/apply-operations-to-make-all-array-elements-equal-to-zero/) 2029
- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/) 2236
- [370. 区间加法](https://leetcode.cn/problems/range-addition/)（会员题）
- [3009. 折线图上的最大交点数量](https://leetcode.cn/problems/maximum-number-of-intersections-on-the-chart/)（会员题）

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
