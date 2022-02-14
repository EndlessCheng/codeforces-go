由于偶数次翻转等于没有翻转，奇数次翻转等于翻转一次，因此我们可以用一个懒标记 $\textit{flip}$ 来表示当前是否处于翻转状态。

用 $s$ 表示位集，$\textit{cnt}_1$ 表示 $s$ 中 $1$ 的个数，讨论每个操作：

- 对于 $\textit{fix}$ 操作，如果没有发生翻转并且 $s[\textit{idx}]=0$，或者发生翻转并且 $s[\textit{idx}]=1$，那么则翻转 $s[\textit{idx}]$ 的值，将 $\textit{cnt}_1$ 加一。
- 对于 $\textit{unfix}$ 操作，如果没有发生翻转并且 $s[\textit{idx}]=1$，或者发生翻转并且 $s[\textit{idx}]=0$，那么则翻转 $s[\textit{idx}]$ 的值，将 $\textit{cnt}_1$ 减一。
- 对于 $\textit{flip}$ 操作，我们可以不去翻转整个 $s$，仅将懒标记 $\textit{flip}$ 取反，同时 $\textit{cnt}_1$ 置为 $\textit{size}-\textit{cnt}_1$。
- 对于 $\textit{all}$ 查询，判断 $\textit{cnt}_1=\textit{size}$ 是否成立。
- 对于 $\textit{one}$ 查询，判断 $\textit{cnt}_1>0$ 是否成立。
- 对于 $\textit{count}$ 查询，返回 $\textit{cnt}_1$。
- 对于 $\textit{toString}$ 查询，如果没有翻转则直接返回 $s$，否则翻转 $s$ 的每一位并返回。


