
public class Arr
{
  public static void CreateArr()
  {
    int[] products = new int[5];

    for (int i = 0; i < 5; i++)
    {
      products[i] = i;
      Console.WriteLine(products);
    }
  }
}


class Program
{
  static void Main(string[] args)
  {
    Arr.CreateArr();
  }
}