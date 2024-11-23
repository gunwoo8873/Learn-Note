fn main()
{
    generic();
}
// Note : Generic 중복되는 코드를 효율적으로 처리하기 위한 도구다. Concrete[구체] 타입 혹은 ETC 속성에 대한 추상화된
//        대역으로 컴파일과 실행 시점에 어떤 타입으로 채워져 있는지 확일 할 필요 없이 동작이나 관계를 표현이 가능하다.

fn generic()
{
    let num_arr = vec![40, 60, 70, 80]; // Int 리스트를 저장 [Collection의 Vector참고]
    let mut largest = &num_arr[0];

    for i in &num_arr
    {
        if i > largest
        {
            largest = i;
        }
    }

    println!("Last Largest num : {}", largest); // Result = 80
}
// Note : &는 참조 즉 데이터를 읽을 수 있도록 하되 mut의 존재 여부에 따라 읽고 수정 범위가 달라진다.