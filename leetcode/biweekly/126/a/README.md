[视频讲解](https://www.bilibili.com/video/BV11x421r7q5/)。

对于 $x=\textit{nums}[i]$，遍历 $x$ 的每个数位，计算其最大数位 $\textit{mx}$。

遍历的同时，计算出一个等于 $111\cdots$ 的值 $\textit{base}$。

那么加密后的结果就是 $\textit{mx}\cdot \textit{base}$，加入答案。

```py [sol-Python3]
class Solution:
    def sumOfEncryptedInt(self, nums: List[int]) -> int:
        ans = 0
        for x in nums:
            mx = base = 0
            while x:
                x, d = divmod(x, 10)
                mx = max(mx, d)
                base = base * 10 + 1
            ans += mx * base
        return ans
```

```java [sol-Java]
class Solution {
    public int sumOfEncryptedInt(int[] nums) {
        int ans = 0;
        for (int x : nums) {
            int mx = 0;
            int base = 0;
            for (; x > 0; x /= 10) {
                mx = Math.max(mx, x % 10);
                base = base * 10 + 1;
            }
            ans += mx * base;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfEncryptedInt(vector<int> &nums) {
        int ans = 0;
        for (int x : nums) {
            int mx = 0, base = 0;
            for (; x; x /= 10) {
                mx = max(mx, x % 10);
                base = base * 10 + 1;
            }
            ans += mx * base;
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumOfEncryptedInt(nums []int) (ans int) {
	for _, x := range nums {
		mx, base := 0, 0
		for ; x > 0; x /= 10 {
			mx = max(mx, x%10)
			base = base*10 + 1
		}
		ans += mx * base
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

更多题单，请点我个人主页 - 讨论发布。
