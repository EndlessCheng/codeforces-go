分类讨论+二分答案

---

先统计分负数乘积个数 $\textit{neg}$、正数乘积个数 $\textit{pos}$ 以及乘积为 $0$ 的个数 $\textit{zero}$，然后分三种情况讨论：

- $k\le \textit{neg}$，我们可以二分负数答案，统计不超过二分值的乘积个数；
- $\textit{neg}<k\le \textit{neg}+\textit{zero}$，此时返回 $0$；
- $k>\textit{neg}+\textit{zero}$，我们可以二分正数答案，统计不超过二分值的乘积个数；

以最后一种情况为例，记二分值为 $t$，我们可以遍历 $\textit{num1}$ 中的正数，并在 $\textit{num2}$ 的正数中二分（或双指针）不超过 $\dfrac{t}{\textit{num1}[i]}$ 的元素个数，然后遍历 $\textit{num1}$ 中的负数，方法同上。遍历结束后，若元素个数小于 $t$ 则说明二分值偏小，否则偏大。





**注**：存在线性做法，见 [Selection in a sorted matrix](https://chaoxu.prof/posts/2014-04-02-selection-in-a-sorted-matrix.html)，或者论文“Selection in x + y and matrices with sorted rows and columns”。

