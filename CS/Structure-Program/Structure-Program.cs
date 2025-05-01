using System;

namespace RootNamespace {
  ///////////// Public /////////////
  // Public keyword is used to declare a class, struct, interface,
  // or namespace that can be accessed from any other
  // code in the same assembly or another assembly that references it.
  public class PublicClass {
    public struct PublicStruct {}

    public interface PublicInterface {
      // Method in the public interface
      void Method();
    }
  }

  ///////////// Private /////////////
  // Private keyword is not used to declare a class, struct, interface,
  // or namespace that can only be accessed within the same class or struct.
  // Private members are not accessible outside the class or struct they are declared in.
  private class PrivateClass {
    private struct PrivateStruct {}

    private interface PrivateInterface {
      // Method in the private interface
      void Method();
    }
  }

  namespace SubNamespace {}

  delegate int DelegateType(int a, int b);
  // Delegate keyword is used to declare a delegate type.
  // A delegate is a type that represents references to methods with a specific parameter list and return type.

  class Compiler {
    public static void Main(string[] args) {
      RootNamespace obj = new PublicClass();
    }
  }
}