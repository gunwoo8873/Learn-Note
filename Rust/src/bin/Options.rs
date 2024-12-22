fn main() {
    options();
    match_options();
    branches_options();
    current_status();
}

#[cfg(test)]


enum Option<T>
{
    None,
    Some(T),
}

fn options() {
    let int_some = Some(15);
    let str_some = Some("Option String");
    let none_some = None;
    let update_some = none_some.unwrap_or(15);
    println!("{}", update_some);
}
// Note : Option<Type>을 지정해도 Value 값이 None으로 명시되어 있다면, 타입을 추론이 불가능 하기 때문에
//        컴파일 에러가 발생한다. 그나마 해결 방안으로 unwarp_or_else, unwrap_or를 사용하여 값을 반환이 가능하다.

fn match_options() {
    let int_some = Some(17);
    match int_some
    {
        Some(i) => println!("{}", i),
        None => println!("None")
    }
}

fn branches_options() {
    let int_some = Some(20);
    if let Some(i) = int_some
    {
        println!("{}", i);
    }
}

//// Example
enum Status
{
    Active,
    Inactive,
    Deleted,
}

impl Status {
    fn active(&self) -> bool
    {
        match self
        {
            Status::Active => true,
            _ => false
        }
    }

    fn inactive(&self) -> bool
    {
        match self
        {
            Status::Inactive => true,
            _ => false
        }
    }

    fn deleted(&self) -> bool
    {
        match self
        {
            Status::Deleted => true,
            _ => false
        }
    }
}
// Note : enum 필드에 존재하는 객체들을 각 fn안의 match를 활용하여 특정 트리거를 불러 올 수 있도록 한다.
//        하지만 struct와 유사한 특징을 가지고 있기 때문에 각 불러오는 객체 외 {_}를 활용해서 예외 처리를 해야한다.

fn current_status() {
    let status_active = Status::Active;
    println!("{}", status_active.active());

    let status_inactive = Status::Inactive;
    println!("{}", status_inactive.active());
}
