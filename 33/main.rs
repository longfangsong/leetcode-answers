fn search(nums: Vec<i32>, target: i32) -> i32 {
  let mut left = 0;
  let mut right = nums.len()-1;

  while left < right {
    let mid = (left + right) / 2;
    if (nums[0] <= nums[mid] && nums[0] <= target && target <= nums[mid]) ||
    (nums[mid] < nums[0] && target <= nums[mid] && nums[mid] < nums[0]) ||
    (nums[mid] < nums[0] && nums[mid] < nums[0] && nums[0] <= target) {
      right = mid;
    } else {
      left = mid+1;
    }
  }
  if left != right || nums[left] != target {
      return -1;
  }
  return left as i32;
}