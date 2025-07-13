public class IntegerType
{
  public static void IntegerFN()
  {
    sbyte sb = sbyte.MaxValue; // +-127 : Have sign 8 bit and 1 byte
    Console.WriteLine($"sbyte output : {sb}, size : {sizeof(sbyte)}");
    byte b = byte.MaxValue; // +255 : Not have sign 8 bit and 1 byte
    Console.WriteLine($"byte output : {b}, size : {sizeof(byte)}");

    short s = short.MaxValue; // +-32,767 : Have sign 16 bit and 2 byte
    Console.WriteLine($"short output : {s}, size : {sizeof(short)}");
    ushort us = ushort.MaxValue; // +65,535 : Not sign 16 bit and 2 byte
    Console.WriteLine($"ushort output : {us}, size : {sizeof(ushort)}");

    int i = int.MaxValue; // +-2,147,483,647 : Have sign 32 bit and 4 byte
    Console.WriteLine($"int output : {i}, size : {sizeof(int)}");
    uint ui = uint.MaxValue; // +4,294,967,295 : Not have sign 32 bit and 4 byte
    Console.WriteLine($"uint output : {ui}, size : {sizeof(uint)}");

    long l = long.MaxValue; // +- 9,223,372,036,854,775,807 =  : Have sign 64 bit and 8 byte
    Console.WriteLine($"long output : {l}, size : {sizeof(long)}");
    ulong ul = ulong.MaxValue; // +18,446,744,073,709,551,615 : Not have sign 64 bit and 8 byte
    Console.WriteLine($"ulong output : {ul}, size : {sizeof(ulong)}");
  }
}

public class FloatType
{
  public static void FloatTypeFN()
  {
    float f = float.MaxValue; // +-1.5e-45 ~ +-3.4e38 : 4 byte
    Console.WriteLine($"float output : {f}, size : {sizeof(float)}");
  }

  public static void DoubleTypeFN()
  {
    double d = double.MaxValue; // +-5.0e-324 ~ +-1.7e308 : 8 byte
    Console.WriteLine($"double output : {d}, size : {sizeof(double)}");
  }

  public static void DecimalTypeFN()
  {
    decimal de = decimal.MaxValue; // +-1.0 * 10-28 ~ +-7.9 * 1028 : 16 byte
    Console.WriteLine($"decimal output : {de}, size : {sizeof(decimal)}");
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

    FloatType.FloatTypeFN();
    FloatType.DoubleTypeFN();
    FloatType.DecimalTypeFN();
  }
}