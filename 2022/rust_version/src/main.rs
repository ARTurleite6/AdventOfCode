use std::{fs::File, io::Read};

fn main() {
    let mut file = File::open(std::env::args().nth(1).unwrap()).unwrap();
    let mut buffer = String::new();

    file.read_to_string(&mut buffer).unwrap();

    let mut max = 0;
    let mut sum = 0;

    for line in buffer.lines() {
        if line.is_empty() {
            if sum > max {
                max = sum;
            }
            sum = 0;
        } else {
            sum += line.parse::<i32>().unwrap();
        }
    }

    println!("{max}");
}
