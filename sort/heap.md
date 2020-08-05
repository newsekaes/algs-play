## 堆排序
> 为了方便阅读和书写：
> 1. 排序后的内容为从小到大
> 2. 总数组长度为 `n`
> 3. 以下代码均用Javascript实现

### 1. 算法介绍
> **堆**只需保证父节点大于子节点即可，兄弟节点的大小关系无需考虑
1. 先将整个序列排成一个**大顶堆**
2. 将**堆顶**元素和**堆尾**元素互换，从而将**最大值**沉到最后
3. 循环步骤`1`和步骤`2`，将前方的序列重新排成大顶堆，同样将堆顶元素和堆尾元素交换
4. 当所有项均下沉完毕后，排序结束

### 2. 算法示例
```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
function sort(nums) {

    // 构建一个大顶堆
    for (let i = Math.floor(nums.length / 2) - 1; i >=0; i--) {
        // 因是自底向顶，故heapManage在递归时，子节点已经是后台节点中的最大值
        heapManage(nums, i, nums.length)
    }
    // 不断将堆顶下沉，构建有序列
    for (let i = nums.length - 1; i >=0; i--) {
        swap(nums, 0, i)
        // 相当于将 nums[0] 逐渐下沉直到最后一层
        heapManage(nums, 0, i)
    }
    return nums
}
function heapManage(nums, i, length) {
    const left = 2 * i + 1
    const right = 2 * i + 2
    let largest = i
    if (left < length && nums[left] > nums[largest]) {
        largest = left
    }
    if (right < length && nums[right] > nums[largest]) {
        largest = right
    }
    if (largest !== i) {
        swap(nums, largest, i)
        // 如果递归成功，说明子节点及其后代节点 原先已经经过整理，此时针对被替换的nums[largest]微调，使其下沉到应该存在的位置
        heapManage(nums, largest, length)
    }
}
function swap(nums, i, j) {
    const temp = nums[j]
    nums[j] = nums[i]
    nums[i] = temp
  
}
```
