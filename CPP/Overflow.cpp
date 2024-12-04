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

#define TAX_RATE 0.15
// Note : define은 컴파일러가 수행하기 전 Preprocess-sor[전처리기]가 처리한다.

void define()
{
  double m_salary, y_salary;
  cin >> m_salary;

  y_salary = m_salary * (1 + TAX_RATE);
  cout << "Monthly Salary : " << y_salary << endl;
}

int main()
{
  overflow();
  radix();
  define();
}