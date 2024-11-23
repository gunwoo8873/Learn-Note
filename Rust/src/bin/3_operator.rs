pub fn main() {
    elementary_arithmetic();
    boolean_type();
    tuple();
    array();
}

const INT_A: i32 = 10;
const INT_B: i32 = 15;

fn elementary_arithmetic() {
    let addition_result = INT_A + INT_B;
    println!("Addition : {}", addition_result);

    let subtraction_result = INT_A * INT_B;
    println!("Subtraction : {}", subtraction_result);

    let multiplication_result = INT_A * INT_B;
    println!("Multiplication : {}", multiplication_result);

    let division_result = INT_A / INT_B;
    println!("Division : {}", division_result);

    let modulo_result = INT_A % INT_B;
    println!("Modulo : {}",modulo_result);
}

const TRUE: bool = true;
const FALSE: bool = false;

fn boolean_type() {
    let t = TRUE;
    println!("true : {}", t);

    let f = FALSE;
    println!("false : {}", f);
}

const TUPLE: (i32, f64, &str, char) = (100, 3.14, "Tuple", 'a');

fn tuple() {
    let a = TUPLE.0;
    println!("Tuple_index = a : {}", a);

    let b = TUPLE.1;
    println!("Tuple_index = b : {}", b);

    let c = TUPLE.2;
    println!("Tuple_index = c : {}", c);

    let d = TUPLE.3;
    println!("Tuple_index = d : {}", d);
}
// Note : tuple 는 각 다른 데이터 타입을 배열 할 수 있는 방법 중 하나이다.
//        let value: (datatype) = (index)
//        let name = value.index

const ARRAY: [i32; 10] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

fn array() {
    println!("Array : {:?}", ARRAY);

    let array_index_5 = ARRAY[5];
    println!("Array_index_5 : {}", array_index_5);
}
// Note : Array 는 하나의 데이터 타입인 값을 배열로 저장하는 방식이다.
//        let value = [index]
//        let name = value[index]
//
//        let value: [type; index] = [index]
//        let name = value[index]