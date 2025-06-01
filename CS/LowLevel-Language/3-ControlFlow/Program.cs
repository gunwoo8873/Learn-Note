public class SelectAngle
{
  public static void OpeningAngle()
  {
    int i = sizeof(int);
    char c = sizeof(char);

    bool b = (int)i < (char)c;
    Console.WriteLine(b);
  }

  public static void CloseingAngle()
  {
    int i = sizeof(int);
    char c = sizeof(char);

    bool b = (i > c);
    Console.WriteLine(b);
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