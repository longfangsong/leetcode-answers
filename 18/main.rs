fn three_sum(sorted: &Vec<i32>, target: i32, from: usize) -> Vec<Vec<i32>> {
    let mut result = vec![];
    let right_two_sum = sorted[sorted.len() - 1] + sorted[sorted.len() - 2];
    for i in from..sorted.len() - 2 {
        if sorted[i] + sorted[i + 1] + sorted[i + 2] > target {
            break;
        }
        if sorted[i] + right_two_sum < target {
            continue;
        }
        if i != from && sorted[i] == sorted[i - 1] {
            continue;
        }
        let mut left = i + 1;
        let mut right = sorted.len() - 1;
        while left < right {
            let current_sum = sorted[i] + sorted[left] + sorted[right];
            if current_sum == target {
                result.push(vec![sorted[i], sorted[left], sorted[right]]);
                while left < right && sorted[left] == sorted[left + 1] {
                    left += 1;
                }
                while left < right && sorted[right] == sorted[right - 1] {
                    right -= 1;
                }
                left += 1;
                right -= 1;
            } else if current_sum < target {
                left += 1;
            } else {
                right -= 1;
            }
        }
    }
    return result;
}

fn four_sum(nums: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
    if nums.len() < 4 {
        return vec![];
    }
    let mut result = vec![];
    let mut sorted = nums;
    sorted.sort_unstable();
    let right_three_sum =
        sorted[sorted.len() - 1] + sorted[sorted.len() - 2] + sorted[sorted.len() - 3];
    for i in 0..sorted.len() - 3 {
        if sorted[i] + sorted[i + 1] + sorted[i + 2] + sorted[i + 3] > target {
            break;
        }
        if sorted[i] + right_three_sum < target {
            continue;
        }
        if i != 0 && sorted[i] == sorted[i - 1] {
            continue;
        }
        let mut three_result = three_sum(&sorted, target - sorted[i], i + 1);
        three_result.iter_mut().for_each(|three_tuple| three_tuple.push(sorted[i]));
        result.append(&mut three_result);
    }
    return result;
}

fn main() {
    println!("{:?}", four_sum(vec![1, 0, -1, 0, -2, 2], 0));
}
