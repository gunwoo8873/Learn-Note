Console.WriteLine("1. Standard String Output : A");

// '\t'을 포함하면 터미널 출력에서 Tap 한번의 띄워쓰기가 적용되어 Format과 유사한 요소를 출력
Console.WriteLine("2. Tap Uniqce String Output :\tB"); // Result : 2. Tap Uniqce String Output :    B

string example_str = "Example String Value";

// 문자열에서 요청하는 값을 검색하여 Bool 타입의 결과값을 출력한다.
Console.WriteLine("3. Search String value :" + example_str.Contains("Value"));

// ToUpper과 ToLower는 값의 모든 문자열을 대문자 혹은 소문자로 교체하여 결과값으로 저장하고 출력한다.
Console.WriteLine("4. Changed Upper string value :" + example_str.ToUpper()); // Result : Example String Value => EXAMPLE STRING VALUE
Console.WriteLine("4. Changed Lower string value :" + example_str.ToLower()); // Result : Example String Value => example string value

// '\\'에서 주의해야 할점은 '\'만 사용할 경우에 시퀸스가 예약되어 오류로 인식하여 컴파일 실패가 발생한다.
Console.WriteLine("5. Sequence : C:\\Doucment\\VSCode"); // Result : C:\Doucment\VSCode

string example_date_str = "2025.09.20";
Console.WriteLine("6. Marged String value :" + example_str + "\t" + example_date_str);