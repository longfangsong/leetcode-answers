impl Solution {
    pub fn all_paths_source_target(graph: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut result: Vec<Vec<i32>> = Vec::new();
        let mut pending: std::collections::VecDeque<Vec<i32>> = std::collections::VecDeque::new();
        pending.push_back(vec![0]);
        while !pending.is_empty() {
            let current = pending.pop_front().unwrap();
            let from = *current.last().unwrap() as usize;
            if !graph[from].is_empty() {
                for to in &graph[from] {
                    let mut new_result = current.clone();
                    new_result.push(*to);
                    if *to as usize == graph.len() - 1 {
                        result.push(new_result);
                    } else {
                        pending.push_back(new_result);
                    }
                }
            }
        }
        result
    }
}
