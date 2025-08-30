use std::thread;

fn main() {
  thread_builder();

  println!("Start Main Thread process");
}

fn info_thread_id() {
  let thread_id = thread::current().id();
  println!("Thread ID : {thread_id:?}");
}

fn thread_builder() {
  for i in 0..=10 {
    let builder_handle = thread::Builder::new().name("Thread Num {i}".to_string())
    .spawn(move || {
      info_thread_id();
    })
    .unwrap();
  }
}