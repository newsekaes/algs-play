## 基数排序

> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍
1. 获取待排序列`A`(项均为`int`型)的**最大值**和**最小值**，并获得其中**位数**最多的值，得到最大位数`k`
2. 考虑到**负数**的存在，构建一个`-9 ~ 9`的长度为**19**的**桶**组
3. 依次遍历`A`的每个项，并按其**个对**，依次步骤`2`构建的桶组中（区分正负），然后从桶组依次取出，得到`个位`有序序列`A'`， 替代`A`进入下一步
4. 依次按**百位**、**千位**……**第k位**的顺序，循环步骤`3`, 最终得到一个按`第k位`有序序列`A'k`
5. 此时`A'k`为有序序列

### 2. 算法示例
```javascript
/**
 * 快速排序
 * @param { number[] } nums
 * @return { number[] }
 */
function sort(nums) {
    const radixNum = getMaxRadix(nums)
    // 从个位开始，依次按对应位大小排序
    for (let i = 0, d = 10; i < radixNum; i++, d *= 10) {
        const buckets = new Array(19).fill(null).map(() => [])
        // 遍历待排序序列，
        for (let j = 0; j < nums.length; j++) {
            const radixNum = parseInt(nums[j] % d / (d / 10))
            // -9~9 映射为 0~18，需要 +9
            buckets[radixNum + 9].push(nums[j])
        }
        // 依次取出每个bucket并保持bucket内的顺序不变
        nums = [].concat(...buckets)
    }
    return nums
}
function getMaxRadix(nums) {
    let max = nums[0]
    if (max < 0) max = -max
    for (let i = 0; i < nums.length; i++) {
        let temp = nums[i]
        if (temp < 0) temp = -temp
        if (temp > max) {
            max = nums[i]
        }
    }
    if (max === 0) return 1
    let radixNum = 0
    let tenE = 1
    while (max >= tenE ) {
        radixNum++
        tenE *= 10
    }
    return radixNum
}
```
