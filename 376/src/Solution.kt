class Solution {
    private fun doWiggleMaxLength(nums: IntArray): Int {
        if (nums.size <= 2) {
            return nums.size
        }
        if ((nums[nums.size - 1] - nums[nums.size - 2]) * (nums[nums.size - 2] - nums[nums.size - 3]) < 0) {
            return doWiggleMaxLength(nums.sliceArray(0 until nums.size - 1)) + 1
        }
        return doWiggleMaxLength(nums.sliceArray(0 until nums.size - 1))
    }
    fun wiggleMaxLength(nums: IntArray): Int {
        if(nums.isEmpty())
            return 0
        val noRepeat = mutableListOf(nums[0])
        for (i in 1 until nums.size) {
            if(nums[i] != nums[i-1]) {
                noRepeat.add(nums[i])
            }
        }
        return doWiggleMaxLength(noRepeat.toIntArray())
    }
}

fun main(args: Array<String>) {
    val s = Solution()
    print(s.wiggleMaxLength(intArrayOf(51,36,245,20,238,238,89,105)))
}