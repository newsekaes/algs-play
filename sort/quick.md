## 快速排序

> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍
1. 从数组中选择一个**基准数**，所有比**基准数**小的都排在其**左边**，比**基准数**大的都排在其**右边**，进行第一次排列
2. **基准数**左右两边各一组数组，每组数组也从中选择一个**基准数**，按照步骤`1`的方式进行排列
3. 重复步骤`1`和步骤`2`，直到所有的小数组全部排列完
4. 排序完成

### 2. 算法示例
代码：
```javascript
/**
 * 快速排序
 * @param { number[] } nums
 * @return { number[] }
 */
function sort(nums) {
    return quickSort(nums, 0, nums.length - 1)
}
function quickSort(nums, left, right) {
    if (left < right) {
        const pivotIndex = partition(nums, left, right)
        quickSort(nums, left, pivotIndex - 1)
        quickSort(nums,  pivotIndex + 1, right)
    }
    return nums
}
function partition(nums, left, right) {
    // 将最后的项作为基准数
    const pivot = right
    let index = left
    for (let i = index; i < right; i++) {
        if (nums[i] < nums[pivot]) {
            // 每次有比基准值小的项，就排到左侧，并让index+1
            swap(nums, i, index)
            index++
        }
    }
    // left ~ index-1均比基准值要小；index ~ right-1均比基准值大
    // 此时交换index和pivot(即right), 使得index左边的均未较小值，index右边均未较大值
    swap(nums, index, pivot)
    return index
}
function swap(nums, i, j) {
    const temp = nums[j]
    nums[j] = nums[i]
    nums[i] = temp
}
```
