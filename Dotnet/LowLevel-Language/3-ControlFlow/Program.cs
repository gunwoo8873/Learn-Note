namespace LogicalNS
{
  public class LogicalOperator
  {
    public static void OpeningAngle()
    {
      int i = sizeof(int);
      byte b = sizeof(byte);

      bool bo = i < b;
      Console.WriteLine(bo);
    }

    public static void CloseingAngle()
    {
      int i = sizeof(int);
      byte b = sizeof(byte);

      bool bo = i > b;
      Console.WriteLine(bo);
    }
  }
}


class CompilerCS
{
  static void Main(string[] args)
  {
    SelectAngle.OpeningAngle();
    SelectAngle.CloseingAngle();
  }
}