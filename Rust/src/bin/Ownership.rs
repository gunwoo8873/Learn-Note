use std::{thread::spawn, sync::Arc};

fn main() {
  nomal_ref();
  static_ref();
  leaking();
  ref_counting();
}

fn nomal_ref() {
  let x: i32 = 15;
  let y: i32 = 20;

  sub_ref(x, y);
}

fn sub_ref(x: i32, y: i32) {
  println!("{} + {} = {}", x, y, x + y);
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