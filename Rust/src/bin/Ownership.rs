use std::{thread, sync::Arc};

fn main()
{
  static_ref();
  leaking();
  ref_counting();
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
  let b = a.clone(); // A Ownership Variable not move to B

  assert_eq!(a.as_ptr(), b.as_ptr());

  thread::spawn(move || dbg!(a));
  thread::spawn(move || dbg!(b));
}