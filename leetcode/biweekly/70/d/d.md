```go [sol1-Go]
func numberOfWays(corridor string) int {
	ans, cntS, pre := 1, 0, 0
	for i, ch := range corridor {
		if ch == 'S' {
			// 对第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意一个位置放置屏风
			cntS++
			if cntS >= 3 && cntS%2 == 1 {
				ans = ans * (i - pre) % (1e9 + 7)
			}
			pre = i // 记录上一个座位的位置
		}
	}
	if cntS == 0 || cntS%2 == 1 { // 座位个数不能为 0 或奇数
		return 0
	}
	return ans
}
```

```C++ [sol1-C++]
class Solution {
public:
    int numberOfWays(string corridor) {
        int ans = 1, cnt_s = 0, pre = 0;
        for (int i = 0; i < corridor.length(); ++i) {
            if (corridor[i] == 'S') {
                // 对第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意一个位置放置屏风
                ++cnt_s;
                if (cnt_s >= 3 && cnt_s % 2) {
                    ans = (long) ans * (i - pre) % 1000000007;
                }
                pre = i; // 记录上一个座位的位置
            }
        }
        return cnt_s && cnt_s % 2 == 0 ? ans : 0; // 座位个数必须为正偶数
    }
};
```

```Python [sol1-Python3]
class Solution:
    def numberOfWays(self, corridor: str) -> int:
        ans, cnt_s, pre = 1, 0, 0
        for i, ch in enumerate(corridor):
            if ch == 'S':
                # 对第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意一个位置放置屏风
                cnt_s += 1
                if cnt_s >= 3 and cnt_s % 2 == 1:
                    ans = ans * (i - pre) % 1000000007
                pre = i  # 记录上一个座位的位置
        return ans if cnt_s and cnt_s % 2 == 0 else 0  # 座位个数必须为正偶数
```

```java [sol1-Java]
class Solution {
    public int numberOfWays(String corridor) {
        var ans = 1L;
        var cntS = 0;
        var pre = 0;
        for (var i = 0; i < corridor.length(); ++i) {
            if (corridor.charAt(i) == 'S') {
                // 对第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意一个位置放置屏风
                ++cntS;
                if (cntS >= 3 && cntS % 2 == 1) {
                    ans = ans * (i - pre) % 1000000007;
                }
                pre = i; // 记录上一个座位的位置
            }
        }
        return cntS > 0 && cntS % 2 == 0 ? (int) ans : 0; // 座位个数必须为正偶数
    }
}
```
