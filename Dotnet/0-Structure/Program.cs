using System;

namespace Public
{
  // Use public keyword is you can using member
  public class PublicCS
  {
    public struct Struct
    {
      // ...
    }

    public interface Interface
    {
      // ...
    }
  }

  namespace MultiNameSpace
  {
    // ...
  }
}

namespace Private
{
  // Not use public keyword is default private keyword in the outside cannot access
  private class MenualPrivateClass
  {
    private struct Struct
    {
      // ...
    }

    private interface Interface
    {
      // ...
    }
  }
}


namespace Default_None_Public
{
  class Default_PricateClass
  {
    struct Default_Struct
    {
      // ...
    }

    interface Interface
    {
      // ...
    }

    static void Main(string[] args)
    {
      Console.WriteLine("Hello World");
    }
  }
}