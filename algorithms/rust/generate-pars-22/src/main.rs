fn generate(s: String, n: i32) -> Vec<String> {
    let mut out: Vec<String> = vec![s.chars().collect()];
    let num_open = s.chars().filter(|c| *c == '(').count() as i32;
    let mut num_clos = s.chars().filter(|c| *c == ')').count() as i32;
    if num_open == n {
        // then we can only add ) to close
        while num_clos < n {
            out.push(')'.to_string());
            num_clos += 1;
        }
        // dbg!(&out.join(""), &num_clos, &num_open);
        return vec![out.join("")];
    }
    if num_open == num_clos && num_open == n {
        // this is a valid one so we can only add (
        return vec![out.join("")];
    };
    if num_open < n {
        let mut sl = s.clone();
        sl.push('(');
        // dbg!(&sl);
        out.append(generate(sl, n).as_mut());
    }
    if num_clos < num_open {
        let mut sr = s.clone();
        sr.push(')');
        dbg!(&out);
        out.append(generate(sr, n).as_mut());
        dbg!(&out);
    }
    out
}
pub fn generate_parenthesis(n: i32) -> Vec<String> {
    let out = generate("".to_string(), n);
    out
}
fn main() {
    let res = generate_parenthesis(2);
    dbg!(res);
}
