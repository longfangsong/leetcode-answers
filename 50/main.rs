struct Solution;

impl Solution {
  fn my_pow_abs(x:f64, n: u32) -> f64 {
    let mut result = 1f64;
    let mut x_contribute = x;
    let offset = 0usize;
    for offset in 0..32 {
      if (1 << offset) & n != 0 {
        result *= x_contribute;
      }
      x_contribute *= x_contribute;
    }
    result
  }
  pub fn my_pow(x: f64, n: i32) -> f64 {
    if n > 0 {
      Solution::my_pow_abs(x, n as _)
    } else {
      1.0 / Solution::my_pow_abs(x, -n as _)
    }
  }
}