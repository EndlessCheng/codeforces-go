说起来，由于「晚上和次日凌晨是相连的」，可能潜意识中会觉得次日凌晨就是今天。比如某人买了一张 8 号凌晨的火车票，在 8 号晚上到火车站候车，结果眼睁睁地看到时间变成了 9 号……

相关问题：

- [如何看待错过前一天凌晨的火车票？](https://www.zhihu.com/question/412332219)
- [为什么日本时间表上会有 25:00？](https://www.zhihu.com/question/20438994)

```py [sol-Python3]
class Solution:
    def findDelayedArrivalTime(self, arrivalTime: int, delayedTime: int) -> int:
        return (arrivalTime + delayedTime) % 24
```

```java [sol-Java]
class Solution {
    public int findDelayedArrivalTime(int arrivalTime, int delayedTime) {
        return (arrivalTime + delayedTime) % 24;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findDelayedArrivalTime(int arrivalTime, int delayedTime) {
        return (arrivalTime + delayedTime) % 24;
    }
};
```

```go [sol-Go]
func findDelayedArrivalTime(arrivalTime, delayedTime int) int {
    return (arrivalTime + delayedTime) % 24
}
```

```js [sol-JavaScript]
var findDelayedArrivalTime = function(arrivalTime, delayedTime) {
    return (arrivalTime + delayedTime) % 24;
};
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
