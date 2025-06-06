public class Array
{
  public static void CreatedArr()
  {
    int[] arr = new int[10]; // Create new array element to not have index however heap for memory size is 10 haved
    Console.WriteLine($"Heap size : {arr.Length} and index check : {arr}"); // The empty is array element output to datatype
  }

  public static void AddIndexArr()
  {
    int[] arr = new int[5];
    arr[1] = 15;
    arr[2] = 30;

    int v = arr[0];
    Console.WriteLine($"Variable value is : ({v}) and all array element in the index : ");
  }
}

class CompilerCS
{
  public static void Main(string[] args)
  {
    Array.CreatedArr();
    Array.AddIndexArr();
  }
}