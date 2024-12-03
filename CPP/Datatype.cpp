# include <iostream> // Header File을 포함시키는 코드 [Screen Output이나 Keyboard Input 기능 제공]

using namespace std; // C++의 표준 lib이며 라이브러리를 namespace라는 명칭을 사용

void int_type()
{
  short a_max = SHRT_MAX;
  short a_min = SHRT_MIN;
  cout << "short Size : " << sizeof(a_max) << " , Byte Max : " << a_max << " , Byte Min : " << a_min << endl; // Result : 2

  int b_max = INT_MAX;
  int b_min = INT_MIN;
  cout << "int Size : " << sizeof(b_max) << " , Byte Max : " << b_max << " , Byte Min : " << b_min << endl; // Result : 4

  long c_max = LONG_MAX;
  long c_min = LONG_MIN;
  cout << "long Size : " << sizeof(c_max) << " , Byte Max : " << c_max << " , Byte Min : " << c_min << endl; // Result : 4

  long long d_max = LLONG_MAX;
  long long d_min = LLONG_MIN;
  cout << "long long Size : " << sizeof(d_max) << " , Byte Max : " << d_max << " , Byte Min : " << d_min << endl; // Result : 8
}

void char_type()
{
  char a = 'c';
  char a_max = CHAR_MAX;
  char a_min = CHAR_MIN;
  cout << "char Size : " << sizeof(a) << " , Byte Max : " << a_max << " , Byte Min : " << a_min << endl; // Result : 1

  char16_t b = u'c';
  cout << "char16_t Size : " << sizeof(b) << endl; // Result : 2

  char32_t c = U'c';
  cout << "char32_t Size : " << sizeof(c) << endl; // Result : 4
}

void float_type()
{
  float a = 10.0e-3f;
  cout << "float Size : " << sizeof(a) << " Float Result : " << a << endl; // Result : 4

  double b = 10.0;
  cout << "double Size : " << sizeof(b) << " Double Result : " << b << endl; // Result : 8
}

// Note : sizeof 연산자를 사용하여 각 데이터 타입들의 Size 확인이 가능하다. 
//        자료형이나 변수의 Size를 Byte 단위로 계산하는 연산자다.

int main()
{
  int_type();
  char_type();
  float_type();
  return 0;
}