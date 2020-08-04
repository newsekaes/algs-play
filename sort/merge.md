## 归并排序
> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍
1. 将所有的**n**项两两比较排列(奇数项和空数组比较)，这样形成**n/2**个有序项

2. 将已经有序的**n/2**项两两比较(奇数项和空数组比较)，这样形成**n/4**个有序项

3. 重复以上归并操作，直到最后**2**个有序数组合成**一个**有序数组

4. 排序完毕

### 2. 算法示例

代码：
```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
function sort(nums) {
    if (nums.length < 2) {
        return nums
    }
    const left = nums
    const right = nums.splice(Math.ceil(nums.length / 2), nums.length - 1)
    return merge(sort(left), sort(right))
}
function merge(left, right) {
    const mergeArray = []
    while (left.length > 0 || right.length > 0) {
        if (right.length === 0 || left[0] < right[0]) {
            mergeArray.push(left.shift())
        } else if (left.length === 0 || right[0] <= left[0]) {
            mergeArray.push(right.shift())
        }
    }
    return mergeArray
}
```
