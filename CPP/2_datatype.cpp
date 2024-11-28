# include <iostream>

using namespace std;

void return_void()
{}

string return_str()
{
  string a = "Return String Datatype";
  return 0;
}

char return_char()
{
  char a = 'c';
  return 0;
}

/* Char 16 bit type */
char16_t return_char_16()
{
  char16_t a = u'c';
  return 0;
}

/* Char 32 bit type */
char32_t return_char_32()
{
  char32_t a = U'c';
  return 0;
}

int return_int()
{
  int long a = 15;
  return 0;
}

float return_float()
{
  float a = 15.000;
  return 0;
}
// Note : 각 Return Type타입을 명시하여 int test() {}나 void, string, double 등 다양한 Return Type을 활용해야 한다.

int main()
{
  return_void();
  return_int();
  return_float();
  return_str();
  return_char();
  return_char_16();
  return_char_32();
}