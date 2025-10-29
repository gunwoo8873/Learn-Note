use std::thread;
use std::rc::Rc;
use std::sync::Arc;

fn main() {
    static_ownership();
    // memory_address_example();
}

fn static_ownership() {
    let atomic_arr = Arc::new([1, 2, 3, 4, 5]);

    thread::scope(|sc| {
        sc.spawn(move || {
            //// Standard static variable
            // static으로 선언된 변수는 일정한 initializer[이니셜라이저 : 초기화]를 갖고 드랍하지 않으며 프로그램이 시작전 생성된다.
            // 어떤 스레드도 static 변수로 부터 값을 빌려올 수 있게된다. 
            // [단, 하나의 static에는 모두 접근할 수 있지만 모두가 스레드를 소유할 수 없다.]
            static X: [i32; 5] = [1, 2, 3, 4, 5];
            dbg!(X);
        });
        sc.spawn(move || {
            //// Leaking[누수] static variable
            // leak는 Box의 소유권을 해제하고 값이 드랍되지 않게 한다. 
            // 하지만 프로그램이 종료될 때까지 존재하게 되어, 다른 스레드에서 값을 빌려갈 수 있게된다.
            let x: &'static [i32; 5] = Box::leak(Box::new([1, 2, 3, 4, 5]));
            dbg!(x);
        });
        sc.spawn(move || {
            //// Reference Counting[레퍼런스 카운팅]
            // Reference Counting은 Rc의 타입을 사용해 구현이 가능하다. 데이터를 복제하면 새로운 데이터가 메모리에
            // 할당되는 것이 아닌 Reference Count 값이 증가하게 된다.
            // [Rc는 Thread safe[스레드 안전성]이 보장되지 않는 타입이기 때문에,
            //  여러 개의 스레드가 특정 값에 대해 사용하면 Reference Count가 변경 될 수 있다.]
            let ref_counting_arr = Rc::new([1, 2, 3, 4, 5]);
            let ref_counting_clone = ref_counting_arr.clone();

            println!("Original pointer address: {:p}", ref_counting_arr.as_ptr());
            println!("Cloned pointer address: {:p}", ref_counting_clone.as_ptr());

        });
        sc.spawn(move || {
            let atomic_clone = atomic_arr.clone();
        });
    });
}


//// Immutable Borrowing[불변 대여]
// '&'을 사용해서 값을 빌리면 기본적으로 Immutable Reference[불변 레퍼런스]가 된다.
// 복사가 가능해 어떤 레퍼런스가 참조하고 있는 값은 이 레퍼런스를 복사한 모든 레퍼런스와 공유한다.

//// Mutable Borrowing[가변 대여]
// '&mut'를 사용하여 Mutable Reference[가변 레퍼런스]를 생성할 수 있다. 어떤 값을 가변으로 대여하는 것은
// 현재 가변 레퍼런스가 유일하게 이 값을 가변으로 대여하고 있다는 의미로 값을 변경 하더라도, 다른 부분에서 값이 바뀌는걸 방지할 수 있다.

fn immutable_borrowing() {
    match index() {
        0 => x(),
        1 => y(),
        _ => z(index),
    }

    let a = [123, 456, 789];
    let b = unsafe { a.get_unchecked(2) }; // Warning : Unsafe code
}