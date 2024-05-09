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
            _ => {
                if val < self.min.unwrap() {
                    self.min = Some(val);
                }
            }
        }
    }

    fn pop(&mut self) {
        let last = self.stack.pop();
        if last == self.min {
            // Jank
            if self.stack.is_empty() {
                self.min = None
            } else {
                self.min = Some(*self.stack.iter().min().unwrap());
            }
        }
    }

    fn top(&self) -> i32 {
        self.stack.last().unwrap().clone()
    }

    fn get_min(&self) -> i32 {
        self.min.unwrap()
    }
}
fn main() {
    // Your MinStack object will be instantiated and called as such:
    let mut obj = MinStack::new();
    let val = 69;
    obj.push(val);
    obj.pop();
    let ret_3: i32 = obj.top();
    let ret_4: i32 = obj.get_min();
    dbg!(ret_3, ret_4);
}
