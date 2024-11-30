### Instruction [명령어]
```txt
CPU가 수행하는 기초적인 연산들 이다.
```

### Programming Language
```txt
프로그래밍 언어는 인간과 기계 사이의 번역을 맡는 역할 이라고 생각하면 된다.
Programming Language를 통해 원하는 작업을 텍스트로 기술한 것을 source code라고 칭하고, 
파일 내에 저장된 것을 source file 라고 부르고, Compiler는 source file을 기계어로 변환하여 파일에 저장하는데, 이를 object file이라 부른다.
```

### Procedural Programming
```txt
절차적 프로그래밍은 절차를 중시하며 이러한 프로그래밍 기법이란 의미를 가진다.
```

### Object-Oriented Programming [OOP]
```txt
Procedural Programming의 문제점을 대응하기 위해 객체 지향 프로그래밍이 개발 되었다. 알고리즘 중점이 아닌 객체 지향으로 데이터를 중점으로 시작으로 문제에 필요한 데이터를 먼저 설계하는 기법이다.
데이터를 나타낼 때 사용하는 도구는 Class로 데이터 외에 데이터를 처리하는 함수가 포함된다.
```
> OOP는 데이터와 알고리즘이 분리되어 있지 않는다. 추가로 C++는 OOP 기법을 베이스로 설계되어 있다.

## 규칙
* 기본 소스 문자 집합
  ```txt
  C++ 표준에서는 소스 파일에서 사용할 수 있는 기본 소스 문자 집합을 지정한다.  
  외부에 문자를 나타내고자 한다면 유니버설 문자 이름을 사용하여 추가 문자 지정이 가능하다.  
  ```
  > Note : MSVC 구현에서는 추가 문자 허용
  * 그래픽 문자 집합
    ```txt
    a b c d e f g h i j k l m n o p q r s t u v w x y z

    A B C D E F G H I J K L M N O P Q R S T U V W X Y Z

    0 1 2 3 4 5 6 7 8 9

    _ { } [ ] # ( ) < > % : ; . ? * + - / ^ & | ~ ! = , \ " '
    ```

  * Data type list
    ```txt
    alignas, alignof, and, and_eq, asm, auto, bitand, bitor, bool, break, case, catch, char, char16_t, char32_t, class, compl, const, constexpr, const_cast, continue, decltype, default, delete, do, double, dynamic_cast, else, enum, explicit, export, extern, false, float, for, friend, goto, if, inline, int, long, mutable, namespace, new, noexcept, not, not_eq, nullptr, operator, or, or_eq, private, protected, public, register, reinterpret_cast, return, short, signed, sizeof, static, static_assert, static_cast, struct, switch, template, this, thread_local, throw, true, try, typedef, typeid, typename, union, unsigned, using, virtual, void, volatile, wchar_t, while, xor, xor_eq
    ```
    > 이 미친 데이터 타입과 연산자를 보면 Low level System Programing Language인 Rust와 굉장히 경악할 정도로 대조된다.

  * Data type Size
    | Type          | Byte  | Size                                                   |
    | :------------ | :---: | :----------------------------------------------------- |
    | char          |   1   | -128 ~ 127                                             |
    | char 16b, 32b |   -   |                                                        |
    | short         |   2   | -32,768 ~ 32,767                                       |
    | int / long    |   4   | -2,147,483,648 ~ 2,147,483,647                         |
    | long long     |   8   | -9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807 |


## ETC
* Algorithm : 문제를 해결하기 위한 일련의 정확한 단계를 의미, 계산 수행, 데이터 처리, 자동 추론 등 다양한 작업에 사용
* C++ Keywords : [CppReference](https://en.cppreference.com/w/cpp/keyword)