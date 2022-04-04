use std::iter;

impl Solution {
    pub fn missing_rolls(rolls: Vec<i32>, mean: i32, remain_rolls: i32) -> Vec<i32> {
        let expected_sum = mean * (remain_rolls + rolls.len() as i32);
        let current_sum: i32 = rolls.iter().sum();
        let sum_to_generate = expected_sum - current_sum;
        if sum_to_generate == 0 {
            return rolls;
        }
        if sum_to_generate < 0 {
            return vec![];
        }
        let average = sum_to_generate / remain_rolls;
        let remain: i32 = sum_to_generate % remain_rolls;
        if average > 6 || average == 0 || (average == 6 && remain > 0) {
            return vec![];
        }
        iter::repeat(average+1).take(remain as usize)
            .chain(iter::repeat(average).take((remain_rolls-remain) as usize))
            .collect()
    }
}