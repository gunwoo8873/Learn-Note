use std::thread;

fn main()
{
  println!("Main Thread"); // ThreadID N.1

  thread::spawn(output); // ThreadID N.2
  thread::spawn(output); // ThreadID N.3

  join();
  closure();
  scope();
  parking();
  condition()
}
fn output()
{
  println!("Another Thread");

  let id = thread::current().id(); // This lib to Current ThreadID Checking Variable

  println!("Thread id : {:?} ", id); // Thread ID Output
}

fn join()
{
  let t1 = thread::spawn(output); // ThreadID Add N.4
  t1.join().unwrap();

  let t2 = thread::spawn(output); // ThreadID Add N.5
  t2.join().unwrap();
}

fn closure()
{
  let num_vec_a = vec![1, 2, 3];

  // New Create Thread and Keyword Move use the Closure
  thread::spawn(move || {
    for t in &num_vec_a
    {
      println!("{}", t);
    }
  })
  .join()
  .unwrap();

  let num_vec_b = Vec::from_iter(0..=1000);
  
    let t = thread::spawn(move || {
        let len = num_vec_b.len();
        let sum = num_vec_b.into_iter().sum::<usize>();
        
        sum / len // Return
    });

    let average = t.join().unwrap();
    println!("{}", average);
}

fn scope()
{
  let num_vec_a = vec![1, 2, 3];

  thread::scope(|s| {
    s.spawn(|| {
      println!("length : {}", num_vec_a.len());
    });

    s.spawn(|| {
      for i in &num_vec_a
      {
        println!("{}", i);
      }
    });
  });
}

use std::{sync::Mutex, time::Duration, collections::VecDeque};

fn parking()
{
  let queue = Mutex::new(VecDeque::new());
  thread::scope(|s| {
    let t = s.spawn(|| loop {
      let item = queue.lock().unwrap().pop_front();
      
      if let Some(item) = item
      {
        dbg!(item);
      }
      else
      {
        thread::park(); // Fake Function return to Thread
      }
    });
    
    for i in 0..
    {
      queue.lock().unwrap().push_back(i);
      t.thread().unpark();
      thread::sleep(Duration::from_millis(100));
    }
  });
}

use std::sync::Condvar;

fn condition()
{
  let queue = Mutex::new(VecDeque::new());
  let not_empty = Condvar::new();

  thread::scope(|s| {
    s.spawn(|| loop {
      let mut queue = queue.lock().unwrap();
      let item = loop
      {
        if let Some(item) = queue.pop_front()
        {
          break item;
        }
        else
        {
          queue = not_empty.wait(queue).unwrap();
        }
      };

      drop(queue);
      dbg!(item);
    });

    for i in 0..
    {
      queue.lock().unwrap().push_back(i);
      not_empty.notify_one();
      thread::sleep(Duration::from_millis(100));
    }
  });
}