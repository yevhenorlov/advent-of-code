pub mod solutions;

use crate::solutions::s01::s0101;
use crate::solutions::s01::s0102;
use std::fs;

fn main() {
    let input = fs::read_to_string("input/01.txt").expect("Can't read the file, check the path");
    println!("{:?}", s0101(&input));
    println!("{:?}", s0102(&input));
}
