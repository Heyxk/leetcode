# [剑指 Offer 59 - I. 滑动窗口的最大值][link] (Hard)

[link]: https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/

给定一个数组 `nums` 和滑动窗口的大小 `k`，请找出所有滑动窗口里的最大值。

**示例:**

```
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```

**提示：**

你可以假设 k 总是有效的，在输入数组 **不为空** 的情况下， `1 ≤ k ≤ nums.length`。

注意：本题与主站 239 题相同： [https://leetcode-cn.com/problems/sliding-window-maximum/](https://lee
tcode-cn.com/problems/sliding-window-maximum/)
