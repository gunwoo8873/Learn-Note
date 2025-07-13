using System;

namespace Public
{
  // Use public keyword is you can using member
  public class PublicCS
  {
    //...
    public struct Struct
    {
      //...
    }

    public interface Interface
    {
      //...
    }
  }
}

namespace Private
{
  // Not use public keyword is default private keyword in the outside cannot access
  private class Menual_PrivateClass
  {
    //...
    private struct Struct
    {
      //...
    }
  }
}


namespace Default_None_Public
{
  class Default_PricateClass
  {
    //...
    struct Default_Struct
    {
      //...
    }

    static void Main(string[] args)
    {
      Console.WriteLine("Hello World");
    }
  }
}