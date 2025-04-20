using System;

namespace RootNamespace {
  public class PublicClass {
    struct PublicStruct {}

    interface PublicInterface {}
  }
  // Public keyword is used to declare a class, struct, interface,
  // or namespace that can be accessed from any other
  // code in the same assembly or another assembly that references it.

  private class PrivateClass {
    struct PrivateStruct {}

    interface PrivateInterface {}
  }
  // Private keyword is not used to declare a class, struct, interface,
  // or namespace that can only be accessed within the same class or struct.
  // Private members are not accessible outside the class or struct they are declared in.

  namespace SubNamespace {}

  delegate int DelegateType(int a, int b);

  class Compiler {
    public static void Main(string[] args) {
      Console.WriteLine("Hello, World!");
    }
  }
}