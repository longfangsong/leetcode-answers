use std::cmp;

impl Solution {
    pub fn min_distance(word1: String, word2: String) -> i32 {
      let mut old_min_distance = vec![0; word1.len() + 1];
      for i in 0..=word1.len() {
        old_min_distance[i] = i;
      }
      for i in 1..=word2.len() {
        let mut new_min_distance = vec![i; word1.len() + 1];
        for j in 1..=word1.len() {
          if word1.as_bytes()[j-1] == word2.as_bytes()[i-1] {
            new_min_distance[j] = old_min_distance[j-1]; 
          } else {
            new_min_distance[j] = cmp::min(
              new_min_distance[j-1],
              cmp::min(old_min_distance[j],
              old_min_distance[j-1])
            ) + 1;
          }
        }
        old_min_distance = new_min_distance;
      }
      old_min_distance[word1.len()] as i32
    }
}