use std::collections::HashMap;

fn main() {
    get_vector();
    get_string();
    string_replace();
    get_hash_map();
}

enum Collection {
    Vector(Vec<i32>),
    Str(String),
    HashMap(HashMap<String, i32>),
}

impl Collection {
    fn vector() -> Self
    {
        Collection::Vector(Vec::new())
    }

    fn string() -> Self
    {
        Collection::Str(String::new())
    }

    fn hashmap() -> Self
    {
        Collection::HashMap(HashMap::new())
    }

    fn display(&self)
    {
        match self
        {
            Collection::Vector(v) => println!("Vector : {:?}", v),
            Collection::Str(s) => println!("String : {}", s),
            Collection::HashMap(hm) => println!("HashMap : {:?}", hm),
        }
    }
}

fn get_vector() {
    let mut vector_a = Collection::vector();
    if let Collection::Vector(v) = &mut vector_a {
        v.push(1);
        v.push(2);
        v.push(3);
    }
    vector_a.display();

    let mut vector_b : Vec<i32> = vec![4, 5, 6, 7];
    for i in &mut vector_b
    {
        *i += 50;
        println!("{}", i);
    }
}

fn get_string() {
    let mut str_a = Collection::string();
    if let Collection::Str(s) = &mut str_a {
        s.push_str("String_A");
        s.push_str("String_B");
    }
    str_a.display();

    let str_b = "Collection String B";
    let str = str_b.to_string();
    println!("{}", str);

    let str_c = "Collection String C";
    for i in str_c.bytes()
    {
        println!("{}", i);
    }
}
// Note : to_string과 String::from의 차이는 단순히 가독성만 추구하는 것이 아닌 효율성도 같이 오기 때문에 적절한
//        상항에서 사용 하는 것을 생각해야 한다. 앞서 String::new가 존재하기 때문에 이 점을 유의
// Tip : push로 문자열을 추가 하면 char 타입으로 인식하지만, push_str는 문자열의 전체를 인식하여 메모리 확장에 사용한다.
//       bytes를 활용하여 각 문자열 값을 확인 하고자 유용하게 사용할 수 있을거라 생각한다.

fn get_hash_map() {
    let mut hash_a = Collection::hashmap();
    if let Collection::HashMap(hm) = &mut hash_a {
        hm.insert("HashMap_A".to_string(), 1);
        hm.insert("HashMap_B".to_string(), 2);
    }
    hash_a.display();

    let str = "HashMap String";
    let mut map = HashMap::new();
    for i in str.split_whitespace()
    {
        let count = map.entry(i).or_insert(0);
        *count += 1;
    }
    println!("{:?}", map);
}
// Note : HashMap는 <K, V>형식의 한 쌍으로 묶어 구조체를 반환한다. 이는 Do5라는 서비스 거부 공격을 저항 기능을 제공하는
//        알고리즘 이지만 속도는 빠르지 않다.
//        entry는 K의 접근을 관리하는 구조체를 반환하여 존재 여부에 따라 처리를 한다.
//        or_insert는 K가 없을 경우 기본 값을 삽입하고 해당 값의 가변 참조를 반환한다. 즉 K의 존재 여부 확인

fn string_replace() {
    let str = "String Replace";
    println!("{}", str.replace("String", "Re_Replace"));
}
// Note : Replace를 사용하면 특정 문자열를 다른 문자열로 변경하는 방법이다 str.replace(from, to)로 나누어 지는데
//        from을 기준으로 to의 문자열이 변경된다.