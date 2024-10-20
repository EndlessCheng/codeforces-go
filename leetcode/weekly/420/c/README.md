定义 $\text{LPF}(x)$ 为 $x$ 的**最小质因子**。规定 $\text{LPF}(1)=1$。

- 如果 $\text{LPF}(x)=x$，说明 $x$ 是 $1$ 或者质数，无法变小。
- 如果 $\text{LPF}(x)<x$，说明 $x$ 是合数，可以变小。由于题目规定只能除以最大真因数，我们可以把 $x$ 除以 $\dfrac{x}{\text{LPF}(x)}$，得到 $\text{LPF}(x)$。

贪心，最后一个数肯定无需减少，所以我们从 $i=n-2$ 开始倒着遍历 $\textit{nums}$：

- 如果 $\textit{nums}[i] > \textit{nums}[i+1]$，那么把 $\textit{nums}[i]$ 更新为 $\text{LPF}(\textit{nums}[i])$，操作次数加一。注意更新后 $\textit{nums}[i]$ 一定是质数或 $1$，无法再变小。
- 更新后，如果 $\textit{nums}[i] > \textit{nums}[i+1]$ 仍然成立，说明无法把 $\textit{nums}$ 变成非降的，返回 $-1$。

代码实现时，由于 $1$ 不比其他数大，所以无需初始化 $\text{LPF}(1)=1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1UcyYY4EnQ/) 第三题，欢迎点赞关注~

```py [sol-Python3]
MX = 1_000_001
LPF = [0] * MX
for i in range(2, MX):
    if LPF[i] == 0:
        for j in range(i, MX, i):
            if LPF[j] == 0:
                LPF[j] = i

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        for i in range(len(nums) - 2, -1, -1):
            if nums[i] > nums[i + 1]:
                nums[i] = LPF[nums[i]]
                if nums[i] > nums[i + 1]:
                    return -1
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MX = 1_000_001;
    private static final int[] lpf = new int[MX];

    static {
        for (int i = 2; i < MX; i++) {
            if (lpf[i] == 0) {
                for (int j = i; j < MX; j += i) {
                    if (lpf[j] == 0) {
                        lpf[j] = i;
                    }
                }
            }
        }
    }

    public int minOperations(int[] nums) {
        int ans = 0;
        for (int i = nums.length - 2; i >= 0; i--) {
            if (nums[i] > nums[i + 1]) {
                nums[i] = lpf[nums[i]];
                if (nums[i] > nums[i + 1]) {
                    return -1;
                }
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
const int MX = 1'000'001;
int LPF[MX];

auto init = [] {
    for (int i = 2; i < MX; i++) {
        if (LPF[i] == 0) {
            for (int j = i; j < MX; j += i) {
                if (LPF[j] == 0) {
                    LPF[j] = i;
                }
            }
        }
    }
    return 0;
}();

class Solution {
public:
    int minOperations(vector<int>& nums) {
        int ans = 0;
        for (int i = nums.size() - 2; i >= 0; i--) {
            if (nums[i] > nums[i + 1]) {
                nums[i] = LPF[nums[i]];
                if (nums[i] > nums[i + 1]) {
                    return -1;
                }
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
const mx = 1_000_001
var lpf = [mx]int{}

func init() {
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func minOperations(nums []int) (ans int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			nums[i] = lpf[nums[i]]
			if nums[i] > nums[i+1] {
				return -1
			}
			ans++
		}
	}
	return
}
```

#### 复杂度分析

预处理的时间复杂度为 $\mathcal{O}(U\log\log U)$，其中 $U=10^6$。注：用欧拉筛（线性筛）可以做到 $\mathcal{O}(U)$。

对于 $\texttt{minOperations}$：

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面贪心题单中的「**§1.4 从最左/最右开始贪心**」，以及数学题单中的「**§1.3 质因数分解**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
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
