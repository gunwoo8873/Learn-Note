fn main()
{
    immutable_instance();
    mutable_instance();
    tuple();
}

#[derive(Debug)]
struct Field
{
    str: String,
    i32: i32,
    bool: bool
}
// Note : Struct는 구조체라고 각 필드에 대한 구체적인 값을 추가하는 것이 아닌 각 다른 type을 명시가 가능하다.

fn immutable_instance()
{
    let field_a = Field
    {
        str: String::from("Field String"),
        i32: 1,
        bool: true,
    };

    println!("Field A : {:?}", field_a.str);
    println!("Field A : {:?}", field_a.i32);
    println!("Field A : {:?}", field_a.bool);
}

fn mutable_instance()
{
    let mut field_b = Field
    {
        str: String::from("Field String 1"),
        i32: 1,
        bool: true,
    };

    field_b.i32 = 2;

    println!("Field B : {:?}", field_b.str);    // Result : Not Update
    println!("Field B : {:?}", field_b.i32);    // Result : Update
    println!("Field B : {:?}", field_b.bool);   // Result : Not Update

    let field_c = Field
    {
        str: String::from("Field String 2"),
        ..field_b
    };

    println!("Field C : {:?}", field_c.str);    // Result : Update
    println!("Field C : {:?}", field_c.i32);    // Result : Not Update
    println!("Field C : {:?}", field_c.bool);   // Result : Not Update
}
// Note : struct를 사용 하더라도 기본적으로 let의 mut 키워드 사용 여부에 따라 값 업데이트 차이가 존재한다.
//        일부 기존의 Instance를 사용하고 싶을 경우 ..를 붙여 사용하고자 하는 Instance를 불러오면 된다.

fn short_instance(str: String, i32: i32, bool: bool) -> Field
{
    Field
    {
        str,
        i32,
        bool,
    }
}
// Note : Short Instance와 같은 코드를 사용 하고자 하면, 기존 Field의 데이터 타입을 지정된 상태에서 추가적인
//        데이터 타입을 명시하지 않아도 된다.

struct RGB(i32, i32, i32); // x, y, z
impl RGB
{
    fn get_hex(&self) -> String
    {
        format!("#{:02X}{:02X}{:02X}", self.0, self.1, self.2)
    }
}
// Note : {:02X} : 16진수 형식으로 변환하는 포멧 지정자이다.
//        [:] : 포멧 지정자의 시작을 나타낸다.
//        02  : 2자리로 제한하여 출력하지만, 빈 자리는 0으로 자동으로 채운다.
//        X   : 16진수로 출력하고 문자열을 대문자로 변환 후 출력

fn tuple()
{
    let rgb = RGB(255, 255, 255);
    let hex = rgb.get_hex(); // rgb의 값을 가져와 get_hex의 포멧팅 형식으로 변환하여 값을 저장한다.

    println!("{}", hex);
}