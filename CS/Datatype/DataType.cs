using System;
using System.Runtime.InteropServices;
using System.Text;

namespace OutputStyle {
  public class OutputStyleClass {
    public static void StringPrintHeader() {
      Console.WriteLine(new string('-', 40));
      Console.WriteLine($"{"String", -10} | {"UTF8 Size", -10} | {"Unicode Size", -10}");
      Console.WriteLine(new string('-', 40));
    }

    public static void CharPrintHeader() {
      Console.WriteLine(new string('-', 40));
      Console.WriteLine($"{"Char", -10} | {"UTF8 Size", -10} | {"Unicode Size", -10}");
      Console.WriteLine(new string('-', 40));
    }

    public static void IntegerPrintHeader() {
      Console.WriteLine(new string('-', 75));
      Console.WriteLine($"{"Integer Type", -10} | {"Bit", -10} | {"Min Value", -20} | {"Max Value", -20}");
      Console.WriteLine(new string('-', 75));
    }

    public static void RealNumberPrintHeader() {
      Console.WriteLine(new string('-', 100));
      Console.WriteLine($"{"RealNumber Type", -10} | {"Bit", -10} | {"Min Value", -30} | {"Max Value", -30}");
      Console.WriteLine(new string('-', 100));
    }
  }
}

namespace StringNS {
  // String is a nomal text type.
  // String text to EN is 1 byte size and KO is 2 byte size.
  public class StringClass {
    const string strEN = "A";
    const string strKO = "ㄱ";

    public static void StringFN() {
      OutputStyle.OutputStyleClass.StringPrintHeader();

      int sizeUTF8EN = Encoding.UTF8.GetByteCount(strEN); // UTF-8 is 1 byte size.
      int sizeUnicodeEN = Encoding.Unicode.GetByteCount(strEN); // UTF-16 is 2 byte size.
      Console.WriteLine($"{strEN, -10} | {sizeUTF8EN, -10} | {sizeUnicodeEN, -10}");

      int sizeUTF8KO = Encoding.UTF8.GetByteCount(strKO); // UTF-8 is 3 byte size.
      int sizeUnicodeKO = Encoding.Unicode.GetByteCount(strKO); // UTF-16 is 2 byte size.
      Console.WriteLine($"{strKO, -9} | {sizeUTF8KO, -10} | {sizeUnicodeKO, -10}");

      Console.WriteLine(new string('-', 40));
    }
  }

  public class CharClass {
    // Char is a single character type.
    const char chEN = 'A'; // UTF-8 is 1 byte size.
    const char chKO = 'ㄱ'; // UTF-8 is 3 byte size.

    public static void CharFN() {
      OutputStyle.OutputStyleClass.CharPrintHeader();

      int sizeUTF8ENCh = Encoding.UTF8.GetByteCount(chEN.ToString()); // UTF-8 is 1 byte size.
      int sizeUnicodeENCh = Encoding.Unicode.GetByteCount(chEN.ToString()); // UTF-16 is 2 byte size.
      Console.WriteLine($"{chEN, -10} | {sizeUTF8ENCh, -10} | {sizeUnicodeENCh, -10}");

      int sizeUTF8KOCh = Encoding.UTF8.GetByteCount(chKO.ToString()); // UTF-8 is 3 byte size.
      int sizeUnicodeKOCh = Encoding.Unicode.GetByteCount(chKO.ToString()); // UTF-16 is 2 byte size.
      Console.WriteLine($"{chKO, -9} | {sizeUTF8KOCh, -10} | {sizeUnicodeKOCh, -10}");

      Console.WriteLine(new string('-', 40));
    }
  }
}
// UTF
// UTF-8   : 가변 길이 인코딩 방식으로, ASCII 문자(1바이트)와 비ASCII 문자(2~4바이트)를 혼합하여 사용한다. 웹과 데이터베이스에서 널리 사용된다.
// UTF-16  : 가변 길이 인코딩 방식으로, 기본 다국어 평면(BMP) 문자는 2바이트로 표현하고, 보조 평면(SMP) 문자는 4바이트로 표현한다. 주로 Windows 운영체제에서 사용된다.
// UTF-32  : 4바이트 고정 길이 인코딩 방식으로, 모든 문자를 동일한 크기로 표현한다. 주로 내부 처리 및 데이터베이스에서 사용된다.
// Unicode : 모든 문자를 표현하기 위한 국제 표준으로, 각 문자는 2 Byte(16Bit)로 표현된다.

namespace IntegerNS {
  public class IntegerClass {
    public static void Integer_8bit() {
      // Signed Integer
      Console.WriteLine($"{"sbyte", -12} | {sizeof(sbyte) *8, -10} | {sbyte.MinValue, -20} | {sbyte.MaxValue, -10}");
      // Unsigned Integer
      Console.WriteLine($"{"byte", -12} | {sizeof(byte) *8, -10} | {byte.MinValue, -20} | {byte.MaxValue, -10}");
    }

    public static void Integer_16bit() {
      // Signed Integer
      Console.WriteLine($"{"short", -12} | {sizeof(short) *8, -10} | {short.MinValue, -20} | {short.MaxValue, -10}");
      // Unsigned Integer
      Console.WriteLine($"{"ushort", -12} | {sizeof(ushort) *8, -10} | {ushort.MinValue, -20} | {ushort.MaxValue, -10}");
    }
    public static void Integer_32bit() {
      // Signed Integer
      Console.WriteLine($"{"int", -12} | {sizeof(int) *8, -10} | {int.MinValue, -20} | {int.MaxValue, -10}");
      // Unsigned Integer
      Console.WriteLine($"{"uint", -12} | {sizeof(uint) *8, -10} | {uint.MinValue, -20} | {uint.MaxValue, -10}");
    }
    public static void Integer_64bit() {
      // Signed Integer
      Console.WriteLine($"{"long", -12} | {sizeof(long) *8, -10} | {long.MinValue, -20} | {long.MaxValue, -10}");
      // Unsigned Integer
      Console.WriteLine($"{"ulong", -12} | {sizeof(ulong) *8, -10} | {ulong.MinValue, -20} | {ulong.MaxValue, -10}");
    }

    public static void IntegerFN() {
      OutputStyle.OutputStyleClass.IntegerPrintHeader();
      IntegerClass.Integer_8bit();
      IntegerClass.Integer_16bit();
      IntegerClass.Integer_32bit();
      IntegerClass.Integer_64bit();

      Console.WriteLine(new string('-', 75));
    }
  }
}
// ------------------------------------------------------------------------
// Integer Type | Bit        | Min Value            | Max Value
// ------------------------------------------------------------------------
// sbyte        | 8          | -128                 | 127       
// byte         | 8          | 0                    | 255
// short        | 16         | -32768               | 32767
// ushort       | 16         | 0                    | 65535
// int          | 32         | -2147483648          | 2147483647
// uint         | 32         | 0                    | 4294967295
// long         | 64         | -9223372036854775808 | 9223372036854775807
// ulong        | 64         | 0                    | 18446744073709551615
// ------------------------------------------------------------------------

namespace RealNumberNS {
  public class FloatClass {
    public static void FloatFN() {
      Console.WriteLine($"{"Float", -15} | {sizeof(float), -10} | {float.MinValue, -30} | {float.MaxValue, -10}");
    }
  }

  public class DoubleClass {
    public static void DoubleFN() {
      Console.WriteLine($"{"Double", -15} | {sizeof(double), -10} | {double.MinValue, -30} | {double.MaxValue, -10}");
    }
  }

  public class DecimalClass {
    public static void DecimalFN() {
      Console.WriteLine($"{"Decimal", -15} | {sizeof(decimal), -10} | {decimal.MinValue, -30} | {decimal.MaxValue, -10}");
    }
  }

  public class RealNumberClass {
    public static void RealNumberFN() {
      OutputStyle.OutputStyleClass.RealNumberPrintHeader();
      FloatClass.FloatFN();
      DoubleClass.DoubleFN();
      DecimalClass.DecimalFN();

      Console.WriteLine(new string('-', 100));
    }
  }
}
// ----------------------------------------------------------------------------------------------------
// RealNumber Type | Bit        | Min Value                      | Max Value
// ----------------------------------------------------------------------------------------------------
// Float           | 4          | -3.4028235E+38                 | 3.4028235E+38
// Double          | 8          | -1.7976931348623157E+308       | 1.7976931348623157E+308
// Decimal         | 16         | -79228162514264337593543950335 | 79228162514264337593543950335
// ----------------------------------------------------------------------------------------------------


// namespace BooleanNS {
//   public class BooleanClass {}
// }

class DatatypeClass {
  public static void Main(string[] args) {
    StringNS.StringClass.StringFN();
    StringNS.CharClass.CharFN();

    IntegerNS.IntegerClass.IntegerFN();

    RealNumberNS.RealNumberClass.RealNumberFN();
  }
}