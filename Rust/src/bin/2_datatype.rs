pub fn main() {
    integer_type();
    float_type();
    char_type();
    string_type();
    boolean_type();
}
// Note : 모든 값은 특정한 데이터 타입을 가지고 있으며, 해당 데이터로 작업하는 방법을 알릴 수 있도록 어떤 종류의 데이터가
//        지정되어 있는지 알려준다. 스칼라 타입과 복합 타입으로 두 가지의 부분 집합으로 나누어져 있다.


fn integer_type()
{
    let signed: i32 = -1;
    println!("Signed : {}", signed);

    let unsigned: u32 = 0;
    println!("Unsigned : {}", unsigned);
}
// Note : 정수형 타입 [Integer Typed]
//        8 ~ 128 bit
//        Signed   : [i8, i16, i32, i64, i128] = -2^n-1 ~ 2^n-1
//        Unsigned : [u8, u16, u32, u64, u128] = -2^7-1 ~ 2^7-1
//        isize    : arch
//        unsize   : arch

fn float_type()
{
    let flat_type: f64 = 20.000;
    println!("Float Datatype : {}", flat_type);

    let flat_type: f32 = 20.000;
    println!("Float Datatype : {}", flat_type);
}
// Note : 부동소수점 타입 [Float Typed]
//        기본적으로 Float은 2배수 정밀도로 설정되어 있으니 별도로 f32를 명시필요 하다.
//        부동소수점은 IEEE-754 표준을 따라 f32는 1배수 정밀도이지만, f64는 2배수 정밀도를 가지고 있다.
// Link : https://ko.wikipedia.org/wiki/IEEE_754 [출처 = 위키백과 : IEEE-754]

fn char_type()
{
    let char: char = 'c'; // Text size memory is 1 byte size and Korean String is 2~4 byte size
    println!("Char Datatype : {}", char);
}
// Note : Char라는 데이터 타입은 기본적으로 UTF-8을 사용하여 1 Byte 라는 크기를 가지고 있지만 한글을 사용하게 되면
//        2에서 4까지의 크기를 가지고 있는다. 즉 유니코드의 스칼라 값을 표현
//        ASCII 보다 많은 값을 표현할 수 있어 다양한 언어, 이모지, 0인 공백 문자 모두가 유효한 타입이다.
//        유니코드 스칼라값의 범위는 U+0000 ~ U+D7FF, U+E000, U10FFFF 값이다.

fn string_type()
{
    let string: String = "String".to_string();
    println!("String Datatype : {}", string);
}

fn boolean_type()
{
    let boolean: bool = true;
    println!("Boolean Datatype : {}", boolean);
}