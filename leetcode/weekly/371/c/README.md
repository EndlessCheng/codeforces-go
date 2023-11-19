请看 [视频讲解](https://www.bilibili.com/video/BV1MG411X7zR/)。

总共就两种情况：

1. 不交换 $\textit{nums}_1[n-1]$ 和 $\textit{nums}_2[n-1]$。
2. 交换 $\textit{nums}_1[n-1]$ 和 $\textit{nums}_2[n-1]$。

对于每种情况，从 $i=0$ 枚举到 $i=n-2$，一旦发现 $\textit{nums}_1[i] > \textit{nums}_1[n-1]$ 或 $\textit{nums}_2[i] > \textit{nums}_2[n-1]$，就**必须**执行交换操作。如果操作后仍然满足 $\textit{nums}_1[i] > \textit{nums}_1[n-1]$ 或 $\textit{nums}_2[i] > \textit{nums}_2[n-1]$，说明这种情况无法满足要求。

如果两种情况都无法满足要求，返回 $-1$。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums1: List[int], nums2: List[int]) -> int:
        def f(last1: int, last2: int) -> int:
            res = 0
            for x, y in zip(nums1, nums2):
                if x > last1 or y > last2:
                    if y > last1 or x > last2:
                        return inf
                    res += 1
            return res
        ans = min(f(nums1[-1], nums2[-1]), f(nums2[-1], nums1[-1]))
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums1, int[] nums2) {
        int n = nums1.length;
        int ans = Math.min(f(nums1[n - 1], nums2[n - 1], nums1, nums2),
                       1 + f(nums2[n - 1], nums1[n - 1], nums1, nums2));
        return ans > n ? -1 : ans;
    }

    private int f(int last1, int last2, int[] nums1, int[] nums2) {
        int res = 0;
        for (int i = 0; i + 1 < nums1.length; ++i) {
            int x = nums1[i], y = nums2[i];
            if (x > last1 || y > last2) {
                if (y > last1 || x > last2) {
                    return nums1.length + 1;
                }
                res++;
            }
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int> &nums1, vector<int> &nums2) {
        auto f = [&](int last1, int last2) -> int {
            int res = 0;
            for (int i = 0; i + 1 < nums1.size(); i++) {
                int x = nums1[i], y = nums2[i];
                if (x > last1 || y > last2) {
                    if (y > last1 || x > last2) {
                        return nums1.size() + 1;
                    }
                    res++;
                }
            }
            return res;
        };
        int ans = min(f(nums1.back(), nums2.back()), 1 + f(nums2.back(), nums1.back()));
        return ans > nums1.size() ? -1 : ans;
    }
};
```

```go [sol-Go]
func minOperations(nums1, nums2 []int) int {
	n := len(nums1)
	f := func(last1, last2 int) (res int) {
		for i, x := range nums1[:n-1] {
			y := nums2[i]
			if x > last1 || y > last2 {
				if y > last1 || x > last2 {
					return n + 1
				}
				res++
			}
		}
		return
	}
	ans := min(f(nums1[n-1], nums2[n-1]), 1+f(nums2[n-1], nums1[n-1]))
	if ans > n {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
