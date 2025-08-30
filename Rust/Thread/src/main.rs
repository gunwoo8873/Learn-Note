use std::thread;

fn main() {
  thread_builder();
}

fn info_thread_id() {
  //// 현재 생성된 스레드 값
  let thread_id = thread::current().id();
  println!("Thread ID : {thread_id:?}");
}

fn thread_builder() {
  let mut v = vec![1, 2, 3, 4, 5];
  
  for i in &v {
    //// 범위 스레드
    let scope_handler = thread::scope(|s| {
      s.spawn(|| {
        println!("Thread First Scope Num {i}");
        info_thread_id();
      });

      s.spawn(|| {
        println!("Thread Second Scope Num {i}");
        info_thread_id();
      });
    });

    //// 사전 구성 설정 스레드
    let builder_handler = thread::Builder::new().name("Thread Num {i}".to_string())
    .spawn(move || {
      info_thread_id(); 
    })
    .unwrap();
  }

  //// End Process delay
  thread::sleep_ms(1000);
}