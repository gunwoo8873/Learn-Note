# include <iostream>

using namespace std;

void overflow()
{
  short a = SHRT_MAX;
  short over_a = a + 1;
  cout << "Overflow : " << over_a << endl;
}
// Note : Overflow는 변수에 최대값을 초과하면 발생하는 문제다. Overflow는 발생하더라도 컴파일러는 경고를 하지 않기 때문에
//        전체적으로 부정확한 결과가 계산값을 출력한다.

void radix()
{
  int binary = 0b000001; // 2 base
  cout << "Binary : " << binary << endl;

  int octal = 062; // 8 base
  cout << "Octal : " << octal << endl;
  
  int decimal = 10; // 10 base
  cout << "Decimal : " << decimal << endl;

  int hexadecimal = 0x2A; // 16 base
  cout << "Hexadecimal : " << hexadecimal << endl;
}
// Note : 각 진수에 형태로 코드를 작성을 하더라도 컴파일 단계에서 10 진수로 변환하여 출력한다.

int main()
{
  overflow();
  radix();
}