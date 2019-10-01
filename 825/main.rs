struct Solution;

impl Solution {
    pub fn num_friend_requests(ages: Vec<i32>) -> i32 {
        let mut age_people_count = [0usize; 121];
        for age in ages {
          age_people_count[age as usize] += 1;
        }
        let mut cumsum = [0usize; 121];
        for i in 1..121 {
          cumsum[i] = cumsum[i-1] + age_people_count[i];
        }
        // thus, people count in age range [a,b]
        // is cumsum[b] - cumsum[a]
        let mut result = 0;
        for age in 1..121 {
          let min_age = (0.5 * age as f64 + 7.0) as usize;
          let max_age = age;
          
          if min_age > max_age {
            continue;
          } else { 
            let peoples_can_send_to_this_age = 
              if min_age < age && age <= max_age && cumsum[max_age] != cumsum[min_age] {
                (cumsum[max_age] - cumsum[min_age] - 1)
              } else {
                (cumsum[max_age] - cumsum[min_age])
              };
            result += peoples_can_send_to_this_age * age_people_count[age];
          }
        }
        return result as i32;
    }
}

fn main() {
  println!("{}", Solution::num_friend_requests(vec![16,16]));
}