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
    int intValue = int.MaxValue;
    long longValue = intValue;
    long longUpdateValue = longValue;
    longUpdateValue = long.MaxValue; // int is changed to long and changed value update for MaxValue
    Console.WriteLine($"Datatype chagned int : {intValue} -> long : {longUpdateValue}");
  }
}

public class BooleanType
{
  public static void BooleanFN()
  {
    bool bo = true;
    Console.WriteLine(bo);

    if (bo == true)
    {
      Console.WriteLine($"{bo} == true ");
    }
  }
}

class CompilerCS
{
  public static void Main(string[] args)
  {
    IntegerType.IntegerFN();
    ChagneType.ChangeTypeFN();
  }
}