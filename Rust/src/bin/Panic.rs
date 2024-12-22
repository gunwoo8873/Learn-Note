use std::fs::File;
use std::io::ErrorKind;
fn main() {
    panic();
    read_file();
}

fn panic() {
    panic!("Test Panic!");
}
// Result : error: process didn't exit successfully: `target\debug\panic.exe` (exit code: 101)
// Note : panic!은 unwinding이라는 되감기 명칭을 가지고 있지만 데이터 청소를 뜻한다. 하지만 정리 작없 없이 즉시 컴파일이
//        종료되는 방법을 러스트는 사용한다. 단 panic!은 어떠한 이유던 간에 오류를 발생시키기 때문에 주의해야 한다.

fn read_file() {
    let file_result = File::open("hello.txt");
    let _result = file_result.unwrap_or_else(|e| match e.kind() // 왜 result를 사용하다 _를 붙여야 할까?
    {
        ErrorKind::NotFound => match File::create("hello.txt")
        {
            Ok(t) => t,
            Err(e) => panic!("Test Create File : {}", e),
        },
        _error => panic!("Test Create File Error : {}", _error),
    });
}
// Note : File lib는 open부터 create, read, write를 가지고 있지만 std::io도 같이 사용해야 하는 경우가 많는거 같다.
//        또한 ErrorKind::NotFound의 개념도 중요한데 unwarp을 과연 사용하지 않아도 될까? 혹은 NotFound를 왜 사용할까? 라는 의문
// Link : https://doc.rust-lang.org/std/fs/struct.File.html