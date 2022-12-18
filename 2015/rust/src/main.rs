pub mod solutions;

use crate::solutions::s01::s0101;
use crate::solutions::s01::s0102;

const INPUT_1: &str = "input/01.txt";

fn main() {
    println!("{:?}", s0101(INPUT_1));
    println!("{:?}", s0102(INPUT_1));
}
