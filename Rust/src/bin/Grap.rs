use std::env;

fn main() {
  let atgs: Vec<String> = env::args().collect(); // Don't Not Adjective unicode do panic!

  let query = &atgs[1];
  let file_path = &atgs[2];

  println!("Searching for {}", query);
  println!("In file {}", file_path);
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_main() {
    main();
  }
}