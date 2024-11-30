# include <iostream>

using namespace std;

int main()
{
  int input_number;

  cout << "Enter the your Number Between 1 and 100 : ";

  cin >> input_number;

  cout << "Your Number is : " << input_number << endl;

  return 0;
}
// Note : cin은 Keyboard의 Input값을 받아 값을 저장한다. 13 Line에서 cin의 값을 반환하여 << String << Input << 순으로 메모리 할당 및 해제 진행