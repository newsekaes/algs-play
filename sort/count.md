## 计数排序
> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现


### 1. 算法介绍
1. 对待排序的数组A进行一次`O(n)`遍历,找出其中的最大值`max`和最小值`min`
2. 用一组长度为`max - min + 1`的数组`B`，遍历数组A，得到记录每个数出现次数后的数组`B`, 来记录数组`A`之中每个数出现的次数
3. 按照每个数出现的次数，遍历数组`B`，得到排序完成的数组C

### 2. 算法示例
代码：
```javascript
/**
 * @param { number[] } nums     
 * @return { number[] }
 */
function sort(nums) {
    const [min, max] = getMinMax(nums)
    const minMaxArray = new Array(max - min + 1).fill(0)
    const sortArray = []
    for (let i = 0; i < nums.length; i++) {
        minMaxArray[nums[i] - min]++
    }
    for (let i = 0; i < minMaxArray.length; i++) {
        for (let j = minMaxArray[i]; j > 0; j--) {
            sortArray.push(i + min)
        }
    }
    return sortArray
}
function getMinMax(nums) {
    let min = nums[0]
    let max = nums[0]
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] > max) {
            max = nums[i]
        }
        else if (nums[i] < min) {
            min = nums[i]
        }   
    }
    return [min, max]
}
```
