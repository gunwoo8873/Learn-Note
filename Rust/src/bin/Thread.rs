use std::thread;

fn main()
{
    thread::spawn(f);
    thread::spawn(f);

    println!("Main Thread");

    join();
}

fn f()
{
    println!("Another Thread");

    let id = thread::current().id();
    println!("Thread id : {:?} ", id);
}

const NUM: Vec<i32> = vec![1, 2, 3];

fn join()
{
    thread::spawn(move ||
        {
            for n in NUM
            {
                println!("Thread number : {:?} ", n);
            }
        }
    )
        .join()
        .unwrap();
}