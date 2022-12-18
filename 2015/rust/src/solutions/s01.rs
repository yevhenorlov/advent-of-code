use std::fs;

pub fn s0101(path: &str) {
    let input = fs::read_to_string(path).expect("Can't read the file, check the path");
    let mut floor = 0;
    for c in input.trim().chars() {
        match c {
            '(' => floor+=1,
            ')' => floor-=1,
            _ => println!("unhandled case"),
        }
    }
    println!("{floor}")
}
