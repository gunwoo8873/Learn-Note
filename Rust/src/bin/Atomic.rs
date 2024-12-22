use std::{
  sync::atomic::{AtomicUsize, Ordering::Relaxed}, 
  thread::{scope, sleep}, 
  time::Duration,
};

fn main() {
  current_status();
}

fn process_item(i: usize) {
  println!("Processing item {}", i);
}

fn current_status() {
  let num_done = AtomicUsize::new(0);

  scope(|s| {
    s.spawn(|| {
      for i in 0..100 {
        process_item(i);
        num_done.store(i + 1, Relaxed);
      }
    });

    loop {
      let num_done = num_done.load(Relaxed);

      if num_done == 100 {
        break;
      }

      println!("Waiting... ({}/100)", num_done);
      sleep(Duration::from_millis(100));
    }
  });

  println!("Done!");
}