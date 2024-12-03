## Using
```cpp
using namespace std; // std라는 이름을 가진 식별자 영역을 사용한다.
```
```txt
using은 식별자들이 사용되고 효율성을 위해서 몇 개의 영역으로 구성되어 있다.
```

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
    | float         |   4   | 1.2E-38 ~ 3.4E38                                       |
    | double        |   8   | 2.2E-308 ~ 1.8E308                                     |
    | unsgined      |   -   | -                                                      |
    | un short      |   2   | 0 ~ 65535                                              |
    | un int / long |   4   | 0 ~ 4294967295                                         |

## ETC
* C++ Keywords : [CppReference](https://en.cppreference.com/w/cpp/keyword)