[这场周赛的视频讲解](https://www.bilibili.com/video/BV1BM4y1W7AQ/)

```py [sol-Python3]
class Solution:
    def numberOfEmployeesWhoMetTarget(self, hours: List[int], target: int) -> int:
        return sum(h >= target for h in hours)
```

```java [sol-Java]
class Solution {
    public int numberOfEmployeesWhoMetTarget(int[] hours, int target) {
        int ans = 0;
        for (int h : hours) {
            if (h >= target) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfEmployeesWhoMetTarget(vector<int> &hours, int target) {
        int ans = 0;
        for (int h: hours) {
            ans += h >= target;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfEmployeesWhoMetTarget(hours []int, target int) (ans int) {
	for _, h := range hours {
		if h >= target {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var numberOfEmployeesWhoMetTarget = function(hours, target) {
    let ans = 0;
    for (const h of hours) {
        ans += h >= target;
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{hours}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
