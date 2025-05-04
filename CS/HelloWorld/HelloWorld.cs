class HelloWorldCS {
  static void Main(string[] args) {
    Console.WriteLine("Hello, World!");
    
    // const keyword is only one used for mutable values.
    const string immutable_value = "Mutable Value";
    Console.WriteLine(immutable_value);

    // not used const keyword to then after can do update value
    string mutable_value = "Immutable Value";
    Console.WriteLine(mutable_value);

    // mutable value is changed to update the value.
    mutable_value = "Changed Immutable Value";
    Console.WriteLine(mutable_value);
  }
}
// Note :
// Class name change is possible and this is developer's choice.

// Warning :
// Don`t change the class name, it is a convention to use PascalCase for class names. 
// Need to Main method to run the program the don't change identifier[식별자] method name
// Main method is the entry point of the program.