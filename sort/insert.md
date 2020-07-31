## 插入排序
> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍
1. 从`index = 0`开始，将第一个数作为已排序序列，将`index = 1`的数与已排序的序列进行比较，如果比`index=0`小，则插入到其前方，否则插入其后方
2. 循环执行第`1`步：从`index = k`开始，将`0`~`k`个数作为已排序序列，将`index = k + 1`的数与已排序序列的每一项一一进行比较，插入到合适的位置
3. 排序完成

### 2. 算法示例
```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
 function sort (nums) {
     for (let i = 1; i < nums.length; i++) {
         for (let j = 0; j < i; j++) {
             if (nums[i] < nums[j]) {
                 // 将 index=i 的项插入到 index=j 处
                 pop = nums.splice(i, 1)[0]
                 nums.splice(j, 0, pop)
             }
         }
     }
     return nums
 }

```
