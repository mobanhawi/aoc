use crate::util::read::file;
pub fn run(size: u32, filename: String) {
    let input = file(filename).
        expect("Failed to read input file");
    let mut result: u128 = 0;
    for line in input.lines() {
        let mut bank: Vec<u32> = line.chars()
            .filter_map(|c| c.to_digit(10))
            .collect();
        for _ in 0..(line.len() - size as usize) {
            let mut index = 0;
            let mut min_index = 0;


            for b in &bank {
                if index == bank.len() - 1 {
                    min_index = index;
                    break;
                }
                if *b < bank[index + 1] {
                    min_index = index;
                    break;
                }
                index += 1;
            }
            bank.remove(min_index);
        }
        let base: u128 = 10;
        for i in 0..bank.len()  {
            result += bank[i] as u128 *base.pow(size-1-i as u32);
        }
    }
    println!("Result: {}", result);
}
