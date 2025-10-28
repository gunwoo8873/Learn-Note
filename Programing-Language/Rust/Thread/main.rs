use std::thread;

fn main() {
    let get_handler = thread::spawn(get_thread); // OR [Long Type] thread::Builder::new().spawn().unwrap();
    get_handler.join().expect("Thread panicked");

    scope_thread();
}

fn get_thread() {
    let current = thread::current();

    //// Get thread ID and Name
    let get_id = current.id(); 
    println!("Current thread ID: {:?}", get_id);

    let get_name = current.name().unwrap_or("Unknown");
    println!("Current thread name: {:?}", get_name);
}

fn scope_thread() {
    let v = vec![1, 2, 3, 4, 5];

    thread::scope(|sc| {
        sc.spawn(|| {
            println!("length : {}", v.len());
        });
        sc.spawn(|| {
            for i in &v {
                println!("{}", i);
            }
        });
    });


    //// *scope 함수에 클로저를 전달하여 해당 클로저는 즉시 실행되며 범위를 나타내는 인수는 'sc'를 입력으로 받는다. 
    //// *[단 범위가 끝날 때 까지 종료되지 않은 스레드를 기다린다.]
    thread::scope(|sc| {
        //// spawn 함수에 'static 타입이 아닌 인수를 입력 받을 수 있게된다.
        sc.spawn(|| {
            println!("Scoped thread running");
        });
        sc.spawn(|| {
            println!("Second scoped thread running");
        });
    });
}