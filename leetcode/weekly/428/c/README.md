## 分析

设第二段的第一个数的下标为 $i$，第三段的第一个数的下标为 $j$。

如果第一段是第二段的前缀，那么必须满足：

- 第一段的长度不超过第二段的长度，即 $i \le j-i$。
- 子数组 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 等于子数组 $\textit{nums}[i]$ 到 $\textit{nums}[2i-1]$。

第二、三段同理。

## 方法一：LCP 数组

为方便判断子数组是否相等，定义 $\textit{lcp}[i][j]$ 表示后缀 $\textit{nums}[i:]$ 和后缀 $\textit{nums}[j:]$ 的**最长公共前缀**（Longest Common Prefix）的长度。

考虑递推。分类讨论：

- 如果 $\textit{nums}[i]\ne \textit{nums}[j]$，那么 $\textit{lcp}[i][j] = 0$。
- 如果 $\textit{nums}[i]= \textit{nums}[j]$，那么问题变成计算后缀 $\textit{nums}[i+1:]$ 和后缀 $\textit{nums}[j+1:]$ 的最长公共前缀的长度，那么 $\textit{lcp}[i][j] = \textit{lcp}[i+1][j+1]+1$。

初始值 $\textit{lcp}[n][j] = \textit{lcp}[i][n] = 0$。

如果第一段是第二段的前缀，那么必须满足：

- 第一段的长度 $i$ 不超过第二段的长度 $j-i$，即 $i \le j-i$。
- $\textit{nums}$ 和后缀 $\textit{nums}[i:]$ 的最长公共前缀的长度至少是第一段的长度，即 $\textit{lcp}[0][i]\ge i$。

如果第二段是第三段的前缀，那么必须满足：

- 第二段的长度 $j-i$ 不超过第三段的长度 $n-j$，即 $j-i \le n-j$。
- 后缀 $\textit{nums}[i:]$ 和后缀 $\textit{nums}[j:]$ 的最长公共前缀的长度至少是第二段的长度，即 $\textit{lcp}[i][j]\ge j-i$。
- 实际上，如果 $\textit{lcp}[i][j]\ge j-i$ 成立，那么第三段的长度必然 $\ge j-i$，所以无需判断 $j-i \le n-j$。

```py [sol-Python3]
class Solution:
    def beautifulSplits(self, nums: List[int]) -> int:
        n = len(nums)
        # lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        lcp = [[0] * (n + 1) for _ in range(n + 1)]
        for i in range(n - 1, -1, -1):
            for j in range(n - 1, i - 1, -1):
                if nums[i] == nums[j]:
                    lcp[i][j] = lcp[i + 1][j + 1] + 1

        ans = 0
        for i in range(1, n - 1):
            for j in range(i + 1, n):
                if i <= j - i and lcp[0][i] >= i or lcp[i][j] >= j - i:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int beautifulSplits(int[] nums) {
        int n = nums.length;
        // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        int[][] lcp = new int[n + 1][n + 1];
        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j >= i; j--) {
                if (nums[i] == nums[j]) {
                    lcp[i][j] = lcp[i + 1][j + 1] + 1;
                }
            }
        }

        int ans = 0;
        for (int i = 1; i < n - 1; i++) {
            for (int j = i + 1; j < n; j++) {
                if (i <= j - i && lcp[0][i] >= i || lcp[i][j] >= j - i) {
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int beautifulSplits(vector<int>& nums) {
        int n = nums.size();
        // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        vector lcp(n + 1, vector<int>(n + 1));
        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j >= i; j--) {
                if (nums[i] == nums[j]) {
                    lcp[i][j] = lcp[i + 1][j + 1] + 1;
                }
            }
        }

        int ans = 0;
        for (int i = 1; i < n - 1; i++) {
            for (int j = i + 1; j < n; j++) {
                if (i <= j - i && lcp[0][i] >= i || lcp[i][j] >= j - i) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func beautifulSplits(nums []int) (ans int) {
	n := len(nums)
	// lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if i <= j-i && lcp[0][i] >= i || lcp[i][j] >= j-i {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二：Z 数组（扩展 KMP）

把空间复杂度优化到 $\mathcal{O}(n)$。时间复杂度不变（常数可能更大一些）。

对于第一、二段，可以计算 $\textit{nums}$ 的 $z$ 数组，其中 $z[i]$ 就是方法一中的 $\textit{lcp}[0][i]$。

对于第二、三段，可以计算 $\textit{nums}[i:]$ 的 $z$ 数组，其中 $z[j-i]$ 就是方法一中的 $\textit{lcp}[i][j]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1pnqZYKEqr/?t=28m05s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def calc_z(self, s: List[int]) -> list[int]:
        n = len(s)
        z = [0] * n
        box_l = box_r = 0  # z-box 左右边界
        for i in range(1, n):
            if i <= box_r:
                # 手动 min，加快速度
                x = z[i - box_l]
                y = box_r - i + 1
                z[i] = x if x < y else y
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                box_l, box_r = i, i + z[i]
                z[i] += 1
        return z

    def beautifulSplits(self, nums: List[int]) -> int:
        z0 = self.calc_z(nums)
        n = len(nums)
        ans = 0
        for i in range(1, n - 1):
            z = self.calc_z(nums[i:])
            for j in range(i + 1, n):
                if i <= j - i and z0[i] >= i or z[j - i] >= j - i:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int beautifulSplits(int[] nums) {
        int n = nums.length;
        int ans = 0;
        int[] z0 = calcZ(nums, 0);
        for (int i = 1; i < n - 1; i++) {
            int[] z = calcZ(nums, i);
            for (int j = i + 1; j < n; j++) {
                if (i <= j - i && z0[i] >= i || z[j - i] >= j - i) {
                    ans++;
                }
            }
        }
        return ans;
    }

    private int[] calcZ(int[] s, int start) {
        int n = s.length - start;
        int[] z = new int[n];
        int boxL = 0;
        int boxR = 0; // z-box 左右边界
        for (int i = 1; i < n; i++) {
            if (i <= boxR) {
                z[i] = Math.min(z[i - boxL], boxR - i + 1);
            }
            while (i + z[i] < n && s[start + z[i]] == s[start + i + z[i]]) {
                boxL = i;
                boxR = i + z[i];
                z[i]++;
            }
        }
        return z;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc_z(vector<int>& s, int start) {
        int n = s.size() - start;
        vector<int> z(n); // 注意这样会每次创建一个新的 vector，复用的写法见写法二
        int box_l = 0, box_r = 0; // z-box 左右边界
        for (int i = 1; i < n; i++) {
            if (i <= box_r) {
                z[i] = min(z[i - box_l], box_r - i + 1);
            }
            while (i + z[i] < n && s[start + z[i]] == s[start + i + z[i]]) {
                box_l = i;
                box_r = i + z[i];
                z[i]++;
            }
        }
        return z;
    }

public:
    int beautifulSplits(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        vector<int> z0 = calc_z(nums, 0);
        for (int i = 1; i < n - 1; i++) {
            vector<int> z = calc_z(nums, i);
            for (int j = i + 1; j < n; j++) {
                if (i <= j - i && z0[i] >= i || z[j - i] >= j - i) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```cpp [sol-C++ 写法二]
class Solution {
    void calc_z(vector<int>& z, vector<int>& s, int start) {
        int n = s.size() - start;
        int box_l = 0, box_r = 0; // z-box 左右边界
        for (int i = 1; i < n; i++) {
            if (i <= box_r) {
                z[i] = min(z[i - box_l], box_r - i + 1);
            } else {
                z[i] = 0;
            }
            while (i + z[i] < n && s[start + z[i]] == s[start + i + z[i]]) {
                box_l = i;
                box_r = i + z[i];
                z[i]++;
            }
        }
    }

public:
    int beautifulSplits(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        vector<int> z0(n), z(n - 1);
        calc_z(z0, nums, 0);
        for (int i = 1; i < n - 1; i++) {
            calc_z(z, nums, i);
            for (int j = i + 1; j < n; j++) {
                if (i <= j - i && z0[i] >= i || z[j - i] >= j - i) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func calcZ(s []int) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

func beautifulSplits(nums []int) (ans int) {
	n := len(nums)
	z0 := calcZ(nums)
	for i := 1; i < n-1; i++ {
		z := calcZ(nums[i:])
		for j := i + 1; j < n; j++ {
			if i <= j-i && z0[i] >= i || z[j-i] >= j-i {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果要分成 $4$ 段呢？分成 $k$ 段呢？

## 相似题目（LCP 数组）

- [2430. 对字母串可执行的最大删除数](https://leetcode.cn/problems/maximum-deletions-on-a-string/) 2102
- [1977. 划分数字的方案数](https://leetcode.cn/problems/number-of-ways-to-separate-numbers/) 2817

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. 【本题相关】[字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
