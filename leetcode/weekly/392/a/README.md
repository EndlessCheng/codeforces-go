读完这篇教程你就会做了：[【简单题杀手】分组循环](https://leetcode.cn/problems/longest-alternating-subarray/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-r57bz/)

注意本题和教程中的情况一样，第一个单调序列末尾和第二个单调序列开头，有一个元素是**重叠**的，所以下面代码在外层循环末尾要把 $i$ 减一。

请看 [视频讲解](https://www.bilibili.com/video/BV1ut421H7Wv/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def longestMonotonicSubarray(self, a: List[int]) -> int:
        ans = 1
        i, n = 0, len(a)
        while i < n - 1:
            if a[i + 1] == a[i]:
                i += 1  # 直接跳过
                continue
            i0 = i  # 记录这一组的开始位置
            inc = a[i + 1] > a[i]  # 定下基调：是严格递增还是严格递减
            i += 2  # i 和 i+1 已经满足要求，从 i+2 开始判断
            while i < n and a[i] != a[i - 1] and (a[i] > a[i - 1]) == inc:
                i += 1
            # 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = max(ans, i - i0)
            i -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int longestMonotonicSubarray(int[] a) {
        int ans = 1;
        int i = 0, n = a.length;
        while (i < n - 1) {
            if (a[i + 1] == a[i]) {
                i++; // 直接跳过
                continue;
            }
            int i0 = i; // 记录这一组的开始位置
            boolean inc = a[i + 1] > a[i]; // 定下基调：是严格递增还是严格递减
            i += 2; // i 和 i+1 已经满足要求，从 i+2 开始判断
            while (i < n && a[i] != a[i - 1] && (a[i] > a[i - 1]) == inc) {
                i++;
            }
            // 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = Math.max(ans, i - i0);
            i--;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestMonotonicSubarray(vector<int> &a) {
        int ans = 1;
        int i = 0, n = a.size();
        while (i < n - 1) {
            if (a[i + 1] == a[i]) {
                i++; // 直接跳过
                continue;
            }
            int i0 = i; // 记录这一组的开始位置
            bool inc = a[i + 1] > a[i]; // 定下基调：是严格递增还是严格递减
            i += 2; // i 和 i+1 已经满足要求，从 i+2 开始判断
            while (i < n && a[i] != a[i - 1] && (a[i] > a[i - 1]) == inc) {
                i++;
            }
            // 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = max(ans, i - i0);
            i--;
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestMonotonicSubarray(a []int) int {
	ans := 1
	i, n := 0, len(a)
	for i < n-1 {
		if a[i+1] == a[i] {
			i++ // 直接跳过
			continue
		}
		i0 := i              // 记录这一组的开始位置
		inc := a[i+1] > a[i] // 定下基调：是严格递增还是严格递减
		i += 2               // i 和 i+1 已经满足要求，从 i+2 开始判断
		for i < n && a[i] != a[i-1] && a[i] > a[i-1] == inc {
			i++
		}
		// 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
		ans = max(ans, i-i0)
		i--
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。时间复杂度乍一看是 $\mathcal{O}(n^2)$，但注意变量 $i$ 减少的次数是 $\mathcal{O}(n)$ 的，其它情况一直在增加，由于 $i$ 最大是 $n$，所以增加的次数是 $\mathcal{O}(n)$，所以二重循环总共循环 $\mathcal{O}(n)$ 次，时间复杂度是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 相似题目

- [1446. 连续字符](https://leetcode.cn/problems/consecutive-characters/) 1165
- [1869. 哪种连续子字符串更长](https://leetcode.cn/problems/longer-contiguous-segments-of-ones-than-zeros/) 1205
- [1957. 删除字符使字符串变好](https://leetcode.cn/problems/delete-characters-to-make-fancy-string/) 1358
- [978. 最长湍流子数组](https://leetcode.cn/problems/longest-turbulent-subarray/) 1393
- [2110. 股票平滑下跌阶段的数目](https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/) 1408
- [228. 汇总区间](https://leetcode.cn/problems/summary-ranges/)
- [2760. 最长奇偶子数组](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/) 1420
- [1887. 使数组元素相等的减少操作次数](https://leetcode.cn/problems/reduction-operations-to-make-the-array-elements-equal/) 1428
- [845. 数组中的最长山脉](https://leetcode.cn/problems/longest-mountain-in-array/) 1437
- [2038. 如果相邻两个颜色均相同则删除当前颜色](https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/) 1468
- [1759. 统计同质子字符串的数目](https://leetcode.cn/problems/count-number-of-homogenous-substrings/) 1491
- [3011. 判断一个数组是否可以变为有序](https://leetcode.cn/problems/find-if-array-can-be-sorted/) 1497
- [1578. 使绳子变成彩色的最短时间](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/) 1574
- [1839. 所有元音按顺序排布的最长子字符串](https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/) 1580
- [2765. 最长交替子序列](https://leetcode.cn/problems/longest-alternating-subarray/) 1581
- [467. 环绕字符串中唯一的子字符串](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/) ~1700
- [2948. 交换得到字典序最小的数组](https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/) 2047
- [2393. 严格递增的子数组个数](https://leetcode.cn/problems/count-strictly-increasing-subarrays/)（会员题）
- [2436. 使子数组最大公约数大于一的最小分割数](https://leetcode.cn/problems/minimum-split-into-subarrays-with-gcd-greater-than-one/)（会员题）
- [2495. 乘积为偶数的子数组数](https://leetcode.cn/problems/number-of-subarrays-having-even-product/)（会员题）
- [3063. 链表频率](https://leetcode.cn/problems/linked-list-frequency/)（会员题）

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
