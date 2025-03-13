use std::{
  i8,
  i16,
  i32,
  i64,
  i128,
  u8,
  u16,
  u32,
  u64,
  u128,
};

fn main() {
  integer_isize();
  integer_usize();
  integer_compile();
  character_compile();
  tuple_compile();
}

fn integer_isize() {
  let i8_bit = i8::BITS;
  let i8_min = i8::MIN;
  let i8_max = i8::MAX;
  println!("Int {}\t Min = {}\t Max = {}", i8_bit, i8_min, i8_max);

  let i16_bit = i16::BITS;
  let i16_min = i16::MIN;
  let i16_max = i16::MAX;
  println!("Int {}\t Min = {}\t Max = {}", i16_bit, i16_min, i16_max);

  let i32_bit = i32::BITS;
  let i32_min = i32::MIN;
  let i32_max = i32::MAX;
  println!("Int {}\t Min = {}\t Max = {}", i32_bit, i32_min, i32_max);

  let i64_bit = i64::BITS;
  let i64_min = i64::MIN;
  let i64_max = i64::MAX;
  println!("Int {}\t Min = {}\t Max = {}", i64_bit, i64_min, i64_max);

  let i128_bit = i128::BITS;
  let i128_min = i128::MIN;
  let i128_max = i128::MAX;
  println!("Int {}\t Min = {}\t Max = {}", i128_bit, i128_min, i128_max);
}

fn integer_usize() {
  let u8_bit = u8::BITS;
  let u8_min = u8::MIN;
  let u8_max = u8::MAX;
  println!("Int {}\t Min = {}\t Max = {}", u8_bit, u8_min, u8_max);

  let u16_bit = u16::BITS;
  let u16_min = u16::MIN;
  let u16_max = u16::MAX;
  println!("Int {}\t Min = {}\t Max = {}", u16_bit, u16_min, u16_max);

  let u32_bit = u32::BITS;
  let u32_min = u32::MIN;
  let u32_max = u32::MAX;
  println!("Int {}\t Min = {}\t Max = {}", u32_bit, u32_min, u32_max);

  let u64_bit = u64::BITS;
  let u64_min = u64::MIN;
  let u64_max = u64::MAX;
  println!("Int {}\t Min = {}\t Max = {}", u64_bit, u64_min, u64_max);

  let u128_bit = u128::BITS;
  let u128_min = u128::MIN;
  let u128_max = u128::MAX;
  println!("Int {}\t Min = {}\t Max = {}", u128_bit, u128_min, u128_max);
}

struct Operator {
    value_a: i32,
    value_b: i32,
}

impl Operator {
    fn sum(&self) -> i32 { self.value_a + self.value_b }
    fn subtraction(&self) -> i32 { self.value_a - self.value_b }
    fn multiplication(&self) -> i32 { self.value_a * self.value_b }
    fn division(&self) -> i32 { self.value_a / self.value_b }
    fn remainder(&self) -> i32 { self.value_a % self.value_b }
}

fn integer_compile() {
    let sum = Operator { value_a: 10, value_b: 20 }.sum();
    println!("sum = {}", sum);

    let subtraction = Operator { value_a: 10, value_b: 20 }.subtraction();
    println!("subtraction = {}", subtraction);

    let multiplication = Operator { value_a: 10, value_b: 20 }.multiplication();
    println!("multiplication = {}", multiplication);

    let division = Operator { value_a: 40, value_b: 20 }.division();
    println!("division = {}", division);

    let remainder = Operator { value_a: 45, value_b: 20 }.remainder();
    println!("remainder = {}", remainder);
}

struct Character {
    char: char,
}

impl Character {
    fn char(&self) -> char { self.char }
}

fn character_compile() {
    let char = Character { char: 'a' }.char();
    println!("char = {}", char);

    let char = Character { char: 'A' }.char();
    println!("char = {}", char);
}

fn tuple_compile() {
    let tuple_index: (i32, f64, char) = (1, 2.0, 'a');
    let tuple_get = tuple_index.0;
    println!("tuple_get = {}", tuple_get);
}