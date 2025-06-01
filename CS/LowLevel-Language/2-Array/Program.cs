public class ArrayCS
{
  public static void NomalArr()
  {
    for (int i = 0; i < 10; i++)
    {
      int[] arr = new int[10]{i}; // Create new array for int datatype variable value to Heap
      Console.WriteLine(arr);
    }
  }
}

class CompilerCS
{
  public static void Main(string[] args)
  {
    ArrayCS.NomalArr();
  }
}