fn main() {
    add_compile();
    borrowing_compile();
    pointer_compile();
    vector_compile();
    return_compile();
}

fn add_compile() {
  let stack_a = String::from("Value A");

  let stack_b = stack_a.clone();

  let stack_c = add(stack_b);
  println!("stack_c = {}", stack_c);
}

fn add(mut v: String) -> String {
  v.push_str(" , Value B");
  return v
}

fn borrowing(v1: String, v2: String) {
  println!("borrowing = {} {}", v1, v2) // v1 and v2 is bind value
}

fn borrowing_compile() {
  let v1 = String::from("V1");
  let v2 = String::from("V2");

  borrowing(v1, v2);
}

fn pointer_compile() {
  pointer();
}

fn pointer() {
  let mut x: Box<i32> = Box::new(5);

  let a: i32 = *x; // Heap a = x = 5
  *x += 1;
  println!("pointer a = {}", a);
  let return_a: &Box<i32> = &x; // Heap return_a = x = 5

  let b: i32 = **return_a; // Stack result = 6 (return x + 1)
  println!("pointer b = {}", b);
  let return_b: &i32 = &*x; // Heap return_b = x = 6

  let c: i32 = *return_b; // Stack result = 6 (return x + 1)
  println!("pointer c = {}", c);
}

fn vector_compile() {
  vector_push()
}

fn vector_push() {
  let mut v: Vec<i32> = vec![1, 2, 3];
  v.push(4);

  println!("vector = {:?}", v);
}

use std::rc::Rc;
fn return_compile() {
  return_str();
}

fn return_str() -> Rc<String> {
  let s = Rc::new(String::from(""));
  Rc::clone(&s)
}