using System;

public record PublicRecord {}

///////////// Public /////////////
// Public keyword is used to declare a class, struct, interface,
// or namespace that can be accessed from any other
// code in the same assembly or another assembly that references it.
public class PublicClass {
    public struct Struct {}

    public interface Interface {
      // ...
    }

    public class SubClassA : Interface {
      public void Method(int parameter) {
        // ...
      }
    }
  }

  ///////////// Private /////////////
  // Private keyword is not used to declare a class, struct, interface,
  // or namespace that can only be accessed within the same class or struct.
  // Private members are not accessible outside the class or struct they are declared in.
  // Note : Not use public keyword to default to private.
private class PrivateClass {
  struct PrivateStruct {}

  interface PrivateInterface {
    // ...
  }
}

delegate int DelegateType(int a, int b);
// Delegate keyword is used to declare a delegate type.
// A delegate is a type that represents references to methods with a specific parameter list and return type.

class Compiler {
  public static void Main(string[] args) {
  }
}
