enum IpAdd {
  V4(u8, u8, u8, u8),
  V6(String),
}

fn main() {
  let home = IpAdd::V4(127, 0, 0, 1);
  let loopback = IpAdd::V6(String::from("::1"));

  match home {
    IpAdd::V4(a, b, c, d) => {
      println!("Home is a V4 address: {}.{}.{}.{}", a, b, c, d);
    }
    IpAdd::V6(s) => {
      println!("Home is a V6 address: {}", s);
    }
  }

  match loopback {
    IpAdd::V4(a, b, c, d) => {
      println!("Loopback is a V4 address: {}.{}.{}.{}", a, b, c, d);
    }
    IpAdd::V6(s) => {
      println!("Loopback is a V6 address: {}", s);
    }
  }
}