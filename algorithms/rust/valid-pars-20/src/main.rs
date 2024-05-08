fn is_opening(c: &char) -> bool {
    ['[', '{', '('].contains(c)
}
fn get_closing(c: &char) -> char {
    match c {
        '{' => return '}',
        '(' => return ')',
        '[' => return ']',
        _ => panic!("Please input opening!"),
    }
}
pub fn is_valid(s: String) -> bool {
    let mut stack: Vec<char> = vec![];
    for c in s.chars() {
        if is_opening(&c) {
            stack.push(c);
        } else {
            let Some(last) = stack.pop() else {
                return false;
            };
            if c != get_closing(&last) {
                return false;
            }
        }
    }
    stack.is_empty()
}
fn main() {
    dbg!(is_valid(String::from("()[]{}")));
    dbg!(is_valid(String::from("({})")));
    dbg!(is_valid(String::from("({[}])")));
}
