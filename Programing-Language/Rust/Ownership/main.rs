use std::thread;
use std::rc::Rc;

fn main() {
    static_ownership();
    // memory_address_example();
}

fn static_ownership() {
    //// Standard static variable
    // static으로 선언된 변수는 일정한 initializer[이니셜라이저 : 초기화]를 갖고 드랍하지 않으며 프로그램이 시작전 생성된다.
    // 어떤 스레드도 static 변수로 부터 값을 빌려올 수 있게된다. 
    // [단, 하나의 static에는 모두 접근할 수 있지만 모두가 스레드를 소유할 수 없다.]
    static X: [i32; 5] = [1, 2, 3, 4, 5];

    //// Leaking[누수] static variable
    // leak는 Box의 소유권을 해제하고 값이 드랍되지 않게 한다. 
    // 하지만 프로그램이 종료될 때까지 존재하게 되어, 다른 스레드에서 값을 빌려갈 수 있게된다.
    let x: &'static [i32; 5] = Box::leak(Box::new([1, 2, 3, 4, 5]));

    thread::scope(|sc| {
        sc.spawn(move || {
            dbg!(X);
        });
        sc.spawn(move || {
            dbg!(x);
        });
        sc.spawn(move || {
            //// Reference Counting[레퍼런스 카운팅]
            // Reference Counting은 Rc의 타입을 사용해 구현이 가능하다. 데이터를 복제하면 새로운 데이터가 메모리에
            // 할당되는 것이 아닌 Reference Count 값이 증가하게 된다.
            // [Rc는 Thread safe[스레드 안전성]이 보장되지 않는 타입이기 때문에, 여러 개의 스레드가 특정 값에 대해 사용하면 Reference Count가 변경 될 수 있다.]
            let a = Rc::new([1, 2, 3, 4, 5]);
            let a_clone = a.clone();
    
            println!("Original pointer address: {:p}", a.as_ptr());
            println!("Cloned pointer address: {:p}", a_clone.as_ptr());
        });
    });
}