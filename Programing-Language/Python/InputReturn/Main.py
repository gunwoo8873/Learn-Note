print("""
    Example Menus:
        1. Test Menu 1
        2. Test Menu 2
        3. Test Menu 3
        4. Test Menu 4
""");

UserInput = input("Please select a menu (1-4):");

match UserInput:
    case "1":
        print("You have selected Test Menu 1");
    case "2":
        print("You have selected Test Menu 2");
    case "3":
        print("You have selected Test Menu 3");
    case "4":
        print("You have selected Test Menu 4");
    case _: # '_' Default case
        print("Invalid selection. Please choose a number between 1 and 4.");