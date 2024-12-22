use std::{thread::spawn, sync::Arc};

fn main() {
  main_ownership();
  static_ref();
  leaking();
  ref_counting();
  variable_ref();
}

fn variable_ref() {
  let a = 10;
  let b = sub_variable_ref(a);

  println!("{} + 10 = {}", a, b);
}

fn sub_variable_ref(a: i32) -> i32 {
  a + 10
}

fn main_ownership() {
  let a: i32 = 15;
  let b = 20;

  sub_ownership_ref(&a, &b);
}

fn sub_ownership_ref(a: &i32, b: &i32) {
  println!("{} + {} = {}", a, b, a + b);
}

fn static_ref() {
  static X: [i32; 5] = [1, 2, 3, 4, 5];
  spawn(|| dbg!(&X));
}

fn leaking() {
  let x: &'static [i32; 5] = Box::leak(Box::new([1, 2, 3, 4, 5]));
  spawn(move || dbg!(x));
}

fn ref_counting() {
  let a = Arc::new([1, 2, 3, 4, 5]);
  let b = a.clone();   // A Ownership Variable not move to B and Pointer move

  assert_eq!(a.as_ptr(), b.as_ptr()); // assert_eq! is Equal Check

  spawn(move || dbg!(a));
  spawn(move || dbg!(b));
}