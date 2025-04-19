using System;

namespace RootNamespace {
  public class PublicClass {
    struct PublicStruct {}

    interface PublicInterface {}
  }

  private class PrivateClass {
    struct PrivateStruct {}

    interface PrivateInterface {}
  }

  namespace SubNamespace {}

  delegate int DelegateType(int a, int b);

  class Compiler {
    public static void Main(string[] args) {
      Console.WriteLine("Hello, World!");
    }
  }
}