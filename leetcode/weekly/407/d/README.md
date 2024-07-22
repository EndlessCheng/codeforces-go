**请先阅读**：[差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

「子数组内的每个元素的值增加或减少 $1$」，这个操作可以转换成修改**差分数组**两个位置上的数。

设 $d_1$ 为 $\textit{nums}$ 的差分数组，$d_2$ 为 $\textit{target}$ 的差分数组。

由于差分数组和原数组是一一对应的，所以问题等价于：

- 把 $d_1$ 变成 $d_2$。
- 每次操作，可以选择两个下标 $i$ 和 $j$（或者只选一个下标 $i$，对应操作 $\textit{nums}$ 后缀的情况），把 $d_1[i]$ 加一（或减一），把 $d_1[j]$ 减一（或加一）。

从左到右遍历 $d_1$ 和 $d_2$，同时维护一个变量 $s$，表示对 $d_1[i]$ 增大/减少的累积量：

- 如果把 $d_1[i]$ 增大了 $k$，那么后面可以把 $d_1[j]$ **免费减少**，至多免费减少 $k$ 次。
- 如果把 $d_1[i]$ 减少了 $k$，那么后面可以把 $d_1[j]$ **免费增大**，至多免费增大 $k$ 次。

设 $k = d_2[i] - d_1[i]$，分类讨论：

- 如果 $k > 0$ 且 $s\ge 0$，那么必须通过操作，把 $d_1[i]$ 增大到 $d_2[i]$，操作 $k$ 次。
- 如果 $k > 0$ 且 $s < 0$，那么可以免费增大至多 $-s$ 次，如果 $k \le -s$ 则无需额外操作，否则要**额外操作** $k+s$ 次。综合一下，就是额外操作 $\max(s+k,0)$ 次。
- 如果 $k \le 0$ 且 $s\le 0$，那么必须通过操作，把 $d_1[i]$ 减少到 $d_2[i]$，操作 $-k$ 次。
- 如果 $k \le 0$ 且 $s > 0$，那么可以免费减少至多 $s$ 次，如果 $-k \le s$ 则无需额外操作，否则要**额外操作** $-k-s$ 次。综合一下，就是额外操作 $-\min(s+k,0)$ 次。
- 最后把 $k$ 加到 $s$ 中。

代码实现时，可以单独计算 $i=0$ 的情况，方便在计算差分数组的同时计算答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16Z421N7P2/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int], target: List[int]) -> int:
        s = target[0] - nums[0]
        ans = abs(s)
        for (a, b), (c, d) in pairwise(zip(nums, target)):
            k = (d - c) - (b - a)
            if k > 0:
                ans += k if s >= 0 else max(k + s, 0)
            else:
                ans -= k if s <= 0 else min(k + s, 0)
            s += k
        return ans
```

```java [sol-Java]
class Solution {
    public long minimumOperations(int[] nums, int[] target) {
        long s = target[0] - nums[0];
        long ans = Math.abs(s);
        for (int i = 1; i < nums.length; i++) {
            int k = (target[i] - target[i - 1]) - (nums[i] - nums[i - 1]);
            if (k > 0) {
                ans += s >= 0 ? k : Math.max(k + s, 0);
            } else {
                ans -= s <= 0 ? k : Math.min(k + s, 0);
            }
            s += k;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumOperations(vector<int>& nums, vector<int>& target) {
        long long s = target[0] - nums[0];
        long long ans = abs(s);
        for (int i = 1; i < nums.size(); i++) {
            int k = (target[i] - target[i - 1]) - (nums[i] - nums[i - 1]);
            if (k > 0) {
                ans += s >= 0 ? k : max(k + s, 0LL);
            } else {
                ans -= s <= 0 ? k : min(k + s, 0LL);
            }
            s += k;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumOperations(nums, target []int) int64 {
	s := target[0] - nums[0]
	ans := abs(s)
	for i := 1; i < len(nums); i++ {
		k := (target[i] - target[i-1]) - (nums[i] - nums[i-1])
		if k > 0 {
			if s >= 0 {
				ans += k
			} else {
				ans += max(k+s, 0)
			}
		} else {
			if s <= 0 {
				ans -= k
			} else {
				ans -= min(k+s, 0)
			}
		}
		s += k
	}
	return int64(ans)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二


```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int], target: List[int]) -> int:
        a = [t - x for x, t in zip(nums, target)]
        return max(a[0], 0) + sum(max(y - x, 0) for x, y in pairwise(a)) + max(-a[-1], 0)
```

```java [sol-Java]

```

```cpp [sol-C++]

```

```go [sol-Go]

```


## 相似题目

- [1526. 形成目标数组的子数组最少增加次数](https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/)

更多差分题目，见下面的数据结构题单中的「**§2.1 一维差分**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
