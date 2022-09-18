下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

题目本质上是求两个区间的交集区间。

交集区间的右端点等于两个区间右端点的最小值，左端点等于两个区间左端点的最大值。

注意交集区间为空的情况。

```py [sol1-Python3]
DAYS_SUM = list(accumulate((31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31), initial=0))

class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        def f(date: str) -> int:
            return DAYS_SUM[int(date[:2]) - 1] + int(date[3:])
        return max(f(min(leaveAlice, leaveBob)) - f(max(arriveAlice, arriveBob)) + 1, 0)
```

```go [sol1-Go]
var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob string) int {
	f := func(s string) (day int) {
		for _, d := range days[:s[0]&15*10+s[1]&15-1] {
			day += d
		}
		return day + int(s[3]&15*10+s[4]&15)
	}
	ans := f(min(leaveAlice, leaveBob)) - f(max(arriveAlice, arriveBob)) + 1
	if ans < 0 { ans = 0 }
	return ans
}

func min(a, b string) string { if b < a { return b }; return a }
func max(a, b string) string { if b > a { return b }; return a }
```
