**前置知识**：[【算法小课堂】差分数组](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)

1. 初始化 $\textit{dec}=0$，表示需要减一的次数。
2. 设 $x=\textit{batteryPercentages}[i]$，那么该电池的实际百分比为 $x - \textit{dec}$，如果 $x - \textit{dec} > 0$，即 $x > \textit{dec}$，那么后面的数都要减一，根据差分数组的思想，把 $\textit{dec}$ 加一即可。
3. 答案就是 $\textit{dec}$。因为每次遇到 $x > \textit{dec}$ 都把 $\textit{dec}$ 加一，这正是题目要求统计的。

[本题视频讲解](https://www.bilibili.com/video/BV1Lj411s7ga/)

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
    int countTestedDevices(vector<int>& batteryPercentages) {
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
        for x in battery_percentages {
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

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
