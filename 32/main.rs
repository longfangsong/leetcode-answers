fn longest_valid_parentheses(s: String) -> i32 {
  let mut left_count = 0;
  let mut right_count = 0;
  let mut result = 0;
  for ch in s.chars() {
    if ch == '(' {
      left_count+=1;
    } else {
      right_count+=1;
    }
    if left_count == right_count && right_count > result {
      result = right_count;
    } else if right_count > left_count {
      left_count = 0;
      right_count = 0;
    }
  }
  left_count = 0;
  right_count = 0;
  for ch in s.chars().rev() {
    if ch == '(' {
      left_count+=1;
    } else {
      right_count+=1;
    }
    if left_count == right_count && left_count > result {
      result = right_count;
    } else if left_count > right_count {
      left_count = 0;
      right_count = 0;
    }                                                     
  }
  return result * 2
}