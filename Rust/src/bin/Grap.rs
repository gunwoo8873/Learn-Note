use std::{env, fs};

fn main() {
  path_grap();
  read_file();
}

struct Config {
  query: String,
  file_path: String,
}

// impl Config {
//   fn new(args: &[String]) -> Config {
//     let query = args[1].clone();
//     let file_path = args[2].clone();

//     Config {query, file_path}
//   }
// }

fn parse_config(args: &[String]) -> Config {
  let query = args[1].clone();
  let file_path = args[2].clone();

  Config {query, file_path}
}

fn path_grap() {
  let args: Vec<String> = env::args().collect(); // Not Adjective unicode is panic!
  let config = parse_config(&args);

  println!("Searching for {}", config.query);
  println!("In file {}", config.file_path);
}

fn read_file() {
  let args: Vec<String> = env::args().collect();
  let config = parse_config(&args);

  let contents = fs::read_to_string(config.file_path).expect("msg");

  println!("With text:\n{}", contents);
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_grap() {
    path_grap();
  }

  #[test]
  fn test_read_file() {
    read_file();
  }
}

// CMD : cargo run --bin grap -- test sample.txt