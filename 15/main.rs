fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    if nums.len() < 3 {
        return vec![];
    }
    let mut sorted_ = nums;
    sorted_.sort();
    let mut result = vec![];
    for i in 0..sorted_.len() - 2 {
        if sorted_[i] + sorted_[i + 1] + sorted_[i + 2] > 0 {
            break;
        }
        if sorted_[i] + sorted_[sorted_.len() - 1] + sorted_[sorted_.len() - 2] < 0 {
            continue;
        }
        if i != 0 && sorted_[i] == sorted_[i - 1] {
            continue;
        }
        let mut left = i + 1;
        let mut right = sorted_.len() - 1;
        while left < right {
            let sum = sorted_[i] + sorted_[left] + sorted_[right];
            if sum == 0 {
                result.push(vec![sorted_[i], sorted_[left], sorted_[right]]);
                while left < right && sorted_[left] == sorted_[left + 1] {
                    left += 1;
                }
                while left < right && sorted_[right] == sorted_[right - 1] {
                    right -= 1;
                }
                left += 1;
                right -= 1;
            } else if sum < 0 {
                left += 1;
            } else {
                right -= 1;
            }
        }
    }
    return result;
}

fn main() {
    three_sum(vec![1, -1, -1, 0]);
}
