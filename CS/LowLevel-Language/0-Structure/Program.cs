using System;

namespace NS
{
  // Use public keyword is you can using member
  public class PublicCS
  {
    public static void Function(string[] args)
    {
      Console.WriteLine("Public Funtion Output");
    }

    public struct Struct
    {
      // ...
    }

    public interface Interface
    {
      // ...
    }
  }

  // Not use public keyword is default private keyword in the outside cannot access
  class PrivateCS
  {
    static void Function(string[] args)
    {
      Console.WriteLine("Private Function Output");
    }

    struct Struct
    {
      // ...
    }

    interface Interface
    {
      // ...
    }
  }

  namespace SubNS
  {
    
  }

  class Complier
  {
    public static void Main(string[] args)
    {
      PublicCS.Function(args); // Can public class access
      PrivateCS.Function(args); // Private class cannot access [Error Code : CS0122]
    }
  }
}