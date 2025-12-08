use std::fs::File;
use std::io::{self, Read};
pub fn file(file_path: String) -> io::Result<String> {
    let mut file = File::open(file_path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}