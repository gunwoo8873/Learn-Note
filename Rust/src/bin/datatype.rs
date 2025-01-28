fn main() {
  integer_compile();
  character_compile();
  tuple_compile();
}

struct Integer {
    value_a: i32,
    value_b: i32,
}

impl Integer {
    fn sum(&self) -> i32 { self.value_a + self.value_b }
    fn subtraction(&self) -> i32 { self.value_a - self.value_b }
    fn multiplication(&self) -> i32 { self.value_a * self.value_b }
    fn division(&self) -> i32 { self.value_a / self.value_b }
    fn remainder(&self) -> i32 { self.value_a % self.value_b }
}

fn integer_compile() {
    let sum = Integer { value_a: 10, value_b: 20 }.sum();
    println!("sum = {}", sum);

    let subtraction = Integer { value_a: 10, value_b: 20 }.subtraction();
    println!("subtraction = {}", subtraction);

    let multiplication = Integer { value_a: 10, value_b: 20 }.multiplication();
    println!("multiplication = {}", multiplication);

    let division = Integer { value_a: 40, value_b: 20 }.division();
    println!("division = {}", division);

    let remainder = Integer { value_a: 45, value_b: 20 }.remainder();
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