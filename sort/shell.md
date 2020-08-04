## 希尔排序

> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍
1. 规定一个序列`ti`, `ti-1`, ... ,`1`为一个递减序列；tn 用 `(3*N + 1)/2` 的方式生成，使其刚好 **小于** 数组长度`n`
2. 每隔`ti`个数，将数组分割为共`ti`组数据（此时每组至少有**1**个项），对每组里的所有项进行**插值排序**；
3. 循环执行`1`、`2`：用`ti-1`分割所有数组，然后对每组的数据进行**插值排序**；直到`t = 1`时，相当于对整个数组(已部分完成排序)进行**插值排序**

### 2. 算法示例
代码：
```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
 function sort (nums) {
     let gap = 1
     while (gap < nums.length) {
         gap = 3 * gap + 1
     }
     while (gap > 0) {
         for (let i = gap; i < nums.length; i++) {
             let temp = nums[i]
             let j = i - gap
             for (; nums[j] > temp; j -= gap) {
                 nums[j + gap] = nums[j]
             }
             nums[j + gap] = temp
         }
         gap = Math.floor(gap / 2)
     }
     return nums
 }
```
