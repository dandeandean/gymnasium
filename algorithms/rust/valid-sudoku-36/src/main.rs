
fn get_3x3(board: &Vec<Vec<char>>, offset: usize) -> Vec<char>{
    let offset_x = (offset/3)*3;
    let offset_y = (offset%3)*3;
    let mut out : Vec<char>= Vec::new() ;
    for i in 0..3{
        for j in 0..3{
            out.push(board[i+offset_x][j+offset_y]);
        }
    }
    return out 
}
fn get_col(board: &Vec<Vec<char>>, offset: usize) -> Vec<char>{
    let mut out: Vec<char> = Vec::new();
    for i in 0..9{
        out.push(board[i][offset]);
    }
    out
}
fn get_row(board: &Vec<Vec<char>>, offset: usize) -> Vec<char>{
    return board[offset].clone();
}
fn has_repeat(area: &Vec<char>) -> bool {
    let mut already_used: Vec<u16> = vec![];
    for c in area {
        let d = c.to_digit(10);
        match d {
            Some (n) => {
                if already_used.contains(&(n as u16)) {
                    return true;
                }
                already_used.push(n as u16);
            }
            None => {}
        }
    }
    false
}
pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {        
    for i in 0..9{
        if has_repeat(&get_3x3(&board, i)) {
            return false
        }
        if has_repeat(&get_col(&board, i)) {
            return false
        }
        if has_repeat(&get_row(&board, i)) {
            return false
        }
    }
    true
}
fn main() {
    let board =
     vec![
        vec!['5','3','.','.','7','.','.','.','.'],
        vec!['6','.','.','1','9','5','.','.','.'],
        vec!['.','9','8','.','.','.','.','6','.'],
        vec!['8','.','.','.','6','.','.','.','3'],
        vec!['4','.','.','8','.','3','.','.','1'],
        vec!['7','.','.','.','2','.','.','.','6'],
        vec!['.','6','.','.','.','.','2','8','.'],
        vec!['.','.','.','4','1','9','.','.','5'],
        vec!['.','.','.','.','8','.','.','7','9']
    ];
    dbg!(is_valid_sudoku(board));
}
