## 桶排序

> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍
1. 进行一次`O(n)`的遍历，得到待排序列`A`的最大值`max`和`min`
2. 以`max`和`min`为总区间，每隔长度`5`，划分一个小区间，称之为`桶`
3. 遍历`A`，按每项的大小，划分到对应的`桶`中，同时使用排序（这里使用`插入排序`）
4. 拼接所有的桶，得到有序序列`B`

### 2. 算法示例
```javascript
/**
 * @param { number[] } nums
 * @return { number[] }
 */
function sort(nums) {
    const bucketSize = 10
    // 获取最大最小值
    const [min, max] = getMaxMin(nums)
    // 构建桶组
    const buckets = new Array(Math.ceil((max - min + 1) / bucketSize) + 1).fill(null).map(() => [])
    // 遍历nums，插入桶组
    for (let i = 0; i < nums.length; i++) {
        const bucketIndex = Math.ceil((nums[i] - min + 1) / bucketSize)
        insertAndSort(buckets[bucketIndex], nums[i])
    }
    // 拼接数组
    return [].concat(...buckets)
}
function getMaxMin(nums) {
    let max = nums[0]
    let min = nums[0]
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] > max) {
            max = nums[i]
        } else if (nums[i] < min) {
            min = nums[i]
        }
    }
    return [min, max]
}
function insertAndSort(nums, value) {
    let i = 0
    for (; i < nums.length; i++) {
        if (value < nums[i]) {
            break
        }
    }
    nums.splice(i, 0, value)
}
```
