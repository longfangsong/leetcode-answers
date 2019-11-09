impl Solution {
  pub fn min_path_sum(grid : Vec<Vec<i32>>)->i32 {
    if grid.is_empty() || grid[0].is_empty() {
      return 0;
    }
    let mut current_sum = grid[0][0];
    let mut last_row = vec![grid[0][0]];
    for j in 1..grid[0].len() {
      current_sum += grid[0][j];
      last_row.push(current_sum);
    }
    for i in 1..grid.len() {
      let mut new_row = vec![last_row[0] + grid[i][0]];
      for j in 1..grid[i].len() {
        new_row.push(i32::min(new_row[j-1], last_row[j]) + grid[i][j]);
      }
      last_row = new_row;
    }
    last_row.last().unwrap().to_owned()
  }
}