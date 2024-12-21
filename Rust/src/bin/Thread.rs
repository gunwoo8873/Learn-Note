use std::{
  thread::{spawn, park, current, scope, sleep},
  sync::{Arc, Condvar},
  rc::Rc,
};

fn main()
{
  println!("Main Thread"); // ThreadID N.1

  spawn(output); // ThreadID N.2
  spawn(output); // ThreadID N.3

  join();
  closure();
  thread_scope();
  parking();
  condition();
  counting();
  arc();
}
fn output()
{
  println!("Another Thread");

  let id = current().id(); // This lib to Current ThreadID Checking Variable

  println!("Thread id : {:?} ", id); // Thread ID Output
}

fn join()
{
  let t1 = spawn(output); // ThreadID Add N.4
  t1.join().unwrap();

  let t2 = spawn(output); // ThreadID Add N.5
  t2.join().unwrap();
}

fn closure()
{
  let num_vec_a = vec![1, 2, 3];

  // New Create Thread and Keyword Move use the Closure
  spawn(move || {
    for t in &num_vec_a
    {
      println!("{}", t);
    }
  })
  .join()
  .unwrap();

  let num_vec_b = Vec::from_iter(0..=1000);
  
    let t = spawn(move || {
        let len = num_vec_b.len();
        let sum = num_vec_b.into_iter().sum::<usize>();
        
        sum / len // Return
    });

    let average = t.join().unwrap();
    println!("{}", average);
}

fn thread_scope()
{
  let num_vec_a = vec![1, 2, 3];

  scope(|s| {
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
  scope(|s| {
    let t = s.spawn(|| loop {
      let item = queue.lock().unwrap().pop_front();
      
      if let Some(item) = item
      {
        dbg!(item);
      }
      else
      {
        park(); // Fake Function return to Thread
      }
    });
    
    for i in 0..
    {
      queue.lock().unwrap().push_back(i);
      t.thread().unpark();
      sleep(Duration::from_millis(100));
    }
  });
}

fn condition()
{
  let queue = Mutex::new(VecDeque::new());
  let not_empty = Condvar::new();

  scope(|s| {
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
      sleep(Duration::from_millis(100));
    }
  });
}

fn counting()
{
    let a = Rc::new([1, 2, 3]);
    let b = a.clone();       // a = [1,2,3]을 복사

    assert_eq!(a.as_ptr(), b.as_ptr()); // a와 b의 raw pointer를 같은지 비교하고, 같은 메모리를 가리킨다.
}
// Note : as_ptr는 if의 == 연산자와 비슷한 성격을 가지고 있다. 같은 위치를 가리키고 있는지 확인하여,
//        메모리 주소를 가리키는 raw pointer를 얻는다. [Waring : 메모리의 안전성을 보장받지 못하는 코드다. = unsafe]
//        assert_eq!는 두 개의 값이 같은지 확인하여, 수행을 이어서 하지만 다르다면 즉시 실행을 중단하고
//        Panic이 발생한다. [즉 오류가 발생하여 오류에 대한 MSG를 출력]

fn arc()
{
    let a = Arc::new([1, 2, 3]);
    let b = a.clone();

    spawn(move || dbg!{a});
    spawn(move || dbg!(b));
}
// Note : Send : 타입의 값과 소요권을 다른 Thread에 이전을 하고 싶다면 Send Trait를 구현하면 된다.
//        Sync : 타입의 값을 여러 Thread에 공유하고 싶다라면 Sync Trait로 구현하면 된다.
//        bool과 str 같은 원시 타입들은 Send와 Sync를 모두 구현하고 있기 때문에 필요한 구조체 필드의
//        타입에 따라서 구조체에 해당 트레이트가 자동으로 구현되는 것을 Auto Trait라고 칭한다.
//        T 타입으로 인식을 하지만 실제로 해당 타입은 존재하지 않아 크기가 0인 메모리를 차지하지 않는다.