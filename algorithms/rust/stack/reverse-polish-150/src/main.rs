pub fn eval_rpn(tokens: Vec<String>) -> i32 {
    let mut stack: Vec<i32> = Vec::new();
    let ops: String = String::from("+-/*");
    for tok in tokens {
        if ops.contains(&tok) {
            let rhs = stack.pop().unwrap();
            let lhs = stack.pop().unwrap();
            match tok.as_str() {
                "+" => stack.push(lhs + rhs),
                "-" => stack.push(lhs - rhs),
                "/" => stack.push(lhs / rhs),
                "*" => stack.push(lhs * rhs),
                _ => panic!("Operation {} not available", tok),
            }
        } else {
            stack.push(tok.parse::<i32>().unwrap());
        }
    }
    stack.pop().unwrap()
}

fn main() {}
