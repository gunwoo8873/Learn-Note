# `""" """`으로 감싸면 여러 줄의 문자열을 표현이 가능하며 유사한 `printf`의 기능을 한다.
print("""
    Example Menus:
        1. Test Menu 1    1 or -1       Test Description 1
        2. Test Menu 2    2 or -2       Test Description 2
        3. Test Menu 3    3 or -3       Test Description 3
        4. Test Menu 4    4 or -4       Test Description 4
""");

UserInput = input("Please select a menu (1-4):");

match UserInput:
    # `case <Parameter> | <Parameter>`으로 여러 값을 하나의 case에서 처리를 하도록 정의할 수 있다.
    case "1" | "-1":
        print(f"{UserInput} You have selected Test Menu 1");
    case "2" | "-2":
        print(f"{UserInput} You have selected Test Menu 2");
    case "3" | "-3":
        print(f"{UserInput} You have selected Test Menu 3");
    case "4" | "-4":
        print(f"{UserInput} You have selected Test Menu 4");
    case _: # '_' Default case
        print("Invalid selection. Please choose a number between 1 and 4.");