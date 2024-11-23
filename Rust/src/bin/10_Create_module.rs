fn main()
{
    module_group_a::module_a();
    module_group_b::module_c();
}

mod module_group_a
{
    pub fn module_a() {}
    pub fn module_b() {}
}

pub mod module_group_b
{
    pub fn module_c() {}
    pub fn module_d() {}
}
// Note : mod는 코드를 구조화 하고 모듈로 만들어 다른 파일에서 쉽게 불러 올 수 있도록 하는 방법이다.
//        하지만 pub 키워드를 사용하지 않으면 현재 작업중인 파일 내에서만 사용이 가능하기 때문에 유의해야 한다.
//        pub을 사용하여 경로를 조회하여 별도 경로를 지정할 필요가 없지만 쉽게 관리하기 위해서 lib.rs 파일을 사용한다.