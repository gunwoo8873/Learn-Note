fn main() {
  let mut count = 0;
  
  loop {
    count += 1;

    if count == 1 {
      println!("{count}st");
      continue;
    } else if count == 2 {
      println!("{count}nd");
      continue;
    } else if count == 3 {
      println!("{count}rd");
      continue;
    } else {
      println!("{count}th");
      break;
    }
  }
}
