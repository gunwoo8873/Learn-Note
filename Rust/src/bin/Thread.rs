use std::thread;

fn main()
{
  println!("Main Thread"); // ThreadID N.1

  thread::spawn(output); // ThreadID N.2
  thread::spawn(output); // ThreadID N.3

  join();
  closure();
  scope();
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