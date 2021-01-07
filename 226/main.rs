use std::rc::Rc;
use std::cell::RefCell;
use std::mem;

impl Solution {
    pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        root.map(|root| {
            let mut root_ref = root.borrow_mut();
            let old_left = root_ref.left.take();
            let old_right = root_ref.right.take();
            root_ref.left = Self::invert_tree(old_right);
            root_ref.right = Self::invert_tree(old_left);
            drop(root_ref);
            root
        })
    }
}