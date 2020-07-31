## 选择排序
> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现


> 选择排序无法提前预判排序是否结束，因为即使某一轮循环发生过交换，仅能证明这一组的最大值刚好已经排好位置，而不能判断剩余值的位置排序

### 1. 算法介绍
1. 第一轮比较
> 确定`index = 0`为最小值

从队首`index = 0`开始，依次比较与其后面每个数字的大小，如果`index = 0`大于`k` (`index< k < n - 1`), 则交换两个数字的位置，一直到`index = 0`和`index = n`比较完毕。
2. 进行剩余的`n - 2`轮比较
> 每次确定`index = k` (`index< k < n - 1`) 为次小值

循环第一步：从`index = 1`、`index = 2`……直到`index = n - 2`
3. 排序完毕

### 2. 算法示例

代码：
```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
function sort(nums){
    for (let i = 0; i < nums.length - 1; i++) {
        let minIndex = i
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[minIndex] > nums[j]) {
                minIndex = j
            }
        }
        if (minIndex !== i) {
            const temp = nums[i]
            nums[i] = nums[minIndex]
            nums[minIndex] = temp
        }
    }
    return nums
}

```


