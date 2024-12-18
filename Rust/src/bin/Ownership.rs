use std::{thread, sync::Arc};

fn main()
{
  main_ownership();
  static_ref();
  leaking();
  ref_counting();
}

fn main_ownership()
{
  let a: i32 = 15;
  let b = 20;

  sub_ownership_ref(&a, &b);
}

fn sub_ownership_ref(a: &i32, b: &i32)
{
  println!("{} + {} = {}", a, b, a + b);
}

fn static_ref()
{
  static X: [i32; 5] = [1, 2, 3, 4, 5];
  thread::spawn(|| dbg!(&X));
}

fn leaking()
{
  let x: &'static [i32; 5] = Box::leak(Box::new([1, 2, 3, 4, 5]));
  thread::spawn(move || dbg!(x));
}

fn ref_counting()
{
  let a = Arc::new([1, 2, 3, 4, 5]);
  let b = a.clone();   // A Ownership Variable not move to B and Pointer move

  assert_eq!(a.as_ptr(), b.as_ptr()); // assert_eq! is Equal Check

  thread::spawn(move || dbg!(a));
  thread::spawn(move || dbg!(b));
}