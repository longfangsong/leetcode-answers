use std::collections::{BTreeMap, HashSet, BTreeSet};
use std::cell::RefCell;
use std::rc::Rc;

#[derive(Clone, Hash, Debug, Ord, PartialOrd, Eq, PartialEq)]
struct Position(i64, i64);

struct DisjointSetUnion<T: Ord + Clone> {
    parent: Rc<RefCell<BTreeMap<T, T>>>
}

impl<T: Ord + Clone> DisjointSetUnion<T> {
    pub fn new() -> Self {
        DisjointSetUnion { parent: Rc::new(RefCell::new(BTreeMap::new())) }
    }
    fn find_root(&self, position: T) -> Option<T> {
        let self_parent = self.parent.clone();
        let self_parent = self_parent.borrow();
        let parent = self_parent.get(&position).cloned();
        drop(self_parent);
        if let Some(parent) = parent {
            if parent == position {
                Some(parent)
            } else {
                let root = self.find_root(parent.clone());
                self.parent.clone().borrow_mut().insert(position, root.clone().unwrap());
                root
            }
        } else {
            None
        }
    }
    pub fn insert(&mut self, position: T, parent: T) {
        self.parent.borrow_mut().insert(position, parent);
    }
    pub fn set_count(&self) -> BTreeSet<T> {
        let mut result = BTreeSet::new();
        let parent = self.parent.borrow();
        let keys: Vec<_> = parent.keys().cloned().collect();
        drop(parent);
        for item in keys {
            result.insert(self.find_root(item.clone()).unwrap());
        }
        result
    }
    pub fn union(&mut self, set1: T, set2: T) {
        let root1 = self.find_root(set1).unwrap();
        let root2 = self.find_root(set2).unwrap();
        self.insert(root1, root2);
    }
}

struct Solution;

impl Solution {
    fn has_earth(grid: &Vec<Vec<char>>, row_index: i64, col_index: i64) -> bool {
        if row_index < 0 || col_index < 0 {
            false
        } else {
            grid[row_index as usize][col_index as usize] == '1'
        }
    }

    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let mut already_searched = DisjointSetUnion::new();
        for (row_index, row) in grid.iter().enumerate() {
            for (col_index, col) in row.iter().enumerate() {
                if col == &'1' {
                    let top = Self::has_earth(&grid, (row_index as i64) - 1, col_index as i64);
                    let left = Self::has_earth(&grid, row_index as i64, (col_index as i64) - 1);
                    if top && left {
                        already_searched.insert(Position(row_index as i64, col_index as i64), Position((row_index as i64) - 1, col_index as i64));
                        already_searched.union(Position(row_index as i64, col_index as i64 - 1), Position((row_index as i64) - 1, col_index as i64));
                    } else if top {
                        already_searched.insert(Position(row_index as i64, col_index as i64), Position((row_index as i64) - 1, col_index as i64))
                    } else if left {
                        already_searched.insert(Position(row_index as i64, col_index as i64), Position(row_index as i64, col_index as i64 - 1))
                    } else {
                        already_searched.insert(Position(row_index as i64, col_index as i64), Position(row_index as i64, col_index as i64))
                    }
                }
            }
        }
        let roots = already_searched.set_count();
        roots.len() as _
    }
}

fn main() {
    println!("{}",
             Solution::num_islands(vec![
                 vec!['1', '0', '1', '1', '1'],
                 vec!['1', '0', '1', '0', '1'],
                 vec!['1', '1', '1', '0', '1'],
             ]));
}
