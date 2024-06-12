枚举 $x=\textit{nums}[j]$，我们需要知道有多少个 $\textit{nums}[i]$，满足 $i<j$ 且 $\textit{nums}[i]$ 的最高位与 $x\bmod 10$ 互质。

需要直接枚举 $\textit{nums}[i]$ 吗？有没有更快的做法？

由于 $\textit{nums}[i]$ 的最高位在 $[1,9]$ 中，我们可以在遍历数组的同时，统计最高位的出现次数，这样就只需枚举 $[1,9]$ 中的与 $x\bmod 10$ 互质的数，把对应的出现次数加到答案中。

具体算法如下：

1. 初始化答案 $\textit{ans}=0$，初始化长为 $10$ 的 $\textit{cnt}$ 数组，初始值均为 $0$。
2. 遍历 $\textit{nums}$，设 $x=\textit{nums}[j]$。
3. 枚举 $[1,9]$ 内的数字 $y$，如果与 $x\bmod 10$ 互质，则 $\textit{ans}$ 增加 $\textit{cnt}[y]$。
4. 计算 $x$ 的最高位，将其出现次数加一。
5. 返回 $\textit{ans}$。

[视频讲解](https://www.bilibili.com/video/BV1du41187ZN/)

```py [sol-Python3]
class Solution:
    def countBeautifulPairs(self, nums: List[int]) -> int:
        ans = 0
        cnt = [0] * 10
        for x in nums:
            for y, c in enumerate(cnt):
                if c and gcd(y, x % 10) == 1:
                    ans += c
            while x >= 10: 
                x //= 10
            cnt[x] += 1  # 统计最高位的出现次数
        return ans
```

```java [sol-Java]
class Solution {
    public int countBeautifulPairs(int[] nums) {
        int ans = 0;
        int[] cnt = new int[10];
        for (int x : nums) {
            for (int y = 1; y < 10; y++) {
                if (cnt[y] > 0 && gcd(y, x % 10) == 1) {
                    ans += cnt[y];
                }
            }
            while (x >= 10) {
                x /= 10;
            }
            cnt[x]++; // 统计最高位的出现次数
        }
        return ans;
    }

    private int gcd(int a, int b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countBeautifulPairs(vector<int>& nums) {
        int ans = 0, cnt[10]{};
        for (int x : nums) {
            for (int y = 1; y < 10; y++) {
                if (cnt[y] && gcd(y, x % 10) == 1) {
                    ans += cnt[y];
                }
            }
            while (x >= 10) { 
                x /= 10;
            }
            cnt[x]++; // 统计最高位的出现次数
        }
        return ans;
    }
};
```

```go [sol-Go]
func countBeautifulPairs(nums []int) (ans int) {
	cnt := [10]int{}
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ans += cnt[y]
			}
		}
		for x >= 10 { // 这里需要 O(log x) 的时间
			x /= 10
		}
		cnt[x]++ // 统计最高位的出现次数
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot(k+\log U))$，其中 $n$ 为 $\textit{nums}$ 的长度，$k=10$，$U=\max(\textit{nums})$。计算 GCD 的时间视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(k)$。

注：也可以预处理 $9$ 以内所有数对的 GCD 到一个数组中，从而加速 GCD 的计算。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
