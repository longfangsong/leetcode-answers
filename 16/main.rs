pub fn three_sum_closest(nums: Vec<i32>, target: i32) -> i32 {
    if nums.len() < 3 {
        return -1;
    }
    let mut sorted_ = nums;
    sorted_.sort();
    let mut result = None;
    for i in 0..sorted_.len() - 2 {
        if i != 0 && sorted_[i] == sorted_[i - 1] {
            continue;
        }
        let mut left = i + 1;
        let mut right = sorted_.len() - 1;
        while left < right {
            let diff = target -  (sorted_[i] + sorted_[left] + sorted_[right]);
            if result.is_none() {
                result = Some(diff)
            }
            if diff == 0 {
                return target
            } else if diff.abs() < result.unwrap().abs() {
                result = Some(diff)
            } else if diff > 0 {
                left += 1;
            } else {
                right -= 1;
            }
        }
    }
    return target - result.unwrap();
}
