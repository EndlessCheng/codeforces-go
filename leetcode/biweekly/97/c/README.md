### 前置知识：同向双指针

[【同向双指针+简洁模板】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)，看完你就掌握双指针啦~

> 注：我一般把窗口大小不固定的叫做**双指针**，窗口大小固定的叫做**滑动窗口**。

### 思路

用双指针处理第二条线段，我们可以强制让线段右端点恰好落在奖品上，设第二条线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，我们需要知道在 $\textit{prizePositions}[\textit{left}]$ 左侧的第一条线段最多可以覆盖多少个奖品。

那么，先想想只有一条线段要怎么做。

同样使用双指针，设线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，那么当前覆盖的奖品个数为 $\textit{right} - \textit{left} + 1$。

同时，用一个数组 $\textit{pre}[\textit{right}+1]$ 记录线段右端点**不超过** $\textit{prizePositions}[\textit{right}]$ 时最多可以覆盖多少个奖品。下标错开一位是为了方便计算。

初始 $\textit{pre}[0]=0$。根据 $\textit{pre}$ 的定义，有

$$
\textit{pre}[\textit{right}+1] = \max(\textit{pre}[\textit{right}],\textit{right} - \textit{left} + 1)
$$

回到第二条线段的计算，根据开头说的，此时最多可以覆盖的奖牌数为

$$
\textit{right}-\textit{left}+1+\textit{pre}[\textit{left}]
$$

遍历过程中取上式的最大值，即为答案。

代码实现时，可以用一次遍历完成上述过程。

```py [sol1-Python3]
class Solution:
    def maximizeWin(self, prizePositions: List[int], k: int) -> int:
        pre = [0] * (len(prizePositions) + 1)
        ans = left = 0
        for right, p in enumerate(prizePositions):
            while p - prizePositions[left] > k:
                left += 1
            ans = max(ans, right - left + 1 + pre[left])
            pre[right + 1] = max(pre[right], right - left + 1)
        return ans
```

```java [sol1-Java]
class Solution {
    public int maximizeWin(int[] prizePositions, int k) {
        int ans = 0, left = 0, n = prizePositions.length;
        int[] pre = new int[n + 1];
        for (int right = 0; right < n; right++) {
            while (prizePositions[right] - prizePositions[left] > k) ++left;
            ans = Math.max(ans, right - left + 1 + pre[left]);
            pre[right + 1] = Math.max(pre[right], right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int maximizeWin(vector<int> &prizePositions, int k) {
        int ans = 0, left = 0, n = prizePositions.size(), pre[n + 1];
        pre[0] = 0;
        for (int right = 0; right < n; right++) {
            while (prizePositions[right] - prizePositions[left] > k) ++left;
            ans = max(ans, right - left + 1 + pre[left]);
            pre[right + 1] = max(pre[right], right - left + 1);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func maximizeWin(prizePositions []int, k int) (ans int) {
	pre := make([]int, len(prizePositions)+1)
	left := 0
	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1)
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

#### 相似题目（同向双指针）

- [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)，[题解](https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/1959540/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-iaks/)
- [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)，[题解](https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/1959532/biao-ti-xia-biao-zong-suan-cuo-qing-kan-k81nh/)
- [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/)，[题解](https://leetcode.cn/problems/subarray-product-less-than-k/solutions/1959538/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-jebq/)
