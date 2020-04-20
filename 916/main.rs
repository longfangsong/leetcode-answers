use std::cmp::max;

impl Solution {
    fn count_character(content: &str) -> [usize; 26] {
        let mut result = [0; 26];
        for ch in content.chars() {
            result[(ch as u32 - 'a' as u32) as usize] += 1;
        }
        result
    }
    fn max_count(vec: &Vec<String>) -> [usize; 26] {
        let mut result = [0; 26];
        let mut iter = vec.iter();
        let first = iter.next();
        if first.is_none() {
            return result;
        }
        let mut current = Self::count_character(first.unwrap());
        for s in iter {
            let this_count = Self::count_character(s);
            for i in 0..26 {
                current[i] = max(current[i], this_count[i]);
            }
        }
        result
    }
    pub fn word_subsets(a: Vec<String>, b: Vec<String>) -> Vec<String> {
        let need = Self::max_count(&b);
        a.iter().filter(|&it| {
            let characters = Self::count_character(it);
            Iterator::zip(need.iter(), characters.iter())
                .all(|(need, has)| has >= need)
        }).cloned().collect()
    }
}
