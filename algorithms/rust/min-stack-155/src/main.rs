struct MinStack {
    // use push_back and pop
    pub stack: Vec<i32>,
    // Store a copy b/c idk if we can use lifetime for leetcode
    pub min: Option<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {
    fn new() -> Self {
        MinStack {
            stack: vec![],
            min: None,
        }
    }

    fn push(&mut self, val: i32) {
        self.stack.push(val);
        match self.min {
            None => self.min = Some(val),
            Some(current_min) => {
                if val < current_min {
                    self.min = Some(val);
                }
            }
        }
    }

    fn pop(&mut self) {
        if self.stack.pop() == self.min {
            if self.stack.is_empty() {
                self.min = None
            } else {
                self.min = Some(*self.stack.iter().min().unwrap());
            }
        }
    }

    fn top(&self) -> i32 {
        match self.stack.last() {
            None => panic!("You called top on an empty stack! That is wrong and immoral"),
            Some(val) => *val,
        }
    }

    fn get_min(&self) -> i32 {
        match self.min {
            None => panic!("You called get min on an empty stack!"),
            Some(val) => val,
        }
    }
}
fn main() {
    // Your MinStack object will be instantiated and called as such:
    let mut obj = MinStack::new();
    obj.push(69);
    obj.push(420);
    obj.pop();
    let ret_3: i32 = obj.top();
    let ret_4: i32 = obj.get_min();
    dbg!(ret_3, ret_4);
}
