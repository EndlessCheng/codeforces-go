请看 [视频讲解](https://www.bilibili.com/video/BV1tw411q7VZ/) 第二题。

把 $0$ 看成 $1$，设 $s_1$ 为 $\textit{nums}_1$ 的元素和，$s_2$ 为 $\textit{nums}_2$ 的元素和。这是元素和的最小值。

分类讨论：

- 如果 $s_1 < s_2$ 并且 $\textit{nums}_1$ 中没有 $0$，那么 $s_1$ 无法增加，返回 $-1$。
- 如果 $s_2 < s_1$ 并且 $\textit{nums}_2$ 中没有 $0$，那么 $s_2$ 无法增加，返回 $-1$。
- 否则答案为 $\max(s_1,s_2)$。

```py [sol-Python3]
class Solution:
    def minSum(self, nums1: List[int], nums2: List[int]) -> int:
        s1 = sum(max(x, 1) for x in nums1)
        s2 = sum(max(x, 1) for x in nums2)
        if s1 < s2 and 0 not in nums1 or s2 < s1 and 0 not in nums2:
            return -1
        return max(s1, s2)
```

```java [sol-Java]
class Solution {
    public long minSum(int[] nums1, int[] nums2) {
        long s1 = 0;
        boolean zero1 = false;
        for (int x : nums1) {
            if (x == 0) {
                zero1 = true;
                s1++;
            } else {
                s1 += x;
            }
        }

        long s2 = 0;
        boolean zero2 = false;
        for (int x : nums2) {
            if (x == 0) {
                zero2 = true;
                s2++;
            } else {
                s2 += x;
            }
        }

        if (!zero1 && s1 < s2 || !zero2 && s2 < s1) {
            return -1;
        }
        return Math.max(s1, s2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minSum(vector<int>& nums1, vector<int>& nums2) {
        long long s1 = 0;
        bool zero1 = false;
        for (int x : nums1) {
            if (x == 0) {
                zero1 = true;
                s1++;
            } else {
                s1 += x;
            }
        }

        long long s2 = 0;
        bool zero2 = false;
        for (int x : nums2) {
            if (x == 0) {
                zero2 = true;
                s2++;
            } else {
                s2 += x;
            }
        }

        if (!zero1 && s1 < s2 || !zero2 && s2 < s1) {
            return -1;
        }
        return max(s1, s2);
    }
};
```

```go [sol-Go]
func minSum(nums1, nums2 []int) int64 {
	s1 := int64(0)
	zero1 := false
	for _, x := range nums1 {
		if x == 0 {
			zero1 = true
			s1++
		} else {
			s1 += int64(x)
		}
	}

	s2 := int64(0)
	zero2 := false
	for _, x := range nums2 {
		if x == 0 {
			zero2 = true
			s2++
		} else {
			s2 += int64(x)
		}
	}

	if !zero1 && s1 < s2 || !zero2 && s2 < s1 {
		return -1
	}
	return max(s1, s2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
