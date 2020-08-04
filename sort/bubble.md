## 冒泡排序
> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍

1. 第一轮比较
> 确定第`n`个数为最大值

从队首`index = 0`开始，依次比较相邻`2`个数字的大小，如果`index = 0`大于`index + 1`, 则交换两个数字的位置，一直到`index = n - 2`和`index = n - 1`比较完毕。
2. 进行剩余的`n - 1`轮比较
> 每次确定第`k`个数(`1 <= k < n`) 为剩余值的最大值

循环第一步
3. 排序完毕

### 2. 算法示例
代码：
```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
function sort(nums) {
    for (let i = 0; i < nums.length - 1; i++) {
        let alreadySort = true // 用以判断是否已经排序完毕
        for (let j = 1; j < nums.length - i; j++) {
            if (nums[j - 1] > nums[j]) {
                const temp = nums[j - 1]
                nums[j - 1] = nums[j]
                nums[j] = temp
                alreadySort = false // 有交换记录，未排序号
            }
        }
        if (alreadySort) break; // 无交换记录，剩余的已排好
    }
    return nums
}
```
