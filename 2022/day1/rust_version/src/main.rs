use std::{fs::File, io::Read, collections::BinaryHeap};

fn main() {
    let mut file = File::open(std::env::args().nth(1).unwrap()).unwrap();
    let mut buffer = String::new();

    file.read_to_string(&mut buffer).unwrap();
    let mut heap = BinaryHeap::new();

    let mut sum = 0;

    for line in buffer.lines() {
        if line.is_empty() {
            heap.push(sum);
            sum = 0;
        } else {
            sum += line.parse::<i32>().unwrap();
        }
    }

    let mut iter = heap.into_iter();
    let result = iter.next().unwrap_or(0) + iter.next().unwrap_or(0) + iter.next().unwrap_or(0);

    println!("{result}");
}
