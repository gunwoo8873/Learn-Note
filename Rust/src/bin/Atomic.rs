use std::{io, thread, sync::atomic::{AtomicBool, Ordering::Relaxed}};

fn main()
{
  atomic();
}

fn atomic()
{
  static STOP: AtomicBool = AtomicBool::new(false); // Stop Flag to Tread stop function

  let back_thread = thread::spawn(|| {
    while !STOP.load(Relaxed)
    {
      some_work();
    }
  });

  for line in io::stdin().lines()
  {
    match line.unwrap().as_str()
    {
      "help" => println!("Commands: help, stop"),
      "stop" => break,
      cmd => println!("Unknown command: {:?}", cmd),
    }
  }

  STOP.store(true, Relaxed); // Background Thread Stop Flag
  back_thread.join().unwrap();          // Background Thread Wait for Stopping
}