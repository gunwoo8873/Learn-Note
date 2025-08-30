use std::thread;
use std::rc;

fn main() {
  //// static은 일정한 이니셜라이저[Initializer]를 갖고 드랍되지 않아 시작전 생성
  static X: [i32; 10] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  thread::spawn(|| dbg!(&X));

  //// leak은 값의 할당을 누수[Leaking] 발생 시키는 방법이며 Box의 소유권을 해제하고 값이 드랍[Drop]되지 않고 유지
  let x: &'static [i32; 10] = Box::leak(Box::new(X));
  thread::spawn(move || dbg!(x));

  /* Warning! */
    
  //// 레퍼런스 카운팅[Reference Counting]으로 새로운 데이터가 메모리에 할당되지 않지만, 레퍼런스 카운터의 값이 증가[같은 메모리 할당 값을 참조]
  //// Rc가 삭제되면 카운터가 감소하고 마지막 값이 드랍되면 메모리에서 값이 해제되지만, 다른 스레드로 보내면 에러 발생 
  
  // let ref_counting = rc::Rc::new(X);
  // let ref_count_clone = ref_counting.clone();

  // thread::spawn(move || dbg!(ref_counting));
  // thread::spawn(move || dbg!(ref_count_clone));

  let ref_arr = sync::Arc::new(X);
  let ref_arr_clone = ref_arr.clone();

  thread::spawn(move || dbg!(ref_arr));
  thread::spawn(move || dbg!(ref_arr_clone));
}