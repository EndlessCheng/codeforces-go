### 视频讲解

直接看[【周赛 316】](https://www.bilibili.com/video/BV1ne4y1e7nu) 第一题，画个图你就明白了。

> APP 用户如果无法打开，可以分享到微信。

```py [sol1-Python3]
class Solution:
    def haveConflict(self, event1: List[str], event2: List[str]) -> bool:
        return event1[0] <= event2[1] and event1[1] >= event2[0]
```

```java [sol1-Java]
class Solution {
    public boolean haveConflict(String[] event1, String[] event2) {
        return event1[0].compareTo(event2[1]) <= 0 && event1[1].compareTo(event2[0]) >= 0;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool haveConflict(vector<string> &event1, vector<string> &event2) {
        return event1[0] <= event2[1] && event1[1] >= event2[0];
    }
};
```

```go [sol1-Go]
func haveConflict(event1, event2 []string) bool {
	return event1[0] <= event2[1] && event1[1] >= event2[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

#### 相似题目

- [2409. 统计共同度过的日子数](https://leetcode.cn/problems/count-days-spent-together/)

[往期每日一题题解](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
