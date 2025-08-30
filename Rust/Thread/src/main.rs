use std::thread;

fn main() {
  //// 4. Thread fn 호출 후 종료
  thread::spawn(fn_a);

  //// 1. main fn 안에 존재하는 값 부터 실행
  println!("Main Thread Task run process");
}

fn fn_a() {
  //// 2. Thread ID 1번 생성
  println!("Fn A Thread Task run process");

  //// 3. Thread ID 2번 생성
  let thread_id = thread::current().id(); // Thread ID 식별자 함수
  println!("Current output Thread id : {thread_id:?}"); // 총 생성된 Thread ID 값을 출력
}
