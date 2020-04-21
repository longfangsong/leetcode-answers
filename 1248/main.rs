struct Solution;

impl Solution {
    pub fn number_of_subarrays(nums: Vec<i32>, k: i32) -> i32 {
      let mut odd_count_count = Vec::new();
      odd_count_count.resize(nums.len()+1, 0);
      odd_count_count[0] = 1;
      let mut current_odd = 0;
      let mut result = 0;
      for num in nums {
        if num % 2 == 1 {
          current_odd += 1;
        }
        result += if current_odd >= k { odd_count_count[(current_odd - k) as usize]} else {0};
        odd_count_count[current_odd as usize] += 1;
      }
      result
    }
}

fn main() {
  let a = vec![2, 1, 2, 1, 1, 1, 2, 2, 1 , 1, 1, 1, 2, 2, 2, 1];
  let r = Solution::number_of_subarrays(a, 3);
  println!("{}", r);
}