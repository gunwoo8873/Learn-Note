class HelloWorldCS {
  // Warning : Don`t change the class name, it is a convention to use PascalCase for class names.
  // Need to Main method to run the program the don't change identifier[식별자] method name
  // Main method is the entry point of the program.
  static void Main(string[] args) {
    Console.WriteLine("Hello, World!");
    
    const string mutable_value = "Mutable Value";
    Console.WriteLine(mutable_value);

    string immutable_value = "Immutable Value";
    Console.WriteLine(immutable_value);
  }
}
// Note : Class name change is possible and this is developer's choice.