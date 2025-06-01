public class IntegerType
{
  public static void IntegerFN()
  {
    sbyte sb = sbyte.MaxValue; // +-127 : Have sign 8 bit
    Console.WriteLine(sb);
    byte b = byte.MaxValue; // +255 : Not have sign 8 bit
    Console.WriteLine(b);

    short s = short.MaxValue; // +-32,767 : Have sign 16 bit
    Console.WriteLine(s);
    ushort us = ushort.MaxValue; // +65,535 : Not sign 16 bit
    Console.WriteLine(us);

    int i = int.MaxValue; // +-2,147,483,647 : Have sign 32 bit
    Console.WriteLine(i);
    uint ui = uint.MaxValue; // +4,294,967,295 : Not have sign 32 bit
    Console.WriteLine(ui);

    long l = long.MaxValue; // +- 9,223,372,036,854,775,807 =  : Have sign 64 bit
    Console.WriteLine(l);
    ulong ul = ulong.MaxValue; // +18,446,744,073,709,551,615 : Not have sign 64 bit
    Console.WriteLine(ul);
  }
}

public class ChagneType
{
  public static void ChangeTypeFN()
  {
    // Don`t Declaration to Implicit Conversion to need Explicit Conversion.
    int i = int.MaxValue;

    // Chagned int datatype to long datatype
    long l = (long)i;
    Console.WriteLine($"int to long changed datatype : int ({i}) => long ({l})");
  }
}

public class BooleanType
{
  public static void BooleanFN()
  {
    bool b = true;
    Console.WriteLine(b);
  }
}

public class StringType
{
  public static void StringTypeFN()
  {
    string s = "Hello Wrold!";
    /*
    sizeof keyword is data type size check then only english is
    not UTF-8 changed encoding however other country language is need to UTF-8 encoding
    */
    Console.WriteLine($"String output : {s}, size : {s.Length}");
  }

  public static void CharTypeFN()
  {
    char c = 'a';
    Console.WriteLine($"Char output : {c}, size : {sizeof(char)}");
  }
}

class CompilerCS
{
  public static void Main(string[] args)
  {
    IntegerType.IntegerFN();
    ChagneType.ChangeTypeFN();

    StringType.StringTypeFN();
    StringType.CharTypeFN();
  }
}