## 视频讲解

见[【周赛 317】](https://www.bilibili.com/video/BV1Em4y1c7Hc)第一题，欢迎点赞投币~

## 思路

能被 $3$ 和 $2$ 整除的数，就是能被最小公倍数 $\text{LCM}(3,2)=6$ 整除的数。

遍历 $\textit{nums}$，统计 $6$ 的倍数的和 $\textit{sum}$，以及 $6$ 的倍数的个数 $\textit{cnt}$。

答案就是 $\left\lfloor\dfrac{\textit{sum}}{\textit{cnt}}\right\rfloor$。

如果 $\textit{cnt}=0$，返回 $0$。

```py [sol-Python3]
class Solution:
    def averageValue(self, nums: List[int]) -> int:
        sum = cnt = 0
        for x in nums:
            if x % 6 == 0:
                sum += x
                cnt += 1
        return sum // cnt if cnt else 0
```

```java [sol-Java]
class Solution {
    public int averageValue(int[] nums) {
        int sum = 0, cnt = 0;
        for (int x : nums) {
            if (x % 6 == 0) {
                sum += x;
                cnt++;
            }
        }
        return cnt > 0 ? sum / cnt : 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int averageValue(vector<int> &nums) {
        int sum = 0, cnt = 0;
        for (int x : nums) {
            if (x % 6 == 0) {
                sum += x;
                cnt++;
            }
        }
        return cnt ? sum / cnt : 0;
    }
};
```

```go [sol-Go]
func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, x := range nums {
		if x%6 == 0 {
			sum += x
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

---

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
