### 本题视频讲解

见[【双周赛 97】](https://www.bilibili.com/video/BV1rM4y1X7z9/)的第三题。

### 前置知识：同向双指针

见 [同向双指针【基础算法精讲 01】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

> 注：我一般把窗口大小不固定的叫做**双指针**，窗口大小固定的叫做**滑动窗口**。

### 思路

我们可以强制让第二条线段的右端点恰好落在奖品上，设**第二条**线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，我们需要知道在 $\textit{prizePositions}[\textit{left}]$ 左侧的第一条线段最多可以覆盖多少个奖品。

那么，先想想只有一条线段要怎么做。

使用双指针，设线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，那么当前覆盖的奖品个数为 $\textit{right} - \textit{left} + 1$。

同时，用一个数组 $\textit{pre}[\textit{right}+1]$ 记录线段右端点**不超过** $\textit{prizePositions}[\textit{right}]$ 时最多可以覆盖多少个奖品。下标错开一位是为了方便下面计算。

初始 $\textit{pre}[0]=0$。根据 $\textit{pre}$ 的定义，有

$$
\textit{pre}[\textit{right}+1] = \max(\textit{pre}[\textit{right}],\textit{right} - \textit{left} + 1)
$$

回到第二条线段的计算，根据开头说的，此时最多可以覆盖的奖品数为

$$
\textit{right}-\textit{left}+1+\textit{pre}[\textit{left}]
$$

这里 $\textit{pre}[\textit{left}]$ 表示**第一条**线段右端点**不超过** $\textit{prizePositions}[\textit{left}-1]$ 时最多可以覆盖多少个奖品。

遍历过程中取上式的最大值，即为答案。

由于我们遍历了所有的奖品作为第二条线段的右端点，且通过 $\textit{pre}[\textit{left}]$ 保证第一条线段与第二条线段没有任何交点，且第一条线段覆盖了第二条线段左侧的最多奖品。那么这样遍历后，算出的答案就一定是所有情况中的最大值。

如果脑中没有一幅直观的图像，可以看看 [视频讲解【双周赛 97】](https://www.bilibili.com/video/BV1rM4y1X7z9/)的第三题。

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

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{prizePositions}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $O(n)$。
- 空间复杂度：$O(n)$。

#### 相似题目（同向双指针）

- [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)，[题解](https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/1959540/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-iaks/)
- [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)，[题解](https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/1959532/biao-ti-xia-biao-zong-suan-cuo-qing-kan-k81nh/)
- [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/)，[题解](https://leetcode.cn/problems/subarray-product-less-than-k/solutions/1959538/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-jebq/)
